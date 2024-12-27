package main

import (
	"4-order-api/configs"
	"4-order-api/pkg/db"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	_ = db.NewDb(conf)
	fmt.Println(conf.Db.Dsn)

	router := http.NewServeMux()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8080")
	server.ListenAndServe()
}
