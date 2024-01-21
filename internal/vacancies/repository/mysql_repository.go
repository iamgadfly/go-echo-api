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

func (r VacanciesRepo) Create(vacancy models.Vacancy) (models.Vacancy, error) {
	return models.Vacancy{}, nil
}
