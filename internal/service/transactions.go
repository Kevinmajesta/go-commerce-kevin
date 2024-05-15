package service

import (
	"github.com/Kevinmajesta/go-commerce-kevin/internal/entity"
	"github.com/Kevinmajesta/go-commerce-kevin/internal/repository"
)

type TransactionService interface {
	FindAllTransaction() ([]entity.Transaction, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{transactionRepository: transactionRepo}
}

func (s *transactionService) FindAllTransaction() ([]entity.Transaction, error) {
	return s.transactionRepository.FindAllTransaction()
}
