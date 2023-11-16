package server

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"time"
)

const (
	certFile       = "ssl/Server.crt"
	keyFile        = "ssl/Server.pem"
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

// Server struct
type Server struct {
	port string
	echo *echo.Echo
	db   *sqlx.DB
}

// NewServer New Server constructor
func NewServer(port string, db *sqlx.DB) *Server {
	return &Server{echo: echo.New(), db: db, port: port}
}

func (s *Server) Run() error {
	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}
	if err := s.echo.Start(s.port); err != nil {
		return err
	}

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()
	return s.echo.Server.Shutdown(ctx)
}
