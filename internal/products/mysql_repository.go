package products

import "github.com/iamgadfly/go-echo-api/internal/models"

type Repository interface {
	Create(product *models.Product) (models.Product, error)
	GetProducts() (models.ProductList, error)
	CreateBatch(products []models.Product) error
	//Find(id interface{}) (models.Product, error)
	//Update(product *models.Product) (models.Product, error)
	//GetProducts() ([]models.Product, error)
}
