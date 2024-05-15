package repository

import (
	"github.com/Kevinmajesta/go-commerce-kevin/internal/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAllTransaction() ([]entity.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) FindAllTransaction() ([]entity.Transaction, error) {
	transactions := make([]entity.Transaction, 0)
	if err := r.db.Find(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil
}
