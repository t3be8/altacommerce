package cart

type InsertCartRequest struct {
	Total string `json:"total"`
	Qty   string `json:"qty"`
}

type UpdateCartRequest struct {
	ID   int `json:"id"`
	Stok int `json:"stok"`
}
type DeleteCartRequest struct {
	ID int `json:"id"`
}
