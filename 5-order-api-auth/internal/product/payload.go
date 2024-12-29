package product

type ProductCreateRequest struct {
	Name        string   `json:"name" validate:"required"`
	Desctiption string   `json:"description"`
	Images      []string `json:"images"`
}

type ProductUpdateRequest struct {
	Name        string   `json:"name" validate:"required"`
	Desctiption string   `json:"description"`
	Images      []string `json:"images"`
}
