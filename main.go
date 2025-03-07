package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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
	mux := http.NewServeMux()
	mux.HandleFunc(fmt.Sprintf("%s /api/owner", http.MethodGet), handlers.HandleGetAllOwner)

	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
