package usecase

import (
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/iamgadfly/go-echo-api/internal/vacancies"
	"go.uber.org/zap"
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

func (v VacancyUC) Create(link string) (models.Vacancy, error) {
	return models.Vacancy{}, nil
}
