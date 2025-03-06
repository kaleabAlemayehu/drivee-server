package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	// "github.com/google/uuid"
	// "github.com/google/uuid"
	"github.com/joho/godotenv"
	connection "github.com/kaleabAlemayehu/drivee-server/connection"
	"github.com/kaleabAlemayehu/drivee-server/model"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("unable to load environment variables")
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil || port == 0 {
		port = 9999
		log.Printf("port is not specified\n running on default port :%d \n", port)
	}
	ctx, conn := connection.DBConnect(os.Getenv("GOOSE_DBSTRING"))
	mux := http.NewServeMux()
	mux.HandleFunc(fmt.Sprintf("%s /api/owner", http.MethodGet), func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
