package services

import (
	"inventory-app/internal/models"
	"inventory-app/internal/repository"
)

// ProductService handles product business logic
type ProductService struct {
	productRepo *repository.ProductRepository
}

// NewProductService creates a new product service
func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

// CreateProduct adds a new product
func (s *ProductService) CreateProduct(product models.Product, userID string) (models.Product, error) {
	return s.productRepo.Create(product, userID)
}

// GetProductByID retrieves a product by its ID
func (s *ProductService) GetProductByID(id string) (models.Product, error) {
	return s.productRepo.GetByID(id)
}

// ListProducts retrieves products with optional filtering
func (s *ProductService) ListProducts(filter *models.ProductFilter) ([]models.Product, error) {
	return s.productRepo.ListProducts(filter)
}

// UpdateProduct updates an existing product
func (s *ProductService) UpdateProduct(product models.Product, userID string) error {
	return s.productRepo.Update(product, userID)
}

// DeleteProduct removes a product
func (s *ProductService) DeleteProduct(id string) error {
	return s.productRepo.Delete(id)
}
