package builder

import (
	"github.com/Kevinmajesta/go-commerce-kevin/internal/http/handler"
	"github.com/Kevinmajesta/go-commerce-kevin/internal/http/router"
	"github.com/Kevinmajesta/go-commerce-kevin/internal/repository"
	"github.com/Kevinmajesta/go-commerce-kevin/internal/service"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/cache"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/encrypt"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/route"
	"github.com/Kevinmajesta/go-commerce-kevin/pkg/token"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func BuildAppPublicRoutes(db *gorm.DB, tokenUseCase token.TokenUseCase, encryptTool encrypt.EncryptTool) []*route.Route {
	userRepository := repository.NewUserRepository(db, nil)
	userService := service.NewUserService(userRepository, tokenUseCase, encryptTool)
	userHandler := handler.NewUserHandler(userService)
	productRepository := repository.NewProductRepository(db, nil)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	return router.AppPublicRoutes(userHandler, *productHandler)
}

func BuildAppPrivateRoutes(db *gorm.DB, redisDB *redis.Client, encryptTool encrypt.EncryptTool) []*route.Route {
	cacheable := cache.NewCacheable(redisDB)
	userRepository := repository.NewUserRepository(db, cacheable)
	userService := service.NewUserService(userRepository, nil, encryptTool)
	userHandler := handler.NewUserHandler(userService)
	productRepository := repository.NewProductRepository(db, cacheable)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	return router.AppPrivateRoutes(userHandler, *productHandler, *transactionHandler)
}
