package service

import (
	"github.com/Kevinmajesta/go-commerce-kevin/internal/entity"
	"github.com/Kevinmajesta/go-commerce-kevin/internal/repository"
	"github.com/google/uuid"
)

type TransactionService interface {
	FindAllTransaction() ([]entity.Transaction, error)
	CreateTransaction(transaction *entity.Transaction) (*entity.Transaction, error)
	UpdateTransaction(transaction *entity.Transaction) (*entity.Transaction, error)
	DeleteTransaction(id uuid.UUID) (bool, error)
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

func (s *transactionService) CreateTransaction(transactions *entity.Transaction) (*entity.Transaction, error) {
	return s.transactionRepository.CreateTransaction(transactions)
}

func (s *transactionService) UpdateTransaction(transactions *entity.Transaction) (*entity.Transaction, error) {
	return s.transactionRepository.UpdateTransaction(transactions)
}

func (s *transactionService) DeleteTransaction(id uuid.UUID) (bool, error) {
	transactions, err := s.transactionRepository.FindTransactionByID(id)
	if err != nil {
		return false, err
	}

	return s.transactionRepository.DeleteTransaction(transactions)
}
