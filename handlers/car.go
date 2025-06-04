package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kaleabAlemayehu/drivee-server/dto"
	"github.com/kaleabAlemayehu/drivee-server/model"
	"github.com/kaleabAlemayehu/drivee-server/utils"
)

func (h *handler) HandleGetAllCars(w http.ResponseWriter, r *http.Request) {
	var params model.ListCarsParams
	queries := r.URL.Query()
	if limit, err := strconv.Atoi(queries.Get("limit")); limit > 0 && err == nil {
		params.Limit = int32(limit)
		log.Printf("limit %v\n", params.Limit)
	} else {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "limit query is not provided correctly")
		return
	}

	if offset, err := strconv.Atoi(queries.Get("offset")); offset >= 0 && err == nil {
		params.Offset = int32(offset)
		log.Printf("offset %v\n", params.Limit)
	} else {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "offset query is not provided correctly")
		return
	}
	cars, err := h.query.ListCars(r.Context(), params)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to fetch cars")
		return
	}

	var response dto.CarResponse
	response.Total, err = h.query.GetCarsCount(r.Context())
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to fetch total car number")
		return
	}
	response.Cars = cars

	if err = utils.SendResponse(w, "success", http.StatusOK, response); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send data")
		return
	}
}

func (h *handler) HandleGetCar(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	car, err := h.query.GetCar(r.Context(), id)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	if err = utils.SendResponse(w, "success", http.StatusOK, car); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send data")
		return
	}
}

func (h *handler) HandleInsertCar(w http.ResponseWriter, r *http.Request) {
	ownerID, err := uuid.Parse(r.Context().Value("userID").(string))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	var body struct {
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
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	params := model.InsertCarParams{
		OwnerID:       ownerID,
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

	car, err := h.query.InsertCar(r.Context(), params)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	if err = utils.SendResponse(w, "succes", http.StatusCreated, car); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send data")
		return
	}
}

func (h *handler) HandleUpdateCar(w http.ResponseWriter, r *http.Request) {
	ownerID, err := uuid.Parse(r.Context().Value("userID").(string))
	// TODO: fix this
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusUnauthorized, "unauthorize")
		return
	}
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send data")
		return
	}

	// FIXME: put it on dto
	var body struct {
		Mileage  int32 `json:"mileage"`
		Location struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		}
		PricePerHour pgtype.Numeric   `json:"price_per_hour"`
		Status       model.StatusType `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	payload := model.UpdateCarParams{
		ID:            id,
		Mileage:       body.Mileage,
		StMakepoint:   body.Location.X,
		StMakepoint_2: body.Location.Y,
		PricePerHour:  body.PricePerHour,
		Status:        model.StatusType(body.Status),
		OwnerID:       ownerID,
	}

	car, err := h.query.UpdateCar(r.Context(), payload)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	if err := utils.SendResponse(w, "succes", http.StatusOK, car); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send data")
		return
	}
}
