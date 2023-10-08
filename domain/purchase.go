package domain

import "time"

type Purchase struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Transaction time.Time `json:"transaction" gorm:"default:current_timestamp"`
}
