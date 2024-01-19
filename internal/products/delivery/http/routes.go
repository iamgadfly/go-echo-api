package http

import "github.com/labstack/echo/v4"

func MapProductRoutes(e *echo.Group, h ProductHandlers) {
	e.POST("/parse", h.ParseByLink())
	e.POST("/get_csv", h.GetCsv())
	e.POST("/parse_cat_wb", h.ParseWbCat())
	e.POST("/search", h.Search())
}
