package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	connection "github.com/kaleabAlemayehu/drivee-server/connection"
	"github.com/kaleabAlemayehu/drivee-server/handlers"
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

	handler := handlers.NewHandler(ctx, conn)
	owner := http.NewServeMux()
	owner.HandleFunc(fmt.Sprintf("%s /", http.MethodGet), handler.HandleGetAllOwner)
	owner.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodGet), handler.HandleGetOwner)
	owner.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodPatch), handler.HandleUpdateOwner)
	owner.HandleFunc(fmt.Sprintf("%s /", http.MethodPost), handler.HandleInsertOwner)

	mux := http.NewServeMux()
	mux.Handle("/api/owner/", http.StripPrefix("/api/owner", owner))

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("unable to run the server %v \n", err.Error())
	}
}
