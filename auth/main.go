package main

import (
	"authenticate/internal/auth"
	"fmt"
	"os"
)

var (
	pgaddr = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	raddr = fmt.Sprintf("%s:%s",
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
	)
	password = os.Getenv("REDIS_PASSWORD")

	key = os.Getenv("S_KEY")
)

func main() {
	srv := auth.NewServer(pgaddr, raddr, password, key)
	srv.StartUp()
}
