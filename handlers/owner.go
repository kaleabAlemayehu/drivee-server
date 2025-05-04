package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/kaleabAlemayehu/drivee-server/model"
	"github.com/kaleabAlemayehu/drivee-server/utils"
)

func (h *handler) HandleUpdateBookingByOwner(w http.ResponseWriter, r *http.Request) {
	ownerID, err := uuid.Parse(r.Context().Value("userID").(string))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}

	var body model.UpdateBookingForOwnerParams
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	body.ID = id
	body.OwnerID = ownerID
	booking, err := h.query.UpdateBookingForOwner(r.Context(), body)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	if err := utils.SendResponse(w, "success", http.StatusOK, booking); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send data")
		return
	}
	return
}

func (h *handler) HandleGetBookingByOwner(w http.ResponseWriter, r *http.Request) {
	ownerID, err := uuid.Parse(r.Context().Value("userID").(string))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusUnauthorized, "unauthorized")
		return
	}
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	params := model.GetBookingForOwnerParams{
		ID:      id,
		OwnerID: ownerID,
	}
	booking, err := h.query.GetBookingForOwner(r.Context(), params)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	if err = utils.SendResponse(w, "success", http.StatusOK, booking); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send date")
		return
	}
}

func (h *handler) HandleGetAllBookingsForOwner(w http.ResponseWriter, r *http.Request) {
	ownerID, err := uuid.Parse(r.Context().Value("userID").(string))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusUnauthorized, "unauthorized")
		return
	}
	bookings, err := h.query.ListBookingsForOwner(r.Context(), ownerID)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to fetch all bookings")
		return
	}
	if bookings == nil {
		utils.SendResponse(w, "success", http.StatusNoContent, "[]")
		log.Println("there is not booking to list")
		return
	}

	if err = utils.SendResponse(w, "success", http.StatusOK, bookings); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send bookings data")
		return
	}
}

func (h *handler) HandleGetAllCarsByOwner(w http.ResponseWriter, r *http.Request) {
	ownerID, err := uuid.Parse(r.Context().Value("userID").(string))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusUnauthorized, "unauthorized")
		return
	}
	cars, err := h.query.ListCarsForOwner(r.Context(), ownerID)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to fetch all cars")
		return
	}
	if cars == nil {
		utils.SendResponse(w, "success", http.StatusNoContent, "[]")
		log.Println("there is not cars to list")
		return
	}
	if err = utils.SendResponse(w, "success", http.StatusOK, cars); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send cars data")
		return
	}
}

func (h *handler) HandleGetCarByOwner(w http.ResponseWriter, r *http.Request) {
	ownerID, err := uuid.Parse(r.Context().Value("userID").(string))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusUnauthorized, "unauthorized")
		return
	}
	carId, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	var params = model.GetCarForOwnerParams{
		ID:      carId,
		OwnerID: ownerID,
	}
	car, err := h.query.GetCarForOwner(r.Context(), params)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to fetch a car")
		return
	}
	if car == nil {
		utils.SendResponse(w, "success", http.StatusNoContent, "[]")
		log.Println("there is no car with this id")
		return
	}
	if err = utils.SendResponse(w, "success", http.StatusOK, car); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send the car data")
		return
	}
}
