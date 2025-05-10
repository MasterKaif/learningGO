package models

import "time"

type Expense struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	Title       string    `json:"title" validate:"required" gorm:"not null"`
	Amount      float64   `json:"amount" validate:"required,gt=0" gorm:"not null"`
	Category    string    `json:"category" validate:"required" gorm:"not null"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}
