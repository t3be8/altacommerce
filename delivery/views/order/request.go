package order

type OrderRequest struct {
	ShipmentID uint   `json:"shipment_id" validate:"required"`
	Address    string `json:"address" validate:"required"`
}

type UpdateStatusOrderRequest struct {
	Status string `json:"status" validate:"required"`
}
