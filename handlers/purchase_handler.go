package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusBBarni/purchase-api/controllers"
)

func NewPurchaseHandler(router fiber.Router, purchaseController *controllers.PurchaseController) {
	router.Post("", purchaseController.SavePurchase)
	router.Get("", purchaseController.GetAllPurchases)
}
