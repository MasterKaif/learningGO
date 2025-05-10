package main

import (
	"net/http"

	"expense-api/handlers"
	"expense-api/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	InitDB()

	// Create a new router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/register", handlers.RegisterUser(DB)).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser(DB)).Methods("POST")

	//Expense routes
	r.HandleFunc("/expenses", middlewares.AuthMiddleware(handlers.CreateExpense(DB))).Methods("POST")
	r.HandleFunc("/expenses", middlewares.AuthMiddleware(handlers.GetExpenses(DB))).Methods("GET")

	//Report and Summary
	r.Handle("/expenses/report", middlewares.AuthMiddleware(handlers.FilteredExpenses(DB))).Methods("GET")
	r.Handle("/expenses/summary", middlewares.AuthMiddleware(handlers.ExpenseSummary(DB))).Methods("GET")


	// Health check route
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")


	// Start the server
	http.ListenAndServe(":8080", r)
}
