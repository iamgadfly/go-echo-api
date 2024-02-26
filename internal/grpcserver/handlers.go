package grpcserver

import (
	pRepo "github.com/iamgadfly/go-echo-api/internal/products/repository"
	productUC "github.com/iamgadfly/go-echo-api/internal/products/usecase"
	"github.com/iamgadfly/go-echo-api/internal/users/repository"
	vRepo "github.com/iamgadfly/go-echo-api/internal/vacancies/repository"
	"github.com/iamgadfly/go-echo-api/pkg/api/product_v1"
	product "github.com/iamgadfly/go-echo-api/pkg/api/product_v1/proto"
	"github.com/iamgadfly/go-echo-api/pkg/api/user_v1"
	"github.com/iamgadfly/go-echo-api/pkg/api/user_v1/proto"
	"github.com/iamgadfly/go-echo-api/pkg/api/vacancy_v1"
	vacancy "github.com/iamgadfly/go-echo-api/pkg/api/vacancy_v1/proto"
	"google.golang.org/grpc"
)

// MapServices ..
func (s *ServerGRPC) MapServices(server *grpc.Server) {
	// user service
	userRepo := repository.NewUsersRepository(s.db)
	sUser := user_v1.NewUserServer(s.cfg, userRepo)

	// vacancy service
	vacancyRepo := vRepo.NewVacanciesRepository(s.db)
	sVacancy := vacancy_v1.NewVacancyServer(s.cfg, vacancyRepo)

	// product service
	prodRepo := pRepo.NewUProductsRepository(s.db, s.logger)
	pUseCase := productUC.NewProductUseCase(s.cfg, prodRepo, s.logger)
	sProduct := product_v1.NewProductServer(s.cfg, prodRepo, pUseCase)

	//register services
	proto.RegisterUserServer(server, sUser)
	vacancy.RegisterVacancyServer(server, sVacancy)
	product.RegisterProductServer(server, sProduct)
}
