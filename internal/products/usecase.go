package products

import (
	"context"
	"github.com/iamgadfly/go-echo-api/internal/models"
)

type ProductUseCase interface {
	ParseByLink(link string) (models.Product, error)
	WriteCsv() error
	ParseWbCat(urls []string) error
	Search(ctx context.Context, string string) ([]models.Product, error)
	GetById(ctx context.Context, id string) (*models.Product, error)
	//Find(id interface{}) (models.Product, error)
	//Update(product *models.Product) (models.Product, error)
	//GetProducts() ([]models.Product, error)
}
