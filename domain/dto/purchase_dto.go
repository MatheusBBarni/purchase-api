package dto

import "github.com/matheusBBarni/purchase-api/domain"

type PurchaseDto struct {
	Description string  `json:"description" validate:"required,max=55"`
	Amount      float64 `json:"amount" validate:"required,min=0"`
}

func (purchaseDto PurchaseDto) ToDomain() *domain.Purchase {
	return &domain.Purchase{
		Description: purchaseDto.Description,
		Amount:      purchaseDto.Amount,
	}
}
