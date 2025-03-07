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
	mux := http.NewServeMux()
	mux.HandleFunc(fmt.Sprintf("%s /api/owner", http.MethodGet), handler.HandleGetAllOwner(ctx, conn))
	mux.HandleFunc(fmt.Sprintf("%s /api/owner/{id}", http.MethodGet), handler.HandleGetOwner(ctx, conn))
	// mux.HandleFunc(fmt.Sprintf("%s /api/owner/{id}", http.MethodGet), handlers.HandleGetAllOwner)

	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
