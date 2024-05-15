package binder

type ProductCreateRequest struct {
	Name     string `json:"name" validate:"required,email"`
	Price    string `json:"price" validate:"required"`
	Category string `json:"category" validate:"required"`
}

type ProductUpdateRequest struct {
	ID       string `param:"id" validate:"required"`
	Name     string `json:"name" validate:"required,email"`
	Price    string `json:"price" validate:"required"`
	Category string `json:"category" validate:"required"`
}

type ProductDeleteRequest struct {
	ID string `param:"id" validate:"required"`
}
