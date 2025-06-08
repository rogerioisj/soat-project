package dtos

type UpdateItemDto struct {
	Name        string `json:"name" validate:"min=3,max=50"`
	Description string `json:"description" validate:"min=10,max=200"`
	Price       int64  `json:"price" validate:"min=0"`
	Category    string `json:"category"`
}
