package service

import (
	"github.com/Kevinmajesta/go-commerce-kevin/internal/entity"
	"github.com/Kevinmajesta/go-commerce-kevin/internal/repository"
	"github.com/google/uuid"
)

type ProductService interface {
	FindAllProduct() ([]entity.Product, error)
	CreateProduct(product *entity.Product) (*entity.Product, error)
	UpdateProduct(product *entity.Product) (*entity.Product, error)
	DeleteProduct(id uuid.UUID) (bool, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	return &productService{
		productRepository: productRepository,
	}
}

func (s *productService) FindAllProduct() ([]entity.Product, error) {
	return s.productRepository.FindAllProduct()
}

func (s *productService) CreateProduct(products *entity.Product) (*entity.Product, error) {
	return s.productRepository.CreateProduct(products)
}

func (s *productService) UpdateProduct(products *entity.Product) (*entity.Product, error) {
	return s.productRepository.UpdateProduct(products)
}

func (s *productService) DeleteProduct(id uuid.UUID) (bool, error) {
	products, err := s.productRepository.FindProductByID(id)
	if err != nil {
		return false, err
	}

	return s.productRepository.DeleteProduct(products)
}
