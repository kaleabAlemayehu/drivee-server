package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/cridenour/go-postgis"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kaleabAlemayehu/drivee-server/model"
)

type carParams struct {
	id             uuid.UUID      `json:"id"`
	ownerid        uuid.UUID      `json:"owner_id"`
	make           string         `json:"make"`
	model          string         `json:"model"`
	year           string         `json:"year"`
	license_plate  string         `json:"license_plate"`
	vin_number     string         `json:"vin_number"`
	transmission   string         `json:"transmission"`
	fuel_type      model.FuelType `json:"fuel_type"`
	mileage        int32          `json:"mileage"`
	location       postgis.PointS `json:"location"`
	price_per_hour pgtype.Numeric `json:"price_per_hour"`
}

func (h *handler) HandleGetAllCars(w http.ResponseWriter, r *http.Request) {

	cars, err := h.query.ListCars(h.ctx)
	if err != nil {
		log.Println("unable to fetch cars")
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&cars)
	if err != nil {
		log.Println("unable to fetch cars")
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *handler) HandleGetCar(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	car, err := h.query.GetCar(h.ctx, id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(&car)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *handler) HandleInsertCar(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		log.Println("unable to marshal body")
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	var body struct {
		OwnerID      uuid.UUID          `json:"owner_id"`
		Make         string             `json:"make"`
		Model        string             `json:"model"`
		Year         string             `json:"year"`
		LicensePlate string             `json:"license_plate"`
		VinNumber    string             `json:"vin_number"`
		Transmission model.Transmission `json:"transmission"`
		FuelType     model.FuelType     `json:"fuel_type"`
		Mileage      int32              `json:"mileage"`
		Location     struct {
			X float64 `json:"X"`
			Y float64 `json:"Y"`
		} `json:"location"`
		PricePerHour pgtype.Numeric   `json:"price_per_hour"`
		Status       model.StatusType `json:"status"`
	}
	err = json.Unmarshal(bodyByte, &body)
	log.Printf(" \n \n struct %v \n", body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	params := model.InsertCarParams{
		OwnerID:       body.OwnerID,
		Make:          body.Make,
		Model:         body.Model,
		Year:          body.Year,
		LicensePlate:  body.LicensePlate,
		VinNumber:     body.VinNumber,
		Transmission:  body.Transmission,
		FuelType:      body.FuelType,
		Mileage:       body.Mileage,
		StMakepoint:   body.Location.X,
		StMakepoint_2: body.Location.Y,
		PricePerHour:  body.PricePerHour,
		Status:        body.Status,
	}

	car, err := h.query.InsertCar(h.ctx, params)
	if err != nil {
		log.Println(err.Error())
		log.Println("the query is not working")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(&car)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
