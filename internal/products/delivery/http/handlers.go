package http

import (
	"context"
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/products"
	"github.com/iamgadfly/go-echo-api/internal/products/usecase"
	"github.com/iamgadfly/go-echo-api/pkg/req"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type ProductHandlers struct {
	productUC products.ProductUseCase
	cfg       *config.Config
	logger    *zap.SugaredLogger
}

func NewProductHandlers(productUC usecase.ProductUseCase, cfg *config.Config, logger *zap.SugaredLogger) ProductHandlers {
	return ProductHandlers{
		productUC: productUC,
		cfg:       cfg,
		logger:    logger,
	}
}

// ParseByLink
// @Summary Parse By Link
// @tags products
// @description Parse by link
// @Param link body models.ProductLink true "link on product"
// @success 200 {object} models.Product
// @router /api/v1/products/parse [post]
func (h *ProductHandlers) ParseByLink() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, _ := req.ParseReq(c)
		prod, err := h.productUC.ParseByLink(data["link"].(string))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, prod)
	}
}

// GetCsv
// @Summary Get Csv
// @tags products
// @description Parse by link
// @accept json
// @produce json
// @success 200 {string} string "created at"
// @router /api/v1/products/get_csv [post]
func (h *ProductHandlers) GetCsv() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := h.productUC.WriteCsv()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, "created at "+time.Now().String())
	}
}

func (h *ProductHandlers) ParseWbCat() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, _ := req.ParseReq(c)
		var urls []string
		for n := 1; n <= 52; n++ {
			urls = append(urls, data["link"].(string)+"&page="+strconv.Itoa(n))
		}

		err := h.productUC.ParseWbCat(urls)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, "start..")
	}
}

// Search
// @Summary Get by product i
// @tags products
// @description Get product by word
// @Param word body models.ProductWord true "word to search"
// @accept json
// @produce json
// @success 200 {object} []models.Product
// @router /api/v1/products/search [post]
func (h *ProductHandlers) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))
		defer cancel()
		data, _ := req.ParseReq(c)
		res, err := h.productUC.Search(ctx, data["word"].(string))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, res)
	}
}

// GetById
// @Summary Get by product i
// @tags products
// @description Get by product id
// @Param id path int true "product id"
// @accept json
// @produce json
// @success 200 {object} models.Product
// @router /api/v1/products/{id} [get]
func (h *ProductHandlers) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))
		defer cancel()
		res, err := h.productUC.GetById(ctx, c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, res)
	}
}
