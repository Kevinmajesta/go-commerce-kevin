package binder

type TransactionCreateRequest struct {
	ProductID string `json:"product_id" validate:"required"`
	Qty       string `json:"qty" validate:"required"`
	UserID    string `json:"user_id" validate:"required"`
	Discount  string `json:"discount"`
	IsPaid    bool   `json:"is_paid" validate:"required"`
}

type TransactionUpdateRequest struct {
	ID        string `param:"id" validate:"required"`
	ProductID string `json:"product_id" validate:"required"`
	Qty       string `json:"qty" validate:"required"`
	UserID    string `json:"user_id" validate:"required"`
	Discount  string `json:"discount" `
	IsPaid    bool   `json:"is_paid" validate:"required"`
}

type TransactionDeleteRequest struct {
	ID string `param:"id" validate:"required"`
}
