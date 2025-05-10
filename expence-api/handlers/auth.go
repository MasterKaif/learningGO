package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"expense-api/models"
	"expense-api/utils"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

// RegisterUser handles POST /register
func RegisterUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			return
		}

		err = validate.Struct(user)
		if err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			return
		}

		hashed, err := utils.HashPassword(user.Password)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		user.Password = hashed

		if err := db.Create(&user).Error; err != nil {
			if strings.Contains(err.Error(), "UNIQUE") {
				http.Error(w, "Email ALready Exists", http.StatusConflict)
				return
			}
			http.Error(w, "Error Creating User", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User Crated!"})
	}
}

// LoginUser Handles POST /login
func LoginUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds = models.LoginRequest{}
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			return
		}

		var user models.User
		if err := db.Where("email = ?", creds.Email).First(&user).Error; err != nil {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}

		if !utils.CheckPasswordHash(creds.Password, user.Password) {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
		token, err := utils.GenerateJWT(user.ID, user.Role)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}
