package service

import (
	"github.com/Kevinmajesta/go-commerce-kevin/internal/entity"
	"github.com/Kevinmajesta/go-commerce-kevin/internal/repository"
)

type ProductService interface {
	FindAllProduct() ([]entity.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepository: productRepo}
}

func (s *productService) FindAllProduct() ([]entity.Product, error) {
	return s.productRepository.FindAllProduct()
}
