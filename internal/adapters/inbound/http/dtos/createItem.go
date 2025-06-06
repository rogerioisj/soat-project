package dtos

type CreateItemDto struct {
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"required,min=10,max=200"`
	Price       int64  `json:"price" validate:"required,min=0"`
	Category    string `json:"category" validate:"required"`
}
