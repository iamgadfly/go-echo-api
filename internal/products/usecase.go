package products

import (
	"github.com/iamgadfly/go-echo-api/internal/models"
)

type ProductUseCase interface {
	ParseByLink(link string) (models.Product, error)
	WriteCsv() error
	ParseWbCat(urls []string) error
	//Find(id interface{}) (models.Product, error)
	//Update(product *models.Product) (models.Product, error)
	//GetProducts() ([]models.Product, error)
}
