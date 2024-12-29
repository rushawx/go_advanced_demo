package main

import (
	"4-order-api/configs"
	"4-order-api/internal/product"
	"4-order-api/pkg/db"
	"4-order-api/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	db := db.NewDb(conf)
	fmt.Println(conf.Db.Dsn)

	productRepository := product.NewProductRepository(db)

	router := http.NewServeMux()

	product.NewProductHandler(router, product.ProductHandlerDeps{ProductRepository: productRepository})

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
