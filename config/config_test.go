package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	t.Setenv("GRPC_PORT", "50051")
	t.Setenv("GRPC_GATEWAY_PORT", "50052")
	t.Setenv("POSTGRES_HOST", "localhost")
	t.Setenv("POSTGRES_PORT", "5432")
	t.Setenv("POSTGRES_DB", "testdb")
	t.Setenv("POSTGRES_USER", "testuser")
	t.Setenv("POSTGRES_PASSWORD", "testpass")
	t.Setenv("POSTGRES_MAX_CONN", "10")

	cfg, err := NewConfig()
	require.NoError(t, err)

	assert.Equal(t, "50051", cfg.GRPC.Port)
	assert.Equal(t, "50052", cfg.GRPC.GatewayPort)
	assert.Equal(t, "localhost", cfg.PG.Host)
	assert.Equal(t, "5432", cfg.PG.Port)
	assert.Equal(t, "testdb", cfg.PG.DB)
	assert.Equal(t, "testuser", cfg.PG.User)
	assert.Equal(t, "testpass", cfg.PG.Password)
	assert.Equal(t, "10", cfg.PG.MaxConn)

	expectedURL := "postgres://testuser:testpass@localhost:5432/testdb?sslmode=disable&pool_max_conns=10"
	assert.Equal(t, expectedURL, cfg.PG.URL)
}
