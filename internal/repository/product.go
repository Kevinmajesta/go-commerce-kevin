package repository

import (
	"github.com/Kevinmajesta/go-commerce-kevin/internal/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAllProduct() ([]entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindAllProduct() ([]entity.Product, error) {
	products := make([]entity.Product, 0)
	if err := r.db.Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}
