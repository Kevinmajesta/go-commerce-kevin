package handler

import (
	"net/http"

	"github.com/Kevinmajesta/go-commerce-kevin/internal/service"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/response"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) FindAllProduct(c echo.Context) error {
	products, err := h.productService.FindAllProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "sukses menampilkan data products", products))
}
