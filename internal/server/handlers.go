package server

import (
	productHttp "github.com/iamgadfly/go-echo-api/internal/products/delivery/http"
	productRepository "github.com/iamgadfly/go-echo-api/internal/products/repository"
	productUC "github.com/iamgadfly/go-echo-api/internal/products/usecase"
	vacancyHttp "github.com/iamgadfly/go-echo-api/internal/vacancies/delivery/http"
	vacancyRepository "github.com/iamgadfly/go-echo-api/internal/vacancies/repository"
	vacancyUC "github.com/iamgadfly/go-echo-api/internal/vacancies/usecase"

	userHttp "github.com/iamgadfly/go-echo-api/internal/users/delivery/http"
	userRepository "github.com/iamgadfly/go-echo-api/internal/users/repository"
	userUseCase "github.com/iamgadfly/go-echo-api/internal/users/usecase"
	"github.com/labstack/echo/v4"
	//echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	uRepo := userRepository.NewUsersRepository(s.db)
	userUC := userUseCase.NewUserUseCase(s.cfg, uRepo, s.logger)
	userHandlers := userHttp.NewAuthHandlers(s.cfg, userUC, s.logger)

	v1 := e.Group("/api/v1")
	usersGroup := v1.Group("/users")
	userHttp.MapUserRoutes(usersGroup, userHandlers)

	pRepo := productRepository.NewUProductsRepository(s.db, s.logger)
	productUC := productUC.NewProductUseCase(s.cfg, pRepo, s.logger)
	productHandlers := productHttp.NewProductHandlers(productUC, s.cfg, s.logger)
	productsGroup := v1.Group("/products")
	productHttp.MapProductRoutes(productsGroup, productHandlers)

	vRepo := vacancyRepository.NewVacanciesRepository(s.db)
	vacancyUC := vacancyUC.NewVacancyUseCase(s.cfg, vRepo, s.logger)
	vacancyHandlers := vacancyHttp.NewVacancyHandlers(s.cfg, s.logger, vacancyUC)
	vacanciesGroup := v1.Group("/vacancy")
	vacancyHttp.MapVacancyRoutes(vacanciesGroup, vacancyHandlers)

	//newsHttp.MapNewsRoutes(newsGroup, newsHandlers, mw)

	return nil
}
