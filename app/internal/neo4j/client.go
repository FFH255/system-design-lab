package neo4j

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Config struct {
	URI      string
	User     string
	Password string
}

func NewClient(ctx context.Context, cfg *Config) neo4j.DriverWithContext {
	driver, err := neo4j.NewDriverWithContext(
		cfg.URI,
		neo4j.BasicAuth(cfg.User, cfg.Password, ""),
	)

	if err != nil {
		panic(err)
	}

	err = driver.VerifyConnectivity(ctx)

	if err != nil {
		panic(err)
	}

	return driver
}
