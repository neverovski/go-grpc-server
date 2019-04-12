package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"go-grpc-server/api/news"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var n news.Repository
	n, err = news.NewPostgresRepository(cfg.DatabaseURL)
	if err != nil {
		log.Println(err)
	}
	defer n.Close()

	log.Println("Listening on port 8080...")
	s := news.NewService(n)
	log.Fatal(news.ListenGRPC(s, 8080))
}
