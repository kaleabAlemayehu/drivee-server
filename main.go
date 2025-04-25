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
	stack := middleware.CreateStack(
		middleware.Logger,
		middleware.Auth,
	)
	_ = stack

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
	carRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodPatch), middleware.ProtectedHandleFunc(stack, handler.HandleUpdateCar))
	carRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodPost), middleware.ProtectedHandleFunc(stack, handler.HandleInsertCar))

	// INFO:
	bookingRouter := http.NewServeMux()
	bookingRouter.HandleFunc(fmt.Sprintf("%s /owner/", http.MethodGet), handler.HandleGetAllBookingsForOwner)
	bookingRouter.HandleFunc(fmt.Sprintf("%s /renter/", http.MethodGet), handler.HandleGetAllBookingsForRenter)

	bookingRouter.HandleFunc(fmt.Sprintf("%s /owner/{id}", http.MethodGet), handler.HandleGetBookingByOwner)
	bookingRouter.HandleFunc(fmt.Sprintf("%s /renter/{id}", http.MethodGet), handler.HandleGetBookingByRenter)
	bookingRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodPost), handler.HandleInsertBooking)
	bookingRouter.HandleFunc(fmt.Sprintf("%s /owner/{id}", http.MethodPatch), handler.HandleUpdateBookingByOwner)
	bookingRouter.HandleFunc(fmt.Sprintf("%s /renter/{id}", http.MethodPatch), handler.HandleUpdateBookingByRenter)

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

	carPhotoRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodGet), handler.HandleGetAllCarPhotos)
	carPhotoRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodGet), handler.HandleGetCarPhoto)
	carPhotoRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodPost), middleware.ProtectedHandleFunc(stack, handler.HandleInsertCarPhoto))
	carPhotoRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodPatch), middleware.ProtectedHandleFunc(stack, handler.HandleUpdateCarPhoto))

	transactionRouter := http.NewServeMux()

	transactionRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodGet), handler.HandleGetAllTransactions)
	transactionRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodGet), handler.HandleGetTransaction)
	transactionRouter.HandleFunc(fmt.Sprintf("%s /", http.MethodPost), handler.HandleCreateTransaction)
	transactionRouter.HandleFunc(fmt.Sprintf("%s /{id}", http.MethodPatch), handler.HandleUpdateTransaction)

	mux := http.NewServeMux()
	mux.HandleFunc(fmt.Sprintf("%s /api/register", http.MethodPost), handler.HandleRegister)
	mux.HandleFunc(fmt.Sprintf("%s /api/login", http.MethodPost), handler.HandleLogin)
	mux.Handle("/api/users/", http.StripPrefix("/api/users", stack(userRouter)))
	mux.Handle("/api/cars/", http.StripPrefix("/api/cars", carRouter))
	mux.Handle("/api/bookings/", http.StripPrefix("/api/bookings", stack(bookingRouter)))
	mux.Handle("/api/payments/", http.StripPrefix("/api/payments", stack(paymentRouter)))
	mux.Handle("/api/reviews/", http.StripPrefix("/api/reviews", stack(reviewRouter)))
	mux.Handle("/api/carphotos/", http.StripPrefix("/api/carphotos", carPhotoRouter))
	mux.Handle("/api/transactions/", http.StripPrefix("/api/transactions", stack(transactionRouter)))

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	log.Printf("running server on localhost:%v", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("unable to run the server %v \n", err.Error())
	}
}
