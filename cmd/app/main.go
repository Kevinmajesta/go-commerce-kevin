package main

import (
	"github.com/Kevinmajesta/go-commerce-kevin/configs"
	"github.com/Kevinmajesta/go-commerce-kevin/internal/builder"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/cache"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/postgres"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/server"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/token"
)

func main() {
	cfg, err := configs.NewConfig(".env")
	checkError(err)

	db, err := postgres.InitPostgres(&cfg.Postgres)
	checkError(err)

	redisDB := cache.InitCache(&cfg.Redis)

	tokenUseCase := token.NewTokenUseCase(cfg.JWT.SecretKey)

	publicRoutes := builder.BuildAppPublicRoutes(db, tokenUseCase)
	privateRoutes := builder.BuildAppPrivateRoutes(db, redisDB)

	srv := server.NewServer("app", publicRoutes, privateRoutes, cfg.JWT.SecretKey)
	srv.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
