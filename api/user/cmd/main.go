package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"go-grpc-server/api/user"
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

	var r user.Repository
	r, err = user.NewPostgresRepository(cfg.DatabaseURL)
	if err != nil {
		log.Println(err)
	}
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := user.NewService(r)
	log.Fatal(user.ListenGRPC(s, 8080))
}
