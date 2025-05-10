package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"expense-api/models"
	"expense-api/middlewares"

	"gorm.io/gorm"
)

func CreateExpense(db *gorm.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var expense models.Expense
		if err := json.NewDecoder((r.Body)).Decode(&expense); err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			return
		}
		userID := r.Context().Value(middlewares.UserIdKey).(uint)
		expense.UserID = userID
		expense.CreatedAt = time.Now()

		if err := db.Create(&expense).Error; err != nil {
			http.Error(w, "Error Creating Expense", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(expense)
		return;
	}
}

func GetExpenses(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expenses []models.Expense
		userID := r.Context().Value(middlewares.UserIdKey).(uint)
		role := r.Context().Value(middlewares.RoleKey).(string)

		if role == "admin" {
			if err := db.Find(&expenses).Error; err != nil {
				http.Error(w, "Error Retrieving Expenses", http.StatusInternalServerError)
				return
			}else {
				db.Where("user_id = ?", userID).Find(&expenses)
			}

			json.NewEncoder(w).Encode(expenses)
			return
		}
	}
}