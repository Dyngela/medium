package product

import (
	"database/sql"
)

type GetProductByIdResponse struct {
	ProductId uint           `json:"product_id"`
	Name      string         `json:"name"`
	Price     float64        `json:"price"`
	Image     sql.NullString `json:"image"`
	IsInStock bool           `json:"is_in_stock"`
}

type GetAllProductResponse struct {
	ProductId uint           `json:"product_id"`
	Name      string         `json:"name"`
	Price     float64        `json:"price"`
	Image     sql.NullString `json:"image"`
	IsInStock bool           `json:"is_in_stock"`
}
