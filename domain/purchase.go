package domain

import "time"

type Purchase struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	TransactionDate time.Time `json:"transaction_date"`
}
