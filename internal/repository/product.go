package repository

import (
	"database/sql"
	"fmt"
	"time"

	"inventory-app/internal/models"

	"github.com/google/uuid"
)

// ProductRepository handles all database operations for products
type ProductRepository struct {
	db *sql.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Create adds a new product to the database
func (r *ProductRepository) Create(product models.Product, userID string) (models.Product, error) {
	// Generate UUID for product ID
	product.ID = uuid.New().String()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	product.CreatedBy = userID
	product.UpdatedBy = userID

	// If status is empty, set default value
	if product.Status == "" {
		product.Status = models.StatusActive
	}

	query := `
		INSERT INTO products (id, product_name, sku, quantity, location, status, created_at, created_by, updated_at, updated_by)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(
		query,
		product.ID,
		product.ProductName,
		product.SKU,
		product.Quantity,
		product.Location,
		product.Status,
		product.CreatedAt,
		product.CreatedBy,
		product.UpdatedAt,
		product.UpdatedBy,
	)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

// GetByID retrieves a product by its ID
func (r *ProductRepository) GetByID(id string) (models.Product, error) {
	var product models.Product
	query := `
		SELECT id, product_name, sku, quantity, location, status, created_at, created_by, updated_at, updated_by
		FROM products
		WHERE id = ?
	`
	err := r.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.ProductName,
		&product.SKU,
		&product.Quantity,
		&product.Location,
		&product.Status,
		&product.CreatedAt,
		&product.CreatedBy,
		&product.UpdatedAt,
		&product.UpdatedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Product{}, fmt.Errorf("product with ID %s not found", id)
		}
		return models.Product{}, err
	}

	return product, nil
}

// ListProducts retrieves a list of products with optional filtering
func (r *ProductRepository) ListProducts(filter *models.ProductFilter) ([]models.Product, error) {
	var products []models.Product

	query := `
		SELECT id, product_name, sku, quantity, location, status, created_at, created_by, updated_at, updated_by
		FROM products
		WHERE 1=1
	`
	args := []interface{}{}

	// Apply filters if provided
	if filter != nil {
		if filter.Status != "" {
			query += " AND status = ?"
			args = append(args, filter.Status)
		}

		if filter.LowStock {
			query += " AND quantity < 10" // Assuming low stock is less than 10 units
		}
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.ProductName,
			&product.SKU,
			&product.Quantity,
			&product.Location,
			&product.Status,
			&product.CreatedAt,
			&product.CreatedBy,
			&product.UpdatedAt,
			&product.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// Update updates an existing product
func (r *ProductRepository) Update(product models.Product, userID string) error {
	product.UpdatedAt = time.Now()
	product.UpdatedBy = userID

	query := `
		UPDATE products
		SET product_name = ?, sku = ?, quantity = ?, location = ?, status = ?, updated_at = ?, updated_by = ?
		WHERE id = ?
	`
	result, err := r.db.Exec(
		query,
		product.ProductName,
		product.SKU,
		product.Quantity,
		product.Location,
		product.Status,
		product.UpdatedAt,
		product.UpdatedBy,
		product.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product with ID %s not found", product.ID)
	}

	return nil
}

// Delete removes a product from the database
func (r *ProductRepository) Delete(id string) error {
	query := `DELETE FROM products WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product with ID %s not found", id)
	}

	return nil
}
