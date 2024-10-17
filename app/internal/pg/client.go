package pg

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Config struct {
	URL string
}

func NewClient(cfg *Config) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), cfg.URL)

	if err != nil {
		panic(err)
	}

	return conn
}
