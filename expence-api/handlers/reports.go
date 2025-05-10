package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"expense-api/models"
	"expense-api/middlewares"
	"gorm.io/gorm"
)

func FilteredExpenses(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(middlewares.UserIdKey).(uint)
		role := r.Context().Value(middlewares.RoleKey).(string)

		var expenses []models.Expense
		query := db.Model(&models.Expense{})

		if role != "admin" {
			query = query.Where("user_id = ?", userID)
		}

		if start := r.URL.Query().Get("start"); start != "" {
			if t, err := time.Parse("2006-01-02", start); err == nil {
				query = query.Where("date >= ?", t)
			}
		}

		if end := r.URL.Query().Get("end"); end != "" {
			if t, err := time.Parse("2006-01-02", end); err == nil {
				query = query.Where("date <= ?", t)
			}
		}

		if category := r.URL.Query().Get("category"); category != "" {
			query = query.Where("category = ?", category)
		}

		if min := r.URL.Query().Get("min_amount"); min != "" {
			if amount, err := strconv.ParseFloat(min, 64); err == nil {
				query = query.Where("amount >= ?", amount)
			}
		}

		if max := r.URL.Query().Get("max_amount"); max != "" {
			if amount, err := strconv.ParseFloat(max, 64); err == nil {
				query = query.Where("amount <= ?", amount)
			}
		}

		if err := query.Find(&expenses).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(expenses)
	}
}



func ExpenseSummary(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(middlewares.UserIdKey).(uint)
		role := r.Context().Value(middlewares.RoleKey).(string)

		type CategoryTotal struct {
			Category string  `json:"category"`
			Total    float64 `json:"total"`
		}

		var result []CategoryTotal

		query := db.Model(&models.Expense{}).Select("category, SUM(amount) as total").Group("category")
		if role != "admin" {
			query = query.Where("user_id = ?", userID)
		}
		if err := query.Scan(&result).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(result)
	}
}