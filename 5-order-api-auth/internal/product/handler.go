package product

import (
	"5-order-api-auth/configs"
	"5-order-api-auth/pkg/middleware"
	"5-order-api-auth/pkg/request"
	"5-order-api-auth/pkg/response"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandlerDeps struct {
	ProductRepository *ProductRepository
	Config            *configs.Config
}

type ProductHandler struct {
	ProductRepository *ProductRepository
	Config            *configs.Config
}

func NewProductHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{
		ProductRepository: deps.ProductRepository,
	}

	router.Handle("POST /product", middleware.IsAuthed(handler.Create(), deps.Config))
	router.Handle("PATCH /product/{id}", middleware.IsAuthed(handler.Update(), deps.Config))
	router.Handle("DELETE /product/{id}", middleware.IsAuthed(handler.Delete(), deps.Config))
	router.HandleFunc("GET /product/{id}", handler.GetById())
}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Create")
		sessionId, ok := r.Context().Value(middleware.ContextSessionIdKey).(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		fmt.Println("SessionId:", sessionId)
		code, ok := r.Context().Value(middleware.CodeKey).(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		fmt.Println("Code:", code)
		body, err := request.HandleBody[ProductCreateRequest](&w, r)
		if err != nil {
			return
		}
		product := NewProduct(body.Name, body.Desctiption, body.Images)
		createdProduct, err := handler.ProductRepository.Create(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, createdProduct, http.StatusCreated)
	}
}

func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Update")
		sessionId, ok := r.Context().Value(middleware.ContextSessionIdKey).(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		fmt.Println("SessionId:", sessionId)
		code, ok := r.Context().Value(middleware.CodeKey).(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		fmt.Println("Code:", code)
		body, err := request.HandleBody[ProductUpdateRequest](&w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product, err := handler.ProductRepository.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        body.Name,
			Desctiption: body.Desctiption,
			Images:      body.Images,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		response.Json(w, product, http.StatusOK)
	}
}

func (handler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Delete")
		sessionId, ok := r.Context().Value(middleware.ContextSessionIdKey).(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		fmt.Println("SessionId:", sessionId)
		code, ok := r.Context().Value(middleware.CodeKey).(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		fmt.Println("Code:", code)
		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = handler.ProductRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, nil, http.StatusOK)
	}
}

func (handler *ProductHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetById")
		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		product, err := handler.ProductRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, product, http.StatusOK)
	}
}
