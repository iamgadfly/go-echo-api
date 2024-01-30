package http

import "github.com/labstack/echo/v4"

func MapVacancyRoutes(e *echo.Group, h vacancyHandlers) {
	e.GET("/create", h.Create())
	e.GET("/:id", h.GetById())

}
