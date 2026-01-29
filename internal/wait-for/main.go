package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	for i := 0; i < 10; i++ {
		err := tryConnecting()
		if err == nil {
			os.Exit(0)
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Println("Failed to connect")
	os.Exit(1)
}

func tryConnecting() error {
	return connectToPostgreSQL()
}

func connectToPostgreSQL() error {
	addr := os.Getenv("WATERMILL_TEST_POSTGRES_HOST")
	if addr == "" {
		addr = "localhost"
	}

	connStr := fmt.Sprintf("postgres://watermill:password@%s/watermill?sslmode=disable", addr)
	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping(context.Background())
	if err != nil {
		return err
	}

	return nil
}
