package middleware

import (
	"5-order-api-auth/configs"
	"5-order-api-auth/pkg/jwt"
	"context"
	"fmt"
	"net/http"
	"strings"
)

type key string

const (
	ContextSessionIdKey key = "ContextSessionIdKey"
	CodeKey             key = "CodeKey"
)

func writeUnauthorized(w http.ResponseWriter) {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			writeUnauthorized(w)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		fmt.Println("Token:", token)
		fmt.Println("Valid:", isValid)
		fmt.Println("Data:", data)
		if !isValid {
			writeUnauthorized(w)
			return
		}
		ctx := context.WithValue(r.Context(), ContextSessionIdKey, data.SessionId)
		ctx = context.WithValue(ctx, CodeKey, data.Code)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
