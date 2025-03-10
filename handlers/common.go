package handlers

import (
	"context"
	"github.com/jackc/pgx/v5"
	model "github.com/kaleabAlemayehu/drivee-server/model"
)

type handler struct {
	query *model.Queries
	ctx   context.Context
}

func NewHandler(ctx context.Context, conn *pgx.Conn) *handler {
	query := model.New(conn)
	return &handler{query: query, ctx: ctx}
}
