package config

import (
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

type (
	Config struct {
		GRPC
		PG
	}

	GRPC struct {
		Port        string `env:"GRPC_PORT"`
		GatewayPort string `env:"GRPC_GATEWAY_PORT"`
	}

	PG struct {
		URL      string
		Host     string `env:"POSTGRES_HOST"`
		Port     string `env:"POSTGRES_PORT"`
		DB       string `env:"POSTGRES_DB"`
		User     string `env:"POSTGRES_USER"`
		Password string `env:"POSTGRES_PASSWORD"`
		MaxConn  string `env:"POSTGRES_MAX_CONN"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	cfg.GRPC.Port = os.Getenv("GRPC_PORT")
	cfg.GRPC.GatewayPort = os.Getenv("GRPC_GATEWAY_PORT")

	if cfg.GRPC.Port == "" {
		cfg.GRPC.Port = "50051"
	}
	if cfg.GRPC.GatewayPort == "" {
		cfg.GRPC.GatewayPort = "50052"
	}

	if err := checkPortAvailability(cfg.GRPC.Port); err != nil {
		log.Fatalf("Ошибка с портом GRPC: %v\n", err)
	}

	if err := checkPortAvailability(cfg.GRPC.GatewayPort); err != nil {
		log.Fatalf("Ошибка с портом Gateway: %v\n", err)
	}

	cfg.PG.Host = os.Getenv("POSTGRES_HOST")
	cfg.PG.Port = os.Getenv("POSTGRES_PORT")
	cfg.PG.DB = os.Getenv("POSTGRES_DB")
	cfg.PG.User = os.Getenv("POSTGRES_USER")
	cfg.PG.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.PG.MaxConn = os.Getenv("POSTGRES_MAX_CONN")

	cfg.PG.URL = "postgres://" + cfg.PG.User + ":" + cfg.PG.Password + "@" + cfg.PG.Host + ":" + cfg.PG.Port + "/" + cfg.PG.DB + "?sslmode=disable&pool_max_conns=" + cfg.PG.MaxConn

	return cfg, nil
}

func checkPortAvailability(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("порт %s уже занят", port)
	}
	defer listener.Close()

	return nil
}
