package vacancies

import "github.com/iamgadfly/go-echo-api/internal/models"

type UseCase interface {
	Create(link string) (models.Vacancy, error)
}
