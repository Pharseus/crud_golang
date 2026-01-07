package payloads

type CreateProductRequest struct {
	Name  string  `json:"name" validate:"required"`
	SKU   string  `json:"sku" validate:"required"`
	Price float64 `json:"price" validate:"required,min=0"`
	Stock int     `json:"stock" validate:"min=0"`
}

type UpdateProductRequest struct {
	Name  string   `json:"name"`
	SKU   string   `json:"sku"`
	Price *float64 `json:"price" validate:"omitempty,min=0"`
	Stock *int     `json:"stock" validate:"omitempty,min=0"`
}
