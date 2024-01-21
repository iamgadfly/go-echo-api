package vacancies

import "github.com/iamgadfly/go-echo-api/internal/models"

type Repository interface {
	Create(vacancy models.Vacancy) (models.Vacancy, error)
}
