package api

import (
	"inventory-app/internal/api/handlers"
	"inventory-app/internal/api/middleware"

	"github.com/gorilla/mux"
)

// SetupRoutes configures all the routes for the API
func SetupRoutes(
	router *mux.Router,
	authMiddleware *middleware.AuthMiddleware,
	authHandler *handlers.AuthHandler,
	productHandler *handlers.ProductHandler,
) {
	// Public routes
	router.HandleFunc("/api/v1/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/api/v1/register", authHandler.Register).Methods("POST")

	// Protected routes
	protected := router.PathPrefix("/api/v1").Subrouter()
	protected.Use(authMiddleware.Authenticate)

	// Product routes
	protected.HandleFunc("/products", productHandler.ListProducts).Methods("GET")
	protected.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	protected.HandleFunc("/products/{id}", productHandler.GetProduct).Methods("GET")
	protected.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	protected.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")
	protected.HandleFunc("/export/products", productHandler.ExportProductsCSV).Methods("GET")
	protected.HandleFunc("/products/{id}/barcode", productHandler.GenerateProductBarcode).Methods("GET")
}
