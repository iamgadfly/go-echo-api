package usecase

import (
	"context"
	"encoding/csv"
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/iamgadfly/go-echo-api/internal/products/repository"
	"github.com/iamgadfly/go-echo-api/pkg/parse/avito"
	"github.com/iamgadfly/go-echo-api/pkg/parse/ozon"
	"github.com/iamgadfly/go-echo-api/pkg/parse/wb"
	"go.uber.org/zap"
	"os"
	"strconv"
	"strings"
	"time"
)

type ProductUseCase struct {
	cfg         *config.Config
	productRepo *repository.ProductRepo
	logger      *zap.SugaredLogger
}

func NewProductUseCase(cfg *config.Config, productRepo *repository.ProductRepo, logger *zap.SugaredLogger) ProductUseCase {
	return ProductUseCase{
		cfg:         cfg,
		productRepo: productRepo,
		logger:      logger,
	}
}

func (p ProductUseCase) ParseByLink(link string) (models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*15))
	defer cancel()
	var err error
	var prod models.Product
	domain := strings.Split(link, "/")
	switch domain[2] {
	case "www.avito.ru":
		prod, err = avito.ParseByLink(link)
	case "www.wildberries.ru":
		prod, err = wb.ParseByLink(ctx, link)
	case "www.ozon.ru":
		prod, err = ozon.Parse(link)
	}
	if err != nil {
		return prod, err
	}
	err = p.productRepo.SearchByShopId(prod)
	p.logger.Info(err)

	if err != nil {
		return prod, err
	}

	return prod, nil
}

func (p ProductUseCase) WriteCsv() error {
	products, err := p.productRepo.GetProducts()
	if err != nil {
		return err
	}

	time := strconv.Itoa(int(time.Now().Unix()))
	file, err := os.Create("storage/products-" + time + ".csv")
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	columns := []string{"ID", "name", "price", "sale_price", "color", "link"}
	writer.Write(columns)
	for _, product := range products {
		writer.Write([]string{strconv.Itoa(product.ID), product.Name, strconv.Itoa(int(product.Price)), strconv.Itoa(int(product.SalePrice)), product.Color, product.Link})
	}

	return nil
}
