package database

import (
	"online-shop/internal/config"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
}

func TestConnectionPostgres(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, err := ConnectPostgres(config.Cfg.DB)
		require.Nil(t, err)
		require.NotNil(t, db)
	})
}
