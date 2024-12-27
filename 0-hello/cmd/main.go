package main

import (
	"0-hello/configs"
	"0-hello/internal/auth"
	"0-hello/internal/link"
	"0-hello/pkg/db"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	db := db.NewDb(conf)
	fmt.Println(conf.Db.Dsn)

	router := http.NewServeMux()

	linkRepository := link.NewLinkRepository(db)

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})
	link.NewLinkHandler(router, link.LinkHandlerDeps{LinkRepository: linkRepository})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
