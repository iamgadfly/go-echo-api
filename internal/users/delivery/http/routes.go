package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MapUserRoutes(authGroup *echo.Group, h userHandlers) {
	// mw *middleware.MiddlewareManager
	authGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[{$time_rfc3339}] ${status} ${method} ${path} ${host} ${remote_ip} ${latency_human}` + "\n",
	}))

	authGroup.POST("/register", h.Register())
	authGroup.POST("/login", h.Login())
	authGroup.GET("/all", h.GetUsers())
	//authGroup.POST("/login", h.Login())
	//authGroup.POST("/logout", h.Logout())
	//authGroup.GET("/find", h.FindByName())
	//authGroup.GET("/all", h.GetUsers())
	//authGroup.GET("/:user_id", h.GetUserByID())
}
