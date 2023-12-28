package usecase

import (
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/iamgadfly/go-echo-api/internal/products/repository"
	"github.com/iamgadfly/go-echo-api/pkg/parse/avito"
	"github.com/iamgadfly/go-echo-api/pkg/parse/ozon"
	"github.com/iamgadfly/go-echo-api/pkg/parse/wb"
	"go.uber.org/zap"
	"strings"
)

type ProductUseCase struct {
	cfg         *config.Config
	productRepo *repository.ProductRepo
	logger      *zap.SugaredLogger
}

func (p ProductUseCase) ParseByLink(link string) (models.Product, error) {
	var err error
	var prod models.Product
	domain := strings.Split(link, "/")
	switch domain[2] {
	case "www.avito.ru":
		prod, err = avito.ParseByLink(link)
	case "www.wildberries.ru":
		prod, err = wb.ParseByLink(link)
	case "www.ozon.ru":
		prod, err = ozon.ParseByLink(link)
	}

	if err != nil {
		return prod, err
	}

	err = p.productRepo.SearchByShopId(prod)
	if err != nil {
		return prod, err
	}

	return prod, nil
}

func NewProductUseCase(cfg *config.Config, productRepo *repository.ProductRepo, logger *zap.SugaredLogger) ProductUseCase {
	return ProductUseCase{
		cfg:         cfg,
		productRepo: productRepo,
		logger:      logger,
	}
}
