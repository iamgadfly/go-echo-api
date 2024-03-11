package tests

import (
	"context"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/iamgadfly/go-echo-api/internal/vacancies/usecase"
	"github.com/iamgadfly/go-echo-api/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UnitTestVacancy struct {
	suite.Suite
}

func TestUnitTestVacancy(t *testing.T) {
	suite.Run(t, &UnitTestVacancy{})
}

func (uts *UnitTestVacancy) TestParse() {
	repo := mocks.VacancyRepository{}
	repo.On("Create", mock.AnythingOfType("*models.Vacancy")).Return(&models.Vacancy{}, nil).Once()
	useCase := usecase.NewVacancyUseCase(nil, &repo, nil)
	_, err := useCase.Create(context.Background(), "https://cheboksary.hh.ru/vacancy/93356282?query=Go+%D1%80%D0%B0%D0%B7%D1%80%D0%B0%D0%B1%D0%BE%D1%82%D1%87%D0%B8%D0%BA&hhtmFrom=vacancy_search_list")
	uts.Nil(err)
}

func (uts *UnitTestVacancy) TestGetById() {
	repo := mocks.VacancyRepository{}
	repo.On("FindById", mock.AnythingOfType("int")).Return(&models.Vacancy{}, nil).Once()
	useCase := usecase.NewVacancyUseCase(nil, &repo, nil)
	_, err := useCase.GetById(context.Background(), "52")
	uts.Nil(err)
}
