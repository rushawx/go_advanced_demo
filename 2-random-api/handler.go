package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
)

type RandomApiHandler struct{}

func NewRandomApiHandler(router *http.ServeMux) {
	handler := &RandomApiHandler{}
	router.HandleFunc("/random", handler.Random())
}

func (handler *RandomApiHandler) Random() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Random")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(rand.IntN(6) + 1)))
	}
}
