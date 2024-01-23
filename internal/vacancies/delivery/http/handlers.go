package http

import (
	"context"
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/vacancies"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type vacancyHandlers struct {
	cfg       *config.Config
	logger    *zap.SugaredLogger
	vacancyUC vacancies.UseCase
}

func NewVacancyHandlers(cfg *config.Config, logger *zap.SugaredLogger, vacancyUC vacancies.UseCase) vacancyHandlers {
	return vacancyHandlers{
		cfg:       cfg,
		logger:    logger,
		vacancyUC: vacancyUC,
	}
}

func (v vacancyHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))
		defer cancel()

		link := c.QueryParam("link")
		if link == "" {
			return c.JSON(http.StatusBadRequest, "link is empty")
		}

		res, err := v.vacancyUC.Create(ctx, link)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, res)
	}
}
