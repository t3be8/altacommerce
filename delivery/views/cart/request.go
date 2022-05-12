package cart

type InsertCartRequest struct {
	Total string `json:"total"`
	Qty   string `json:"qty"`
}

type UpdateCartRequest struct {
	ID    int `json:"id"`
	Total int `json:"total"`
}
type DeleteCartRequest struct {
	ID int `json:"id"`
}
