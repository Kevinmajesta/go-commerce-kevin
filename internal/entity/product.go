package entity

import "github.com/google/uuid"

type Product struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    string    `json:"price"`
	Category string    `json:"category"`
	Auditable
}

func NewProduct(name, price, category string) *Product {
	return &Product{
		ID:        uuid.New(),
		Name:      name,
		Price:     price,
		Category:  category,
		Auditable: NewAuditable(),
	}
}

func UpdateProduct(id uuid.UUID, name, price, category string) *Product {
	return &Product{
		ID:        id,
		Name:      name,
		Price:     price,
		Category:  category,
		Auditable: UpdateAuditable(),
	}
}
