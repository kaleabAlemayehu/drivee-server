package handlers

import (
	"context"
	"net/http"

	"encoding/json"
	"log"

	"github.com/google/uuid"
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

func (h *handler) HandleGetAllOwner(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO:  fetch owner form database

		owners, err := h.query.ListOwner(h.ctx)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "unable to fetch owner list", http.StatusInternalServerError)
			return
		}

		// TODO:  sender owner  to the client
		err = json.NewEncoder(w).Encode(&owners)
		if err != nil {
			log.Println("unable to send owner list")
			http.Error(w, "unable to send owner list", http.StatusInternalServerError)
			return
		}
	}
}
func (h *handler) HandleGetOwner(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: fetch owner form database
		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			log.Println("unable to parse uuid that sent")
			http.Error(w, "bad request", http.StatusBadRequest)
			return

		}
		owner, err := h.query.GetOwner(h.ctx, id)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "unable to fetch owner", http.StatusInternalServerError)
			return
		}

		// TODO: sender owner  to the client
		err = json.NewEncoder(w).Encode(&owner)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "unable to send owner", http.StatusInternalServerError)
			return
		}
	}
}
