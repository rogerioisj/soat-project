package dtos

type CreateOrderRequest struct {
	UserID   string    `json:"user_id" validate:"required,int"`
	Products []Product `json:"products" validate:"required,dive,required"`
}

type Product struct {
	ID       string `json:"id" validate:"required,int"`
	Quantity int    `json:"quantity" validate:"required,int"`
}
