package router

import (
	"net/http"

	"github.com/Kevinmajesta/go-commerce-kevin/internal/http/handler"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/route"
)

func AppPublicRoutes(userHandler handler.UserHandler) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: userHandler.Login,
		},
	}
}

func AppPrivateRoutes(userHandler handler.UserHandler, productHandler handler.ProductHandler) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: userHandler.FindAllUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: userHandler.CreateUser,
		},
		{
			Method:  http.MethodPut,
			Path:    "/users/:id",
			Handler: userHandler.UpdateUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/:id",
			Handler: userHandler.DeleteUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/products",
			Handler: productHandler.FindAllProduct,
		},
	}
}
