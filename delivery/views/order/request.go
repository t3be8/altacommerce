package order

type OrderRequest struct {
	CartID     uint   `json:"cart_id" validate:"required"`
	ShipmentID uint   `json:"shipment_id" validate:"required"`
	Address    string `json:"address" validate:"required"`
}

type UpdateStatusOrderRequest struct {
	Status string `json:"status" validate:"required"`
}
