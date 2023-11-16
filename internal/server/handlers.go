package server

import (
	userHttp "github.com/iamgadfly/go-echo-api/internal/users/delivery/http"
	userRepository "github.com/iamgadfly/go-echo-api/internal/users/repository"
	userUseCase "github.com/iamgadfly/go-echo-api/internal/users/usecase"
	"github.com/labstack/echo/v4"
	//echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	uRepo := userRepository.NewUsersRepository(s.db)
	userUC := userUseCase.NewUserUseCase(uRepo)
	userHandlers := userHttp.NewAuthHandlers(userUC)

	v1 := e.Group("/api/v1")
	usersGroup := v1.Group("/users")
	userHttp.MapUserRoutes(usersGroup, userHandlers)
	//newsHttp.MapNewsRoutes(newsGroup, newsHandlers, mw)

	return nil
}
