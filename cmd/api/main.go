package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/matheusBBarni/purchase-api/config"
	handlers "github.com/matheusBBarni/purchase-api/handlers"
	"github.com/matheusBBarni/purchase-api/repositories"
	"github.com/matheusBBarni/purchase-api/services"
)

func main() {
	app := fiber.New()

	db, err := config.ConnectToDatabase()

	if err != nil {
		log.Panic("Could not connect to database!")
	}

	currencyApiClient := &http.Client{Timeout: 10 * time.Second}
	currencyApiUrl := "https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange"
	currencyHttpClient := config.NewHttpClient(currencyApiClient, currencyApiUrl)

	purchaseRepository := repositories.NewPurchaseRepository(db)
	purchaseService := services.NewPurchaseService(purchaseRepository, currencyHttpClient)

	v1 := app.Group("/v1")
	handlers.NewPurchaseHandler(v1.Group("/purchase"), purchaseService)

	app.Listen(":8080")
}
