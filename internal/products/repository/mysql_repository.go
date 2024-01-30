package repository

import (
	"fmt"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golang.org/x/net/context"
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

func (r ProductRepo) CreateBatch(products []models.Product) error {
	if len(products) == 0 {
		return nil
	}

	_, err := r.db.NamedExec(CreateBatch, products)
	if err != nil {
		return err
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

func (r ProductRepo) Search(ctx context.Context, word string) ([]models.Product, error) {
	var products []models.Product
	if err := r.db.SelectContext(ctx, &products, Search, word+"%"); err != nil {
		return []models.Product{}, err
	}

	return products, nil
}
func (r ProductRepo) GetById(ctx context.Context, id int) (*models.Product, error) {
	prod := models.Product{}
	err := r.db.Get(&prod, GetById, id)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	return &prod, nil
}
