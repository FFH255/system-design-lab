package elastic_search

import "github.com/elastic/go-elasticsearch/v8"

type Config struct {
	Address string
}

func NewClient(cfg *Config) *elasticsearch.Client {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{cfg.Address},
	})

	if err != nil {
		panic(err)
	}

	return client
}
