package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MapUserRoutes(e *echo.Group, h userHandlers) {
	// mw *middleware.MiddlewareManager
	//authGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//	Format: `[{$time_rfc3339}] ${status} ${method} ${path} ${host} ${remote_ip} ${latency_human}` + "\n",
	//}))

	e.POST("/register", h.Register())
	e.POST("/login", h.Login())

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(h.cfg.Server.JwtSecretKey),
	}))

	e.GET("/all", h.GetUsers())
	//authGroup.GET("/:user_id", h.GetUserByID())
}
