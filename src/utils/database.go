package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

var DbPool *pgxpool.Pool

func ConnectToPostgres() {
	var err error
	DSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_ADDR"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DbPool, err = pgxpool.New(context.Background(), DSN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	err = DbPool.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		os.Exit(1)
	} else {
		log.Println("Connected to postgres")
	}
}
