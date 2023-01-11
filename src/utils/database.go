package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

var DbPool *pgxpool.Pool

func ConnectToPostgres() {
	var err error

	DbPool, err = pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/medium")
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
