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

func HandleGetAllOwner(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO:  fetch owner form database
		query := model.New(conn)
		owners, err := query.ListOwner(ctx)
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
func HandleGetOwner(ctx context.Context, conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: fetch owner form database
		query := model.New(conn)
		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			log.Println("unable to parse uuid that sent")
			http.Error(w, "bad request", http.StatusBadRequest)
			return

		}
		owner, err := query.GetOwner(ctx, id)
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
