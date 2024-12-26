package auth

import (
	"0-hello/configs"
	"0-hello/pkg/request"
	"0-hello/pkg/response"
	"fmt"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login")

		body, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(*body)

		resp := LoginResponse{
			Token: handler.Config.Auth.Secret,
		}
		response.Json(w, resp, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")

		body, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(*body)

		resp := RegisterResponse{
			Token: handler.Config.Auth.Secret,
		}
		response.Json(w, resp, http.StatusOK)
	}
}
