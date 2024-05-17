package repository

import (
	"encoding/json"
	"time"

	"github.com/Kevinmajesta/go-commerce-kevin/internal/entity"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/cache"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProductByID(id uuid.UUID) (*entity.Product, error)
	FindAllProduct() ([]entity.Product, error)
	CreateProduct(product *entity.Product) (*entity.Product, error)
	UpdateProduct(product *entity.Product) (*entity.Product, error)
	DeleteProduct(product *entity.Product) (bool, error)
}

type productRepository struct {
	db        *gorm.DB
	cacheable cache.Cacheable
}

func NewProductRepository(db *gorm.DB, cacheable cache.Cacheable) ProductRepository {
	return &productRepository{db: db, cacheable: cacheable}
}

func (r *productRepository) FindAllProduct() ([]entity.Product, error) {
	products := make([]entity.Product, 0)

	key := "FindAllProducts"

	data, _ := r.cacheable.Get(key)
	if data == "" {
		if err := r.db.Find(&products).Error; err != nil {
			return products, err
		}
		marshalledproducts, _ := json.Marshal(products)
		err := r.cacheable.Set(key, marshalledproducts, 5*time.Minute)
		if err != nil {
			return products, err
		}
	} else {
		// Data ditemukan di Redis, unmarshal data ke products
		err := json.Unmarshal([]byte(data), &products)
		if err != nil {
			return products, err
		}
	}
	return products, nil
}

func (r *productRepository) FindProductByID(id uuid.UUID) (*entity.Product, error) {
	products := new(entity.Product)
	if err := r.db.Where("id = ?", id).Take(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func (r *productRepository) CreateProduct(products *entity.Product) (*entity.Product, error) {
	if err := r.db.Create(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func (r *productRepository) UpdateProduct(products *entity.Product) (*entity.Product, error) {
	// Use map to store fields to be updated.
	fields := make(map[string]interface{})

	// Update fields only if they are not empty.
	if products.Name != "" {
		fields["name"] = products.Name
	}
	if products.Price != "" {
		fields["price"] = products.Price
	}
	if products.Category != "" {
		fields["category"] = products.Category
	}

	// Update the database in one query.
	if err := r.db.Model(products).Where("id = ?", products.ID).Updates(fields).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (r *productRepository) DeleteProduct(products *entity.Product) (bool, error) {
	if err := r.db.Delete(&products).Error; err != nil {
		return false, err
	}
	return true, nil
}
