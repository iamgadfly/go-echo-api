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
