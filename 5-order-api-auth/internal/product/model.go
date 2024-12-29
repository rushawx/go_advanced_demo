package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Desctiption string
	Images      pq.StringArray `gorm:"type:text[]"`
}

func NewProduct(name, description string, images []string) *Product {
	return &Product{
		Name:        name,
		Desctiption: description,
		Images:      images,
	}
}
