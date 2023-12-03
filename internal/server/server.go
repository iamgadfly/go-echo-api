package server

import (
	"context"
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
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
	cfg    *config.Config
	port   string
	echo   *echo.Echo
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

// NewServer New Server constructor
func NewServer(cfg *config.Config, port string, db *sqlx.DB, logger *zap.SugaredLogger) *Server {
	return &Server{cfg: cfg, echo: echo.New(), db: db, port: port, logger: logger}
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
