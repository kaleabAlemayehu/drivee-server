package connection

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func DBConnect(connectionString string) (context.Context, *pgx.Conn) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		log.Fatalf("unable to connect database: %s", connectionString)
	}
	return ctx, conn

}
