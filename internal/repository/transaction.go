package repository

import (
	"github.com/Kevinmajesta/go-commerce-kevin/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactionByID(id uuid.UUID) (*entity.Transaction, error)
	FindAllTransaction() ([]entity.Transaction, error)
	CreateTransaction(transaction *entity.Transaction) (*entity.Transaction, error)
	UpdateTransaction(transaction *entity.Transaction) (*entity.Transaction, error)
	DeleteTransaction(transaction *entity.Transaction) (bool, error)
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

func (r *transactionRepository) FindTransactionByID(id uuid.UUID) (*entity.Transaction, error) {
	transactions := new(entity.Transaction)
	if err := r.db.Where("id = ?", id).Take(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil
}
func (r *transactionRepository) CreateTransaction(transactions *entity.Transaction) (*entity.Transaction, error) {
	if err := r.db.Create(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *transactionRepository) UpdateTransaction(transactions *entity.Transaction) (*entity.Transaction, error) {
	// Use map to store fields to be updated.
	fields := make(map[string]interface{})

	// Update fields only if they are not empty.
	if transactions.ProductID != "" {
		fields["product_id"] = transactions.ProductID
	}
	if transactions.Qty != "" {
		fields["qty"] = transactions.Qty
	}
	if transactions.UserID != "" {
		fields["user_id"] = transactions.UserID
	}
	if transactions.Discount != "" {
		fields["Discount"] = transactions.Discount
	}
	if transactions.IsPaid {
		fields["is_paid"] = transactions.IsPaid
	}

	// Update the database in one query.
	if err := r.db.Model(transactions).Where("id = ?", transactions.ID).Updates(fields).Error; err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionRepository) DeleteTransaction(transactions *entity.Transaction) (bool, error) {
	if err := r.db.Delete(&transactions).Error; err != nil {
		return false, err
	}
	return true, nil
}
