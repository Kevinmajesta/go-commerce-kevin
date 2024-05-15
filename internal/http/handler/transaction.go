package handler

import (
	"net/http"

	"github.com/Kevinmajesta/go-commerce-kevin/internal/service"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/response"
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
