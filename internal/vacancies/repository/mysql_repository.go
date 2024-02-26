package repository

import (
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type VacanciesRepo struct {
	db *sqlx.DB
}

// NewVacanciesRepository ..
func NewVacanciesRepository(db *sqlx.DB) *VacanciesRepo {
	return &VacanciesRepo{
		db: db,
	}
}

// Create ..
func (r VacanciesRepo) Create(vacancy *models.Vacancy) (*models.Vacancy, error) {
	check, err := r.FindByVacancyId(vacancy.VacancyId)
	if err != nil {
		return &models.Vacancy{}, err
	}

	if check.ID != 0 {
		vacancy.ID = check.ID
		vacancy, err = r.Update(vacancy)
		if err != nil {
			return nil, err
		}

		return vacancy, nil
	}
	_, err = r.db.NamedExec(Create, vacancy)
	if err != nil {
		return &models.Vacancy{}, err
	}

	return vacancy, nil
}

// Update ..
func (r VacanciesRepo) Update(vacancy *models.Vacancy) (*models.Vacancy, error) {
	vacancy.UpdatedAt = time.Now().Format(time.DateTime)
	_, err := r.db.NamedExec(Update, vacancy)
	if err != nil {
		return nil, err
	}
	vacancy, err = r.FindById(vacancy.ID)
	if err != nil {
		return nil, err
	}
	return vacancy, nil
}

// FindByVacancyId ..
func (r VacanciesRepo) FindByVacancyId(id int) (models.Vacancy, error) {
	vacancy := models.Vacancy{}
	err := r.db.Get(&vacancy, FindByVacancyId, id)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return models.Vacancy{}, err
	}
	return vacancy, nil
}

// FindById ..
func (r VacanciesRepo) FindById(id int) (*models.Vacancy, error) {
	vacancy := models.Vacancy{}
	err := r.db.Get(&vacancy, FindById, id)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	return &vacancy, nil
}
