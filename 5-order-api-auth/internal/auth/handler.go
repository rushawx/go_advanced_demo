package auth

import (
	"5-order-api-auth/configs"
	"5-order-api-auth/pkg/jwt"
	"5-order-api-auth/pkg/request"
	"5-order-api-auth/pkg/response"
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
	router.HandleFunc("POST /auth/session", handler.Session())
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Session() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login")

		body, err := request.HandleBody[SessionRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(*body)

		sessionId, err := handler.AuthService.GetSessionId(body.Phone)
		if err != nil {
			response.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}

		fmt.Println(sessionId)

		resp := SessionResponse{
			SessionId: sessionId,
		}
		response.Json(w, resp, http.StatusOK)
	}
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

		sessionId, err := handler.AuthService.Login(body.SessionId, body.Code)
		if err != nil {
			response.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}
		fmt.Println(sessionId)

		j := jwt.NewJWT(handler.Config.Auth.Secret)
		token, err := j.Create(jwt.JWTData{SessionId: sessionId, Code: body.Code})
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

		email, err := handler.AuthService.Register(body.Phone)
		if err != nil {
			response.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}
		fmt.Println(email)

		resp := RegisterResponse{
			Message: "Register success",
		}
		response.Json(w, resp, http.StatusOK)
	}
}
