package vacancy_v1

import (
	"context"
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/iamgadfly/go-echo-api/internal/vacancies/repository"
	vacancy "github.com/iamgadfly/go-echo-api/pkg/api/vacancy_v1/proto"
	"github.com/iamgadfly/go-echo-api/pkg/parse/hh"
)

// VacancyServer ..
type VacancyServer struct {
	vacancy.UnimplementedVacancyServer
	cfg  *config.Config
	repo *repository.VacanciesRepo
}

// NewVacancyServer ..
func NewVacancyServer(cfg *config.Config, repo *repository.VacanciesRepo) *VacancyServer {
	return &VacancyServer{
		cfg:  cfg,
		repo: repo,
	}
}

// Create ..
func (v VacancyServer) Create(ctx context.Context, req *vacancy.CreateRequest) (*vacancy.VacancyResponse, error) {
	parsedVacancy, err := hh.Parse(ctx, req.GetLink())
	if err != nil {
		return nil, err
	}
	parsedVacancy, err = v.repo.Create(parsedVacancy)
	if err != nil {
		return nil, err
	}

	return &vacancy.VacancyResponse{Data: createResponse(parsedVacancy)}, nil
}

// GetById ..
func (v VacancyServer) GetById(ctx context.Context, req *vacancy.GetByIdRequest) (*vacancy.VacancyResponse, error) {
	vac, err := v.repo.FindById(int(req.GetId()))
	if err != nil {
		return nil, err
	}

	return &vacancy.VacancyResponse{Data: createResponse(vac)}, nil
}

func createResponse(v *models.Vacancy) *vacancy.VacancyData {
	return &vacancy.VacancyData{
		Id:          uint64(v.ID),
		Name:        v.Name,
		VacancyId:   uint64(v.VacancyId),
		Experience:  v.Experience,
		Remote:      v.Remote,
		Salary:      v.Salary,
		Status:      v.Status,
		Link:        v.Link,
		Description: v.Desc,
		Skills:      v.Skills,
		CompanyName: v.CompanyName,
		CompanyUrl:  v.CompanyUrl,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}
}
