package repository

import (
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type VacanciesRepo struct {
	db *sqlx.DB
}

func NewVacanciesRepository(db *sqlx.DB) *VacanciesRepo {
	return &VacanciesRepo{
		db: db,
	}
}

func (r VacanciesRepo) Create(vacancy *models.Vacancy) (*models.Vacancy, error) {
	check, err := r.FindByVacancyId(vacancy.VacancyId)
	if err != nil {
		return &models.Vacancy{}, err
	}
	if check.ID != 0 {
		return &check, nil
	}
	_, err = r.db.NamedExec(Create, vacancy)
	if err != nil {
		return &models.Vacancy{}, err
	}

	return vacancy, nil
}

func (r VacanciesRepo) FindByVacancyId(id int) (models.Vacancy, error) {
	vacancy := models.Vacancy{}
	err := r.db.Get(&vacancy, FindByVacancyId, id)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return models.Vacancy{}, err
	}
	return vacancy, nil
}
