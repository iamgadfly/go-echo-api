package http

import (
	"fmt"
	"github.com/AleksK1NG/api-mc/pkg/httpErrors"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/iamgadfly/go-echo-api/internal/users"
	"github.com/labstack/echo/v4"
	"net/http"
)

type userHandlers struct {
	userUC users.UseCase
}

func NewAuthHandlers(userUC users.UseCase) userHandlers {
	return userHandlers{userUC: userUC}
}

func (h *userHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := echo.Map{}
		c.Bind(&data)
		user := models.User{}
		userToken, err := h.userUC.Register(data, &user)
		if err != nil {
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, userToken)
	}
}

func (h *userHandlers) Test() echo.HandlerFunc {
	return func(c echo.Context) error {
		//data, _ := c.FormParams()
		data := echo.Map{}
		c.Bind(&data)
		fmt.Println(data)
		fmt.Println(data["test"])
		fmt.Println(data["name"])

		return c.JSON(http.StatusOK, map[string]string{
			"yes": "true",
		})
	}
}
