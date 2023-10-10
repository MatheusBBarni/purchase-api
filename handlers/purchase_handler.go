package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/matheusBBarni/purchase-api/handlers/dto"
	"github.com/matheusBBarni/purchase-api/services"
)

func NewPurchaseHandler(router fiber.Router, purchaseService *services.PurchaseService) {
	router.Post("", func(c *fiber.Ctx) error {
		purchaseDto := new(dto.PurchaseDto)

		if err := c.BodyParser(purchaseDto); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		purchase, err := purchaseService.SavePurchase(purchaseDto)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(purchase)
	})

	router.Get("", func(c *fiber.Ctx) error {
		purchases := purchaseService.GetAllPurchases()

		return c.Status(http.StatusOK).JSON(purchases)
	})

	router.Get(":id/convert/:currency", func(c *fiber.Ctx) error {
		purchaseId := c.Params("id")
		currency := c.Params("currency")
		response, err := purchaseService.ConvertPurchaseAmount(purchaseId, currency)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(response)
	})
}
