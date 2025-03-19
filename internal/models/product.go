package models

import (
	"time"
)

// ProductStatus represents the status of a product
type ProductStatus string

const (
	StatusActive      ProductStatus = "active"
	StatusInactive    ProductStatus = "inactive"
	StatusDiscontinued ProductStatus = "discontinued"
)

// Product represents a product in the inventory
type Product struct {
	ID         string        `json:"id"`
	ProductName string       `json:"product_name" validate:"required"`
	SKU        string        `json:"sku" validate:"required"`
	Quantity   int           `json:"quantity" validate:"gte=0"`
	Location   string        `json:"location"`
	Status     ProductStatus `json:"status"`
	CreatedAt  time.Time     `json:"created_at"`
	CreatedBy  string        `json:"created_by"`
	UpdatedAt  time.Time     `json:"updated_at"`
	UpdatedBy  string        `json:"updated_by"`
}

// ProductFilter represents filters for querying products
type ProductFilter struct {
	Status   ProductStatus `json:"status"`
	LowStock bool          `json:"low_stock"`
}