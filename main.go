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

	// INFO: user router (maybe i need to put it in its own package)
	userRouter := http.NewServeMux()
	userRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodGet), handler.HandleGetAllUser)
	userRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodGet), handler.HandleGetUser)
	userRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodPatch), handler.HandleUpdateUser)
	userRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodPost), handler.HandleInsertUser)

	// INFO: car router (maybe i need to put it in its own packaga)
	carRouter := http.NewServeMux()
	carRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodGet), handler.HandleGetAllCars)
	carRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodGet), handler.HandleGetCar)
	carRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodPatch), handler.HandleUpdateCar)
	carRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodPost), handler.HandleInsertCar)

	// INFO:
	bookingRouter := http.NewServeMux()
	bookingRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodGet), handler.HandleGetAllBookings)
	bookingRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodGet), handler.HandleGetBooking)
	bookingRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodPost), handler.HandleInsertBooking)
	bookingRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodPatch), handler.HandleUpdateBooking)

	// INFO:

	paymentRouter := http.NewServeMux()
	paymentRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodGet), handler.HandleGetAllPayments)
	paymentRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodGet), handler.HandleGetPayment)

	mux := http.NewServeMux()
	mux.Handle("/api/users/", http.StripPrefix("/api/users", userRouter))
	mux.Handle("/api/cars/", http.StripPrefix("/api/cars", carRouter))
	mux.Handle("/api/bookings/", http.StripPrefix("/api/bookings", bookingRouter))
	mux.Handle("/api/payments/", http.StripPrefix("/api/payments", paymentRouter))

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	log.Printf("running server on localhost:%v", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("unable to run the server %v \n", err.Error())
	}
}
