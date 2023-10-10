package dto

import (
	"github.com/matheusBBarni/purchase-api/domain"
	"github.com/matheusBBarni/purchase-api/utils"
)

type PurchaseDto struct {
	Description     string  `json:"description" validate:"required,max=55"`
	Amount          float64 `json:"amount" validate:"required,min=0"`
	TransactionDate string  `json:"transaction_date" validate:"required"`
}

func (purchaseDto PurchaseDto) ToDomain() (*domain.Purchase, error) {
	transformedDate, err := utils.ConvertStringToDate(purchaseDto.TransactionDate)
	if err != nil {
		return &domain.Purchase{}, err
	}

	return &domain.Purchase{
		Description:     purchaseDto.Description,
		Amount:          purchaseDto.Amount,
		TransactionDate: transformedDate,
	}, nil
}
