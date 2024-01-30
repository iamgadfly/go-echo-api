package usecase

import (
	"context"
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/iamgadfly/go-echo-api/internal/vacancies"
	"github.com/iamgadfly/go-echo-api/pkg/parse/hh"
	"go.uber.org/zap"
	"strconv"
)

type VacancyUC struct {
	cfg         *config.Config
	vacancyRepo vacancies.Repository
	logger      *zap.SugaredLogger
}

func NewVacancyUseCase(cfg *config.Config, repository vacancies.Repository, logger *zap.SugaredLogger) VacancyUC {
	return VacancyUC{
		cfg:         cfg,
		vacancyRepo: repository,
		logger:      logger,
	}
}

func (v VacancyUC) Create(ctx context.Context, link string) (*models.Vacancy, error) {
	vacancy, err := hh.Parse(ctx, link)
	if err != nil {
		return &models.Vacancy{}, err
	}
	res, err := v.vacancyRepo.Create(vacancy)
	if err != nil {
		return &models.Vacancy{}, err
	}

	return res, nil
}

func (v VacancyUC) GetById(ctx context.Context, rawId string) (*models.Vacancy, error) {
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return nil, err
	}

	vacancy, er := v.vacancyRepo.FindById(id)
	if er != nil {
		return nil, er
	}

	return vacancy, nil
}
