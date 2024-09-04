package config

import (
	"flag"
	"os"
)

type Config struct {
	Port int
	Env  string
	Api  string
	Db   struct {
		Dsn string
	}
	Stripe struct {
		Secret string
		Key    string
	}
	Secretkey string
	Frontend  string
}

func Read() Config {
	var cfg Config

	flag.IntVar(&cfg.Port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "development", "Application environment {development|production}")
	flag.StringVar(&cfg.Db.Dsn, "dsn", "root:root@tcp(mariadb:3306)/testdb", "DSN for MariaDB")
	flag.StringVar(&cfg.Api, "api", "http://localhost:4001", "URL to api")
	flag.StringVar(&cfg.Secretkey, "secret", "bRWmrwNUTqNUuzckjxsFlHZjxHkjrzKP", "secret key")
	flag.StringVar(&cfg.Frontend, "frontend", "http://localhost:4000", "url to front end")

	flag.Parse()

	cfg.Stripe.Key = os.Getenv("STRIPE_KEY")
	cfg.Stripe.Secret = os.Getenv("STRIPE_SECRET")

	return cfg
}
