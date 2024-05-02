package main

import (
	"github.com/Kevinmajesta/go-commerce-kevin/internal/builder"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/server"
)

func main() {
	publicRoutes := builder.BuildAuthPublicRoutes()
	privateRoutes := builder.BuildAuthPrivateRoutes()

	srv := server.NewServer("auth", publicRoutes, privateRoutes, "")
	srv.Run()
}
