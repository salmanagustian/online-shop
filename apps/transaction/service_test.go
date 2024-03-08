package transaction

import (
	"context"
	"online-shop/external/database"
	"online-shop/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	repo := newRepository(db)
	svc = newService(repo)
}

func TestCreateTransaction(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := CreateTransactionRequestPayload{
			ProductSKU:   "d7edce37-e2c0-4335-9212-a76458ab1a22",
			Amount:       2,
			UserPublicId: "faf611ae-5a70-4a45-9b43-875cf6366377",
		}

		err := svc.CreateTransaction(context.Background(), req)
		require.Nil(t, err)
	})
}
