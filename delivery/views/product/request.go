package product

type InsertProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stok        int    `json:"stok"`
	Image       string `json:"image"`
}

type UpdateProductRequest struct {
	ID   int `json:"id"`
	Stok int `json:"stok"`
}
type DeleteProductRequest struct {
	ID int `json:"id"`
}
