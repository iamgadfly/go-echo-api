package vacancies

import "github.com/iamgadfly/go-echo-api/internal/models"

type Repository interface {
	Create(vacancy *models.Vacancy) (*models.Vacancy, error)
	FindByVacancyId(id int) (models.Vacancy, error)
	FindById(id int) (*models.Vacancy, error)
}
