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
	userRouter.HandleFunc("GET /{id}", handler.HandleGetUser)
	userRouter.HandleFunc("PATCH /", handler.HandleUpdateUser)

	// INFO: car router (maybe i need to put it in its own packaga)
	carRouter := http.NewServeMux()

	carRouter.HandleFunc("GET /", handler.HandleGetAllCars)
	carRouter.HandleFunc("GET /{id}", handler.HandleGetCar)
	carRouter.HandleFunc("PATCH /{id}", secureFunc(handler.HandleUpdateCar))
	carRouter.HandleFunc("POST /", secureFunc(handler.HandleInsertCar))

	// INFO:
	bookingRouter := http.NewServeMux()
	bookingRouter.HandleFunc("GET /", handler.HandleGetAllBookingsForRenter)

	bookingRouter.HandleFunc("GET /{id}", handler.HandleGetBookingByRenter)
	bookingRouter.HandleFunc("POST /", handler.HandleInsertBooking)
	bookingRouter.HandleFunc("PATCH /{id}", handler.HandleUpdateBookingByRenter)

	ownerRouter := http.NewServeMux()
	ownerRouter.HandleFunc("GET /bookings", handler.HandleGetAllBookingsForOwner)
	ownerRouter.HandleFunc("GET /bookings/{id}", handler.HandleGetBookingByOwner)
	ownerRouter.HandleFunc("PATCH /bookings/{id}/status", handler.HandleUpdateBookingByOwner)
	ownerRouter.HandleFunc("GET /cars", handler.HandleGetAllCarsByOwner)
	ownerRouter.HandleFunc("GET /cars/{id}", handler.HandleGetCarByOwner)

	// INFO:
	paymentRouter := http.NewServeMux()
	paymentRouter.HandleFunc("GET /", handler.HandleGetAllPayments)
	paymentRouter.HandleFunc("GET /{id}", handler.HandleGetPayment)
	paymentRouter.HandleFunc("POST /", handler.HandleInsertPayment)
	paymentRouter.HandleFunc("PATCH /{id}", handler.HandleUpdatePayment)

	// INFO:
	reviewRouter := http.NewServeMux()
	reviewRouter.HandleFunc("GET /user/{id}", handler.HandleGetAllReviews)
	reviewRouter.HandleFunc("GET /{id}", handler.HandleGetReview)
	reviewRouter.HandleFunc("POST /", handler.HandleInsertReview)
	reviewRouter.HandleFunc("PATCH /{id}", handler.HandleUpdateReview)

	// INFO:
	carPhotoRouter := http.NewServeMux()

	carPhotoRouter.HandleFunc("GET /car/{id}", handler.HandleGetAllCarPhotos)
	carPhotoRouter.HandleFunc("GET /{id}", handler.HandleGetCarPhoto)
	carPhotoRouter.HandleFunc("POST /", secureFunc(handler.HandleInsertCarPhoto))
	carPhotoRouter.HandleFunc("PATCH /{id}", secureFunc(handler.HandleUpdateCarPhoto))

	transactionRouter := http.NewServeMux()

	transactionRouter.HandleFunc("GET /", handler.HandleGetAllTransactions)
	transactionRouter.HandleFunc("GET /{id}", handler.HandleGetTransaction)
	transactionRouter.HandleFunc("POST /", handler.HandleCreateTransaction)
	transactionRouter.HandleFunc("PATCH /{id}", handler.HandleUpdateTransaction)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/register", handler.HandleRegister)
	mux.HandleFunc("POST /api/login", handler.HandleLogin)
	mux.HandleFunc("POST /api/auth/google", handler.HandleGoogleAuth)
	mux.HandleFunc("GET /api/auth/x/pkce", handler.GeneratePKCE)
	mux.HandleFunc("POST /api/request-reset", handler.HandleRequestReset)
	mux.HandleFunc("POST /api/reset-password", handler.HandleResetPassword)
	mux.HandleFunc("POST /api/verify-token", handler.HandleVerifyToken)

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
