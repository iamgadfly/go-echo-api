package grpcserver

import (
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

// ServerGRPC ..
type ServerGRPC struct {
	db     *sqlx.DB
	cfg    *config.Config
	logger *zap.SugaredLogger
}

// NewServerGRPC ..
func NewServerGRPC(db *sqlx.DB, cfg *config.Config, logger *zap.SugaredLogger) *ServerGRPC {
	return &ServerGRPC{
		db:     db,
		cfg:    cfg,
		logger: logger,
	}
}

// Run ..
func (s *ServerGRPC) Run(er chan error) {
	log.Println("starting grpc server!")

	server := grpc.NewServer()
	s.MapServices(server)

	l, err := net.Listen("tcp", s.cfg.Server.PprofPort)
	if err != nil {
		er <- err
	}
	defer l.Close()
	defer server.Stop()

	err = server.Serve(l)
	if err != nil {
		er <- err
	}

	er <- nil
}
