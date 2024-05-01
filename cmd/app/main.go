package main

import (
	"github.com/Kevinmajesta/go-commerce-kevin/configs"
	"github.com/Kevinmajesta/go-commerce-kevin/internal/builder"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/postgres"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/server"
)

func main() {
	cfg, err := configs.NewConfig(".env")
	checkError(err)

	db, err := postgres.InitPostgres(&cfg.Postgres)
	checkError(err)

	publicRoutes := builder.BuildAppPublicRoutes(db)
	privateRoutes := builder.BuildAppPrivateRoutes()

	srv := server.NewServer("app", publicRoutes, privateRoutes)
	srv.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
