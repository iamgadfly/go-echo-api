package repository

import (
	"fmt"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewUProductsRepository(db *sqlx.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
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
