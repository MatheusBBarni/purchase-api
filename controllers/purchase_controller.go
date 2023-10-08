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

	purchase, err := purchaseController.purchaseRepository.Save(purchaseDto.ToDomain())

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(purchase)
}

type currencies = []*dto.CurrencyDto
type currencyApiResponse struct {
	Data currencies `json:"data"`
}

func (purchaseController *PurchaseController) getExchangeRate(currencies currencies, selectedCurrency string) (float64, error) {
	for _, value := range currencies {
		if value.Currency == selectedCurrency {
			return strconv.ParseFloat(value.ExchangeRate, 64)
		}
	}

	return 0.0, errors.New("currency not found")
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

	exchangeRate, err := purchaseController.getExchangeRate(currencyApiResponse.Data, currency)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"exchangeRate": exchangeRate,
		"purchaseId":   purchaseId,
		"currency":     currency,
	})
}
