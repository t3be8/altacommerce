package product

type InsertProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
	Images      string `json:"images"`
	CategoryID  uint   `json:"category_id" validate:"required"`
}

type UpdateProductRequest struct {
	ID   int `json:"id"`
	Stok int `json:"stok"`
}
type DeleteProductRequest struct {
	ID int `json:"id"`
}
