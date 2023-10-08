package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/matheusBBarni/purchase-api/domain/dto"
	"github.com/matheusBBarni/purchase-api/repositories"
	"github.com/matheusBBarni/purchase-api/utils"
)

type PurchaseController struct {
	purchaseRepository *repositories.PurchaseRepository
}

func NewPurchaseController(purchaseRepository *repositories.PurchaseRepository) *PurchaseController {
	return &PurchaseController{
		purchaseRepository: purchaseRepository,
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
