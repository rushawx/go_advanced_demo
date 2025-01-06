package jwt_test

import (
	"0-hello/pkg/jwt"
	"testing"
)

func TestJwtCreate(t *testing.T) {
	const email = "a@a.ru"
	jwtService := jwt.NewJWT("42")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("token is invalid")
	}
	if data.Email != email {
		t.Fatalf("email %s not equal %s", data.Email, email)
	}
}
