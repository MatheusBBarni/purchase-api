package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/matheusBBarni/purchase-api/config"
	"github.com/matheusBBarni/purchase-api/domain/dto"
	"github.com/matheusBBarni/purchase-api/repositories"
	"github.com/matheusBBarni/purchase-api/utils"
)

type PurchaseController struct {
	purchaseRepository *repositories.PurchaseRepository
	currencyHttpClient *config.HttpClient
}

func NewPurchaseController(purchaseRepository *repositories.PurchaseRepository, currencyHttpClient *config.HttpClient) *PurchaseController {
	return &PurchaseController{
		purchaseRepository: purchaseRepository,
		currencyHttpClient: currencyHttpClient,
	}
}

func (purchaseController PurchaseController) GetAllPurchases(c *fiber.Ctx) error {
	purchases := purchaseController.purchaseRepository.GetAll()

	return c.Status(http.StatusOK).JSON(purchases)
}

func (purchaseController *PurchaseController) SavePurchase(c *fiber.Ctx) error {
	purchaseDto := new(dto.PurchaseDto)

	if err := c.BodyParser(purchaseDto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := utils.ValidateStruct(purchaseDto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	purchaseTransformed, err := purchaseDto.ToDomain()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	purchase, err := purchaseController.purchaseRepository.Save(purchaseTransformed)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(purchase)
}

type currencyApiResponse struct {
	Data []*dto.CurrencyDto `json:"data"`
}

func (purchaseController *PurchaseController) getCurrency(currencies []*dto.CurrencyDto, selectedCurrency string) (*dto.CurrencyDto, error) {
	for _, value := range currencies {
		if value.Currency == selectedCurrency {
			return value, nil
		}
	}

	return &dto.CurrencyDto{}, errors.New("currency not found")
}

func (purchaseController *PurchaseController) convertAmount(amount, exchangeRate float64) float64 {
	return amount * exchangeRate
}

func (purchaseController *PurchaseController) ConvertPurchaseAmount(c *fiber.Ctx) error {
	purchaseId := c.Params("id")
	currency := c.Params("currency")
	var currencyApiResponse *currencyApiResponse

	err := purchaseController.currencyHttpClient.Get(&currencyApiResponse)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	selectedCurrency, err := purchaseController.getCurrency(currencyApiResponse.Data, currency)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	parsedId, err := utils.ConvertStringToUint(purchaseId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	purchase := purchaseController.purchaseRepository.GetById(parsedId)

	currencyRecordDate, err := utils.ConvertStringToDate(selectedCurrency.RecordDate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	dateDifferenceInMonths := purchase.TransactionDate.Sub(currencyRecordDate).Hours() / 24 / 30
	if dateDifferenceInMonths > 6 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Purchase cannot be converted to the target currency",
		})
	}

	exchangeRate, err := strconv.ParseFloat(selectedCurrency.ExchangeRate, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	convertedAmount := utils.ConvertFloatToTwoDecimals(purchaseController.convertAmount(purchase.Amount, exchangeRate))

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"id":               purchase.ID,
		"description":      purchase.Description,
		"transaction_date": purchase.TransactionDate,
		"purchase_amount":  purchase.Amount,
		"exchange_rate":    exchangeRate,
		"converted_amount": convertedAmount,
	})
}
