package http

import "github.com/labstack/echo/v4"

func MapUserRoutes(authGroup *echo.Group, h userHandlers) {
	// mw *middleware.MiddlewareManager

	authGroup.POST("/register", h.Register())
	//authGroup.POST("/login", h.Login())
	//authGroup.POST("/logout", h.Logout())
	//authGroup.GET("/find", h.FindByName())
	//authGroup.GET("/all", h.GetUsers())
	//authGroup.GET("/:user_id", h.GetUserByID())
}
