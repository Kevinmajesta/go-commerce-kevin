package entity

import "github.com/google/uuid"

type Transaction struct {
	ID        uuid.UUID `json:"id"`
	ProductID string    `json:"product_id"`
	Qty       string    `json:"qty"`
	UserID    string    `json:"user_id"`
	Discount  string    `json:"discount"`
	IsPaid    bool      `json:"is_paid"`
	Auditable
}

func NewTransaction(product_id, qty, user_id, discount string, is_paid bool) *Transaction {
	return &Transaction{
		ID:        uuid.New(),
		ProductID: product_id,
		Qty:       qty,
		UserID:    user_id,
		Discount:  discount,
		IsPaid:    is_paid,
		Auditable: NewAuditable(),
	}
}

func UpdateTransaction(id uuid.UUID, product_id, qty, user_id, discount string, is_paid bool) *Transaction {
	return &Transaction{
		ID:        uuid.New(),
		ProductID: product_id,
		Qty:       qty,
		UserID:    user_id,
		Discount:  discount,
		IsPaid:    is_paid,
		Auditable: UpdateAuditable(),
	}
}
