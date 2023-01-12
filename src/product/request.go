package product

type CreateProductRequest struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Image      string  `json:"image"`
	IsInStock  bool    `json:"is_in_stock"`
	CategoryId uint    `json:"category_id"`
}

type UpdateProductRequest struct {
	ProductId  uint    `json:"product_id" binding:"required"`
	Name       string  `json:"name" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	Image      string  `json:"image" binding:"required"`
	IsInStock  bool    `json:"is_in_stock" binding:"required"`
	CategoryId uint    `json:"category_id" binding:"required"`
}
