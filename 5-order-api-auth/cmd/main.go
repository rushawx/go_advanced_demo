package main

import (
	"5-order-api-auth/configs"
	"5-order-api-auth/internal/auth"
	"5-order-api-auth/internal/product"
	"5-order-api-auth/internal/user"
	"5-order-api-auth/pkg/db"
	"5-order-api-auth/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	db := db.NewDb(conf)
	fmt.Println(conf.Db.Dsn)

	productRepository := product.NewProductRepository(db)
	userRepository := user.NewUserRepository(db)
	sessionRepository := user.NewSessionRepository(db)

	router := http.NewServeMux()

	authService := auth.NewAuthService(userRepository, sessionRepository)

	product.NewProductHandler(router, product.ProductHandlerDeps{ProductRepository: productRepository})
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf, AuthService: authService})

	stack := middleware.Chain(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8080")
	server.ListenAndServe()
}
