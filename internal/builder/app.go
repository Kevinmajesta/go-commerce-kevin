package builder

import (
	"github.com/Kevinmajesta/go-commerce-kevin/internal/http/handler"
	"github.com/Kevinmajesta/go-commerce-kevin/internal/http/router"
	"github.com/Kevinmajesta/go-commerce-kevin/internal/repository"
	"github.com/Kevinmajesta/go-commerce-kevin/internal/service"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/route"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/token"
	"gorm.io/gorm"
)

func BuildAppPublicRoutes(db *gorm.DB, tokenUseCase token.TokenUseCase) []*route.Route {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, tokenUseCase)
	userHandler := handler.NewUserHandler(userService)
	return router.AppPublicRoutes(userHandler)
}

func BuildAppPrivateRoutes(db *gorm.DB) []*route.Route {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, nil)
	userHandler := handler.NewUserHandler(userService)
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	transactionRepository := repository.NewTransactionRepository(db) 
	transactionService := service.NewTransactionService(transactionRepository) 
	transactionHandler := handler.NewTransactionHandler(transactionService) 
	return router.AppPrivateRoutes(userHandler, *productHandler, *transactionHandler) 
}
