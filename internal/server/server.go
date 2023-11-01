package server

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
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
	err := s.echo.Start(s.port)
	if err != nil {
		return err
	}
	return nil

	//if err := s.MapHandlers(s.echo); err != nil {
	//	return err
	//}
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	//<-quit
	//
	//ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	//defer shutdown()
	//
	//return s.echo.Server.Shutdown(ctx)
}
