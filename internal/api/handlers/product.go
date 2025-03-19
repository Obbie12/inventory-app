package handlers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"image/png"
	"net/http"
	"time"

	"inventory-app/internal/models"
	"inventory-app/internal/services"
	"inventory-app/internal/utils"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/gorilla/mux"
)

// ProductHandler handles HTTP requests for products
type ProductHandler struct {
	productService *services.ProductService
	validator      *utils.Validator
}

// NewProductHandler creates a new product handler
func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		validator:      utils.NewValidator(),
	}
}

// CreateProduct handles the creation of a new product
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Validate the product data
	if err := h.validator.Validate(product); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation error", err)
		return
	}

	// Get user ID from context
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to get user ID", nil)
		return
	}

	// Create the product
	createdProduct, err := h.productService.CreateProduct(product, userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create product", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, createdProduct)
}

// GetProduct handles retrieving a product by ID
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// Get product ID from URL parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Get the product
	product, err := h.productService.GetProductByID(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Product not found", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, product)
}

// ListProducts handles retrieving a list of products with optional filtering
func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters for filtering
	status := r.URL.Query().Get("status")
	lowStock := r.URL.Query().Get("low_stock") == "true"

	var filter *models.ProductFilter
	if status != "" || lowStock {
		filter = &models.ProductFilter{
			Status:   models.ProductStatus(status),
			LowStock: lowStock,
		}
	}

	// Get products
	products, err := h.productService.ListProducts(filter)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve products", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, products)
}

// UpdateProduct handles updating a product
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Get product ID from URL parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Parse the request body
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Set the product ID from the URL parameter
	product.ID = id

	// Validate the product data
	if err := h.validator.Validate(product); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation error", err)
		return
	}

	// Get user ID from context
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to get user ID", nil)
		return
	}

	// Update the product
	err := h.productService.UpdateProduct(product, userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update product", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Product updated successfully"})
}

// DeleteProduct handles deleting a product
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Get product ID from URL parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete the product
	err := h.productService.DeleteProduct(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete product", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Product deleted successfully"})
}

func (h *ProductHandler) ExportProductsCSV(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters for filtering (same as in ListProducts)
	status := r.URL.Query().Get("status")
	lowStock := r.URL.Query().Get("low_stock") == "true"

	var filter *models.ProductFilter
	if status != "" || lowStock {
		filter = &models.ProductFilter{
			Status:   models.ProductStatus(status),
			LowStock: lowStock,
		}
	}

	// Get products
	products, err := h.productService.ListProducts(filter)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve products", err)
		return
	}

	// Set headers for CSV download
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=products.csv")

	// Create CSV writer
	csvWriter := csv.NewWriter(w)
	defer csvWriter.Flush()

	// Write CSV header
	headers := []string{"ID", "Name", "SKU", "Quantity", "Status", "Created At", "Updated At"}
	if err := csvWriter.Write(headers); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to write CSV header", err)
		return
	}

	// Write product data
	for _, product := range products {
		row := []string{
			product.ID,
			product.ProductName,
			product.SKU,
			fmt.Sprintf("%d", product.Quantity),
			string(product.Status),
			product.CreatedAt.Format(time.RFC3339),
			product.UpdatedAt.Format(time.RFC3339),
		}
		if err := csvWriter.Write(row); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to write CSV data", err)
			return
		}
	}
}

func (h *ProductHandler) GenerateProductBarcode(w http.ResponseWriter, r *http.Request) {
	// Get product ID from URL parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Get the product
	product, err := h.productService.GetProductByID(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Product not found", err)
		return
	}

	// Check if SKU is available
	if product.SKU == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Product has no SKU for barcode generation", nil)
		return
	}

	// Generate Code128 barcode (commonly used for product barcodes)
	barcodeData, err := code128.Encode(product.SKU)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate barcode", err)
		return
	}

	// Scale the barcode to reasonable dimensions (300x80 is a common size)
	barcode, err := barcode.Scale(barcodeData, 300, 80)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to scale barcode", err)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s-barcode.png", product.SKU))

	// Encode and send the barcode as PNG image
	if err := png.Encode(w, barcode); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to encode barcode image", err)
		return
	}
}
