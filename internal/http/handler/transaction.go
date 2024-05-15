package handler

import (
	"net/http"

	"github.com/Kevinmajesta/go-commerce-kevin/internal/entity"
	"github.com/Kevinmajesta/go-commerce-kevin/internal/http/binder"
	"github.com/Kevinmajesta/go-commerce-kevin/internal/service"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/response"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) FindAllTransaction(c echo.Context) error {
	transactions, err := h.transactionService.FindAllTransaction()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "sukses menampilkan data transactions", transactions))
}

func (h *TransactionHandler) CreateTransaction(c echo.Context) error {
	input := binder.TransactionCreateRequest{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "ada kesalahan input"))
	}

	newTransaction := entity.NewTransaction(input.ProductID, input.Qty, input.UserID, input.Discount, input.IsPaid)

	transaction, err := h.transactionService.CreateTransaction(newTransaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "sukses membuat transaction baru", transaction))
}

func (h *TransactionHandler) UpdateTransaction(c echo.Context) error {
	var input binder.TransactionUpdateRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "ada kesalahan input"))
	}

	id := uuid.MustParse(input.ID)

	inputTransaction := entity.UpdateTransaction(id, input.ProductID, input.Qty, input.UserID, input.Discount, input.IsPaid)

	updatedTransaction, err := h.transactionService.UpdateTransaction(inputTransaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "sukses update transaction", updatedTransaction))
}

func (h *TransactionHandler) DeleteTransaction(c echo.Context) error {
	var input binder.TransactionDeleteRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "ada kesalahan input"))
	}

	id := uuid.MustParse(input.ID)

	isDeleted, err := h.transactionService.DeleteTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "sukses delete transaction", isDeleted))
}
