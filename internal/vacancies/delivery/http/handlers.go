package http

import (
	"context"
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/vacancies"
	"github.com/iamgadfly/go-echo-api/pkg/req"
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

// Create
// @Summary Create vacancies by link
// @tags vacancy
// @description Create vacancy by link
// @Param word body models.VacancyLink true "link to parse"
// @accept json
// @produce json
// @success 200 {object} models.Vacancy
// @router /api/v1/vacancy/create [post]
func (v vacancyHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))
		defer cancel()
		data, er := req.ParseReq(c)
		if er != nil {
			return c.JSON(http.StatusBadRequest, er.Error())
		}
		if data["link"].(string) == "" {
			return c.JSON(http.StatusBadRequest, "link is empty")
		}
		res, err := v.vacancyUC.Create(ctx, data["link"].(string))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, res)
	}
}

// GetById
// @Summary Get vacancies by id
// @tags vacancy
// @description Get vacancy by id
// @Param id path int true "vacancy id"
// @accept json
// @produce json
// @success 200 {object} models.Vacancy
// @router /api/v1/vacancy/{id} [get]
func (v vacancyHandlers) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))
		defer cancel()

		res, err := v.vacancyUC.GetById(ctx, c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, res)
	}
}
