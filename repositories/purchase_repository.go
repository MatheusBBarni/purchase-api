package repositories

import (
	"github.com/matheusBBarni/purchase-api/domain"
	"gorm.io/gorm"
)

type PurchaseRepository struct {
	db *gorm.DB
}

func NewPurchaseRepository(db *gorm.DB) *PurchaseRepository {
	return &PurchaseRepository{
		db: db,
	}
}

func (purchaseRepository *PurchaseRepository) GetById(id uint) *domain.Purchase {
	var purchase *domain.Purchase

	purchaseRepository.db.First(&purchase, id)

	return purchase
}

func (purchaseRepository *PurchaseRepository) GetAll() []*domain.Purchase {
	var purchases []*domain.Purchase

	purchaseRepository.db.Find(&purchases)

	return purchases
}

func (purchaseRepository *PurchaseRepository) Save(purchase *domain.Purchase) (*domain.Purchase, error) {
	err := purchaseRepository.db.Create(&purchase).Error

	if err != nil {
		return nil, err
	}

	return purchase, nil
}

func (purchaseRepository *PurchaseRepository) Update(purchase *domain.Purchase) (*domain.Purchase, error) {
	err := purchaseRepository.db.Model(&purchase).UpdateColumns(purchase).Error

	if err != nil {
		return nil, err
	}

	return purchase, nil
}

func (purchaseRepository *PurchaseRepository) Delete(id uint) error {
	var purchase *domain.Purchase

	err := purchaseRepository.db.Delete(&purchase, id).Error

	if err != nil {
		return err
	}

	return nil
}
