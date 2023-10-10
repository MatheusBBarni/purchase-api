package services

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/matheusBBarni/purchase-api/config"
	"github.com/matheusBBarni/purchase-api/domain"
	"github.com/matheusBBarni/purchase-api/handlers/dto"
	"github.com/matheusBBarni/purchase-api/repositories"
	"github.com/matheusBBarni/purchase-api/utils"
)

type PurchaseService struct {
	purchaseRepository *repositories.PurchaseRepository
	currencyHttpClient *config.HttpClient
}

func NewPurchaseService(purchaseRepository *repositories.PurchaseRepository, currencyHttpClient *config.HttpClient) *PurchaseService {
	return &PurchaseService{
		purchaseRepository: purchaseRepository,
		currencyHttpClient: currencyHttpClient,
	}
}

func (purchaseService PurchaseService) GetAllPurchases() []*domain.Purchase {
	purchases := purchaseService.purchaseRepository.GetAll()

	return purchases
}

func (purchaseService *PurchaseService) SavePurchase(purchaseDto *dto.PurchaseDto) (*domain.Purchase, error) {

	if err := utils.ValidateStruct(purchaseDto); err != nil {
		return nil, err
	}

	purchaseTransformed, err := purchaseDto.ToDomain()
	if err != nil {
		return nil, err
	}

	purchase, err := purchaseService.purchaseRepository.Save(purchaseTransformed)

	if err != nil {
		return nil, err
	}

	return purchase, nil
}

type currencyApiResponse struct {
	Data []*dto.CurrencyDto `json:"data"`
}

func (purchaseService *PurchaseService) getCurrency(currencies []*dto.CurrencyDto, selectedCurrency string) (*dto.CurrencyDto, error) {
	for _, value := range currencies {
		if value.Currency == selectedCurrency {
			return value, nil
		}
	}

	return &dto.CurrencyDto{}, errors.New("currency not found")
}

func (purchaseService *PurchaseService) convertAmount(amount, exchangeRate float64) float64 {
	return amount * exchangeRate
}

func (purchaseService *PurchaseService) ConvertPurchaseAmount(purchaseId string, currency string) (interface{}, error) {
	var currencyApiResponse *currencyApiResponse

	err := purchaseService.currencyHttpClient.Get(&currencyApiResponse)
	if err != nil {
		return nil, err
	}

	selectedCurrency, err := purchaseService.getCurrency(currencyApiResponse.Data, currency)
	if err != nil {
		return nil, err
	}

	parsedId, err := utils.ConvertStringToUint(purchaseId)
	if err != nil {
		return nil, err
	}
	purchase := purchaseService.purchaseRepository.GetById(parsedId)

	currencyRecordDate, err := utils.ConvertStringToDate(selectedCurrency.RecordDate)
	if err != nil {
		return nil, err
	}

	dateDifferenceInMonths := purchase.TransactionDate.Sub(currencyRecordDate).Hours() / 24 / 30
	if dateDifferenceInMonths > 6 {
		return nil, err
	}

	exchangeRate, err := strconv.ParseFloat(selectedCurrency.ExchangeRate, 64)
	if err != nil {
		return nil, err
	}

	convertedAmount := utils.ConvertFloatToTwoDecimals(purchaseService.convertAmount(purchase.Amount, exchangeRate))

	return fiber.Map{
		"id":               purchase.ID,
		"description":      purchase.Description,
		"transaction_date": purchase.TransactionDate,
		"purchase_amount":  purchase.Amount,
		"exchange_rate":    exchangeRate,
		"converted_amount": convertedAmount,
	}, nil
}
