package link

import (
	"0-hello/pkg/request"
	"0-hello/pkg/response"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}

	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.GoTo())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Create")
		body, err := request.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		for {
			linkFromDb, _ := handler.LinkRepository.GetByHash(link.Hash)
			if linkFromDb == nil {
				break
			}
			link.GenerateHash()
		}
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, createdLink, http.StatusCreated)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Update")
		body, err := request.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, link, http.StatusCreated)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Delete")
		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = handler.LinkRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, nil, http.StatusOK)
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GoTo")
		hash := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		fmt.Println(link.Url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, link.Url, http.StatusMovedPermanently)
	}
}
