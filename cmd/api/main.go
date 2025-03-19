package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	handler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"inventory-app/internal/api"
	"inventory-app/internal/api/handlers"
	"inventory-app/internal/api/middleware"
	"inventory-app/internal/config"
	"inventory-app/internal/repository"
	"inventory-app/internal/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Set up database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// // Configure connection pool
	// db.SetMaxIdleConns(10)
	// db.SetMaxOpenConns(100)
	// db.SetConnMaxLifetime(time.Hour)

	// Test database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Connected to database successfully")

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	productService := services.NewProductService(productRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	productHandler := handlers.NewProductHandler(productService)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(authService)

	// Set up router
	router := mux.NewRouter()
	api.SetupRoutes(router, authMiddleware, authHandler, productHandler)

	corsHandler := handler.CORS(
		handler.AllowedOrigins([]string{"*"}),
		handler.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}),
		handler.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handler.AllowCredentials(),
	)

	// Start the server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, corsHandler(router)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
