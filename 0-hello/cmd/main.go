package main

import (
	"0-hello/configs"
	"0-hello/internal/auth"
	"0-hello/internal/hello"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	fmt.Println(conf.Db.Dsn)

	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
