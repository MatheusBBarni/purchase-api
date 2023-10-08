package domain

import "time"

type Purchase struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Description    string    `json:"description"`
	PurchaseAmount float64   `json:"purchase_amount"`
	Transaction    time.Time `json:"transaction" gorm:"default:current_timestamp"`
}
