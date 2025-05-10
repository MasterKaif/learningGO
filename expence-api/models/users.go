package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name		string	`json:"name" validate:"required"`
	Email		string	`gorm:"uniqueIndex" json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required,min=6"`
	Role 		string	`json:"role" validate:"required,oneof=admin employee"`	
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
