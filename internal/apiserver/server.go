package apiserver

import (
	"fmt"
	"log/slog"

	genericoptions "github.com/RadishXZ/cloudtest/pkg/options"
)

type Config struct {
	MySQLOptions *genericoptions.MySQLOptions
}

type Server struct {
	cfg *Config
}

func (cfg *Config) NewServer() (*Server, error) {
	return &Server{cfg: cfg}, nil
}

func (s *Server) Run() error {
	slog.Info("Read MySQL host from config", "mysql.addr", s.cfg.MySQLOptions.Addr)
	fmt.Printf("Read MySQL host from config: %s\n", s.cfg.MySQLOptions.Addr)

	return nil
}
