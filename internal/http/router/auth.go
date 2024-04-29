package router

import (
	"net/http"

	"github.com/Kevinmajesta/go-commerce-kevin/pkg/route"
	"github.com/labstack/echo/v4"
)

func PublicRoutes() []*route.Route {
	return []*route.Route{
		{
			Method: http.MethodGet,
			Path:   "/login",
			Handler: func(c echo.Context) error {
				return c.JSON(http.StatusOK, "Ini adalah method login !")
			},
		},
	}
}

func PrivateRoutes() []*route.Route {
	return nil
}
