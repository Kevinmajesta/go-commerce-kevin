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

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{userService: userService}
}

func (h *UserHandler) Login(c echo.Context) error {
	input := binder.UserLoginRequest{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "ada kesalahan input"))
	}

	user, err := h.userService.Login(input.Email, input.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "login success", user))
}

func (h *UserHandler) FindAllUser(c echo.Context) error {
	users, err := h.userService.FindAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "sukses menampilkan data user", users))
}
func (h *UserHandler) CreateUser(c echo.Context) error {
	input := binder.UserCreateRequest{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "ada kesalahan input"))
	}

	newUser := entity.NewUser(input.Email, input.Password, input.Role, input.Alamat, input.NoHp)

	user, err := h.userService.CreateUser(newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "sukses membuat user baru", user))
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	var input binder.UserUpdateRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "ada kesalahan input"))
	}

	id := uuid.MustParse(input.ID)

	inputUser := entity.UpdateUser(id, input.Email, input.Password, input.Role, input.Alamat, input.NoHp)

	updatedUser, err := h.userService.UpdateUser(inputUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "sukses update user", updatedUser))
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	var input binder.UserDeleteRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "ada kesalahan input"))
	}

	id := uuid.MustParse(input.ID)

	isDeleted, err := h.userService.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "sukses delete user", isDeleted))
}
