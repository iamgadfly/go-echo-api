package repository

import (
	"fmt"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type ProductRepo struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewUProductsRepository(db *sqlx.DB, logger *zap.SugaredLogger) *ProductRepo {
	return &ProductRepo{
		db:     db,
		logger: logger,
	}
}

func (r ProductRepo) Create(product models.Product) error {
	return nil
}

func (r ProductRepo) SearchByShopId(prod models.Product) error {
	check := models.Product{}
	err := r.db.Get(&check, SearchByShopId, prod.ShopId)
	if err.Error() == "sql: no rows in result set" {
		_, err = r.db.NamedExec(CreateProduct, prod)
	} else {
		fmt.Println("этот товар уже есть!")
		_, err = r.db.NamedExec(Update, prod)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r ProductRepo) GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Select(&products, GetProducts)
	if err != nil {
		return []models.Product{}, err
	}

	return products, nil
}
