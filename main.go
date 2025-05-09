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
	"github.com/kaleabAlemayehu/drivee-server/middleware"
	"github.com/rs/cors"
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
	pool := connection.DBConnect(os.Getenv("GOOSE_DBSTRING"))
	defer pool.Close()

	c := cors.New(cors.Options{
		// AllowedOrigins: []string{"http://localhost:5173"},
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions, http.MethodHead},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
		// Debug:            true,
	})

	handler := handlers.NewHandler(pool)
	secure := middleware.CreateStack(
		middleware.Auth,
	)

	secureFunc := middleware.CreateStackFunc(
		middleware.Auth,
	)
	free := middleware.CreateStack(
		c.Handler,
		middleware.Logger,
	)

	// INFO: user router (maybe i need to put it in its own package)
	userRouter := http.NewServeMux()
	// userRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodGet), handler.HandleGetAllUser)
	// userRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodPost), handler.HandleInsertUser)
	userRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodGet), handler.HandleGetUser)
	userRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodPatch), handler.HandleUpdateUser)

	// INFO: car router (maybe i need to put it in its own packaga)
	carRouter := http.NewServeMux()

	carRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodGet), handler.HandleGetAllCars)
	carRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodGet), handler.HandleGetCar)
	carRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodPatch), secureFunc(handler.HandleUpdateCar))
	carRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodPost), secureFunc(handler.HandleInsertCar))

	// INFO:
	bookingRouter := http.NewServeMux()
	bookingRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodGet), handler.HandleGetAllBookingsForRenter)

	bookingRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodGet), handler.HandleGetBookingByRenter)
	bookingRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodPost), handler.HandleInsertBooking)
	bookingRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodPatch), handler.HandleUpdateBookingByRenter)

	ownerRouter := http.NewServeMux()
	ownerRouter.HandleFunc(fmt.Sprintf("%s /bookings", http.MethodGet), handler.HandleGetAllBookingsForOwner)
	ownerRouter.HandleFunc(fmt.Sprintf("%s /bookings/{id}", http.MethodGet), handler.HandleGetBookingByOwner)
	ownerRouter.HandleFunc(fmt.Sprintf("%s /bookings/{id}/status", http.MethodPatch), handler.HandleUpdateBookingByOwner)
	ownerRouter.HandleFunc(fmt.Sprintf("%s /cars", http.MethodGet), handler.HandleGetAllCarsByOwner)
	ownerRouter.HandleFunc(fmt.Sprintf("%s /cars/{id}", http.MethodGet), handler.HandleGetCarByOwner)

	// INFO:
	paymentRouter := http.NewServeMux()
	paymentRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodGet), handler.HandleGetAllPayments)
	paymentRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodGet), handler.HandleGetPayment)
	paymentRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodPost), handler.HandleInsertPayment)
	paymentRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodPatch), handler.HandleUpdatePayment)

	// INFO:
	reviewRouter := http.NewServeMux()
	reviewRouter.HandleFunc(fmt.Sprintf("%s /user/{id}", http.MethodGet), handler.HandleGetAllReviews)
	reviewRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodGet), handler.HandleGetReview)
	reviewRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodPost), handler.HandleInsertReview)
	reviewRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodPatch), handler.HandleUpdateReview)

	// INFO:
	carPhotoRouter := http.NewServeMux()

	carPhotoRouter.HandleFunc(fmt.Sprintf("%s /car/{id}", http.MethodGet), handler.HandleGetAllCarPhotos)
	carPhotoRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodGet), handler.HandleGetCarPhoto)
	carPhotoRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodPost), secureFunc(handler.HandleInsertCarPhoto))
	carPhotoRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodPatch), secureFunc(handler.HandleUpdateCarPhoto))

	transactionRouter := http.NewServeMux()

	transactionRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodGet), handler.HandleGetAllTransactions)
	transactionRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodGet), handler.HandleGetTransaction)
	transactionRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodPost), handler.HandleCreateTransaction)
	transactionRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodPatch), handler.HandleUpdateTransaction)

	mux := http.NewServeMux()
	mux.HandleFunc(fmt.Sprintf("%s /api/register", http.MethodPost), handler.HandleRegister)
	mux.HandleFunc(fmt.Sprintf("%s /api/login", http.MethodPost), handler.HandleLogin)
	// TODO: may get rid of after move to cloudnary
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	mux.Handle("/api/users/", http.StripPrefix("/api/users", secure(userRouter)))
	mux.Handle("/api/cars/", http.StripPrefix("/api/cars", carRouter))
	mux.Handle("/api/bookings/", http.StripPrefix("/api/bookings", secure(bookingRouter)))
	mux.Handle("/api/owner/", http.StripPrefix("/api/owner", secure(ownerRouter)))
	mux.Handle("/api/payments/", http.StripPrefix("/api/payments", secure(paymentRouter)))
	mux.Handle("/api/reviews/", http.StripPrefix("/api/reviews", secure(reviewRouter)))
	mux.Handle("/api/carphotos/", http.StripPrefix("/api/carphotos", carPhotoRouter))
	mux.Handle("/api/transactions/", http.StripPrefix("/api/transactions", secure(transactionRouter)))

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: free(mux),
	}

	log.Printf("running server on localhost:%v", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("unable to run the server %v \n", err.Error())
	}
}
