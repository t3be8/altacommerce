package cart

type InsertCartRequest struct {
	Price float64 `json:"price"`
	Qty   int     `json:"qty"`
}

type UpdateCartRequest struct {
	ID    int     `json:"id"`
	Price float64 `json:"price"`
}
type DeleteCartRequest struct {
	ID int `json:"id"`
}
