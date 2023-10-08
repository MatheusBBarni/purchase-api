package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/matheusBBarni/purchase-api/config"
	"github.com/matheusBBarni/purchase-api/controllers"
	handlers "github.com/matheusBBarni/purchase-api/handlers"
	"github.com/matheusBBarni/purchase-api/repositories"
)

func main() {
	app := fiber.New()

	db, err := config.ConnectToDatabase()

	if err != nil {
		log.Panic("Could not connect to database!")
	}

	purchaseRepository := repositories.NewPurchaseRepository(db)
	purchaseController := controllers.NewPurchaseController(purchaseRepository)

	v1 := app.Group("/v1")
	handlers.NewPurchaseHandler(v1.Group("/purchase"), purchaseController)

	app.Listen(":8080")
}
