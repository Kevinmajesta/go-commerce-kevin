package builder

import (
	"github.com/Kevinmajesta/go-commerce-kevin/internal/http/router"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/route"
)

func BuildAuthPublicRoutes() []*route.Route {
	return router.AuthPublicRoutes()
}

func BuildAuthPrivateRoutes() []*route.Route {
	return router.AuthPublicRoutes()
}