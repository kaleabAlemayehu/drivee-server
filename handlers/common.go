package handlers

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/jackc/pgx/v5/pgxpool"
	model "github.com/kaleabAlemayehu/drivee-server/model"
)

type handler struct {
	query *model.Queries
	cache *memcache.Client
}

func NewHandler(conn *pgxpool.Pool, cache *memcache.Client) *handler {
	return &handler{query: model.New(conn), cache: cache}
}
