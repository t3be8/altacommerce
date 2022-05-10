package cart

type InsertCartRequest struct {
	TotalQty int `json:"totalqty"`
}

type UpdateCartRequest struct {
	ID   int `json:"id"`
	Stok int `json:"stok"`
}
type DeleteCartRequest struct {
	ID int `json:"id"`
}
