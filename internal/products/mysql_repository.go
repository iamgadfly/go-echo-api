package products

import (
	"context"
	"github.com/iamgadfly/go-echo-api/internal/models"
)

type Repository interface {
	GetProducts() ([]models.Product, error)
	CreateBatch(products []models.Product) error
	SearchByShopId(prod models.Product) error
	Create(product models.Product) error
	Search(ctx context.Context, word string) ([]models.Product, error)
}
