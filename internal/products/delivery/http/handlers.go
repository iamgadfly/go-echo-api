package http

import (
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/products"
	"github.com/iamgadfly/go-echo-api/internal/products/usecase"
	"github.com/iamgadfly/go-echo-api/pkg/req"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
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
