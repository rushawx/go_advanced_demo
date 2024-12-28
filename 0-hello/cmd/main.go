package main

import (
	"0-hello/configs"
	"0-hello/internal/auth"
	"0-hello/internal/link"
	"0-hello/internal/user"
	"0-hello/pkg/db"
	"0-hello/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	db := db.NewDb(conf)
	fmt.Println(conf.Db.Dsn)

	router := http.NewServeMux()

	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)

	authService := auth.NewAuthService(userRepository)

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf, AuthService: authService})
	link.NewLinkHandler(router, link.LinkHandlerDeps{LinkRepository: linkRepository})

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
