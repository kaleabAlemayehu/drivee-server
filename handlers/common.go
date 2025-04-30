package handlers

import (
	"github.com/jackc/pgx/v5/pgxpool"
	model "github.com/kaleabAlemayehu/drivee-server/model"
)

type handler struct {
	query *model.Queries
}

func NewHandler(conn *pgxpool.Pool) *handler {
	return &handler{query: model.New(conn)}
}
