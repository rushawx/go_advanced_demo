package main

import (
	"0-hello/configs"
	"0-hello/internal/auth"
	"0-hello/internal/link"
	"0-hello/internal/stat"
	"0-hello/internal/user"
	"0-hello/pkg/db"
	"0-hello/pkg/event"
	"0-hello/pkg/middleware"
	"fmt"
	"net/http"
)

func App() http.Handler {
	conf := configs.LoadConfig()

	db := db.NewDb(conf)
	fmt.Println(conf.Db.Dsn)

	router := http.NewServeMux()

	eventBus := event.NewEventBus()

	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(stat.StatServiceDeps{EventBus: eventBus, StatRepository: statRepository})

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf, AuthService: authService})
	link.NewLinkHandler(router, link.LinkHandlerDeps{EventBus: eventBus, LinkRepository: linkRepository, Config: conf})
	stat.NewStatHandler(router, stat.StatHandlerDeps{StatRepository: statRepository, Config: conf})

	go statService.AddClick(1)

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	return stack(router)
}

func main() {
	app := App()

	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
