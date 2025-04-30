package connection

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DBConnect(connectionString string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("unable to connect database: %s", connectionString)
	}
	log.Printf("database connected succesfully: %s", connectionString)
	return pool
}
