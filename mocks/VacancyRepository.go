// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	models "github.com/iamgadfly/go-echo-api/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// VacancyRepository is an autogenerated mock type for the VacancyRepository type
type VacancyRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: vacancy
func (_m *VacancyRepository) Create(vacancy *models.Vacancy) (*models.Vacancy, error) {
	ret := _m.Called(vacancy)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *models.Vacancy
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Vacancy) (*models.Vacancy, error)); ok {
		return rf(vacancy)
	}
	if rf, ok := ret.Get(0).(func(*models.Vacancy) *models.Vacancy); ok {
		r0 = rf(vacancy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Vacancy)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.Vacancy) error); ok {
		r1 = rf(vacancy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *VacancyRepository) FindById(id int) (*models.Vacancy, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *models.Vacancy
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.Vacancy, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *models.Vacancy); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Vacancy)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByVacancyId provides a mock function with given fields: id
func (_m *VacancyRepository) FindByVacancyId(id int) (models.Vacancy, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindByVacancyId")
	}

	var r0 models.Vacancy
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (models.Vacancy, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) models.Vacancy); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Vacancy)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewVacancyRepository creates a new instance of VacancyRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVacancyRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *VacancyRepository {
	mock := &VacancyRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
