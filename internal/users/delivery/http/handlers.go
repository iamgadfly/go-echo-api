package http

import (
	"github.com/AleksK1NG/api-mc/pkg/httpErrors"
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/iamgadfly/go-echo-api/internal/users"
	"github.com/iamgadfly/go-echo-api/pkg/req"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type userHandlers struct {
	cfg    *config.Config
	userUC users.UseCase
	logger *zap.SugaredLogger
}

func NewAuthHandlers(cfg *config.Config, userUC users.UseCase, logger *zap.SugaredLogger) userHandlers {
	return userHandlers{
		cfg:    cfg,
		userUC: userUC,
		logger: logger,
	}
}

func (h *userHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := models.User{}
		if err := c.Bind(&user); err != nil {
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		createdUser, err := h.userUC.Register(&user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, createdUser)
	}
}

func (h *userHandlers) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := h.userUC.GetUsers()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (h *userHandlers) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := req.ParseReq(c)
		if err != nil {
			return err
		}
		res, er := h.userUC.Login(data["password"].(string), data["email"].(string))
		if er != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": er.Error(),
			})
		}

		return c.JSON(http.StatusOK, res)
	}
}
