package products

import (
	"context"
	"github.com/iamgadfly/go-echo-api/internal/models"
)

type ProductRepository interface {
	GetProducts() ([]models.Product, error)
	CreateBatch(products []models.Product) error
	SearchByShopId(prod models.Product) error
	Create(product models.Product) error
	Search(ctx context.Context, word string) ([]models.Product, error)
	GetById(ctx context.Context, id int) (*models.Product, error)
}
