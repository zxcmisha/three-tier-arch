package repository

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connection(ctx context.Context) (*pgx.Conn, error) {
	connstr := os.Getenv("DATABASE_URL")
	return pgx.Connect(ctx, connstr)
}
