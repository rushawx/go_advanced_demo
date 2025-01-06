package auth

import (
	"0-hello/configs"
	"0-hello/pkg/jwt"
	"0-hello/pkg/request"
	"0-hello/pkg/response"
	"fmt"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
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

		email, err := handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			response.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}
		fmt.Println(email)

		j := jwt.NewJWT(handler.Config.Auth.Secret)
		token, err := j.Create(jwt.JWTData{Email: email})
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := LoginResponse{
			Token: token,
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

		email, err := handler.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			response.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}
		fmt.Println(email)

		j := jwt.NewJWT(handler.Config.Auth.Secret)
		token, err := j.Create(jwt.JWTData{Email: email})
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := RegisterResponse{
			Token: token,
		}
		response.Json(w, resp, http.StatusCreated)
	}
}
