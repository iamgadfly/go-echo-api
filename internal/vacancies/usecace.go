package vacancies

import (
	"context"
	"github.com/iamgadfly/go-echo-api/internal/models"
)

type UseCase interface {
	Create(ctx context.Context, link string) (*models.Vacancy, error)
	GetById(ctx context.Context, id string) (*models.Vacancy, error)
}
