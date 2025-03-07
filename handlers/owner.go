package handlers

import (
	"net/http"

	"encoding/json"
	connection "github.com/kaleabAlemayehu/drivee-server/connection"
	model "github.com/kaleabAlemayehu/drivee-server/model"
	"log"
	"os"
)

func HandleGetAllOwner(w http.ResponseWriter, r *http.Request) {
	ctx, conn := connection.DBConnect(os.Getenv("GOOSE_DBSTRING"))
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
	return
}
