package middlewares

import (
	"context"
	"net/http"
	"strings"
	"expense-api/utils"
)

type contextKey string

const UserIdKey contextKey = "user_id"
const RoleKey contextKey = "role"

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}
		calims, err := utils.ParseJWT(strings.TrimPrefix(token, "Bearer "))
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIdKey, calims.UserID)
		ctx = context.WithValue(ctx, RoleKey, calims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

