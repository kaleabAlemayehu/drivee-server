package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/kaleabAlemayehu/drivee-server/model"
	"github.com/kaleabAlemayehu/drivee-server/utils"
)

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

func (h *handler) HandleGetAllBookingsForRenter(w http.ResponseWriter, r *http.Request) {
	renterID, err := uuid.Parse(r.Context().Value("userID").(string))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusUnauthorized, "unauthorized")
		return
	}
	bookings, err := h.query.ListBookingsForRenter(r.Context(), renterID)
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

func (h *handler) HandleGetBookingByRenter(w http.ResponseWriter, r *http.Request) {
	renterID, err := uuid.Parse(r.Context().Value("userID").(string))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusForbidden, "unauthorized")
		return
	}
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	params := model.GetBookingForRenterParams{
		ID:       id,
		RenterID: renterID,
	}
	booking, err := h.query.GetBookingForRenter(r.Context(), params)
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

func (h *handler) HandleInsertBooking(w http.ResponseWriter, r *http.Request) {
	renterID, err := uuid.Parse(r.Context().Value("userID").(string))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusForbidden, "unauthorized")
		return
	}
	var body model.InsertBookingParams
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	body.RenterID = renterID
	booking, err := h.query.InsertBooking(r.Context(), body)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	if err := utils.SendResponse(w, "success", http.StatusOK, booking); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send created data")
		return
	}
}

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

func (h *handler) HandleUpdateBookingByRenter(w http.ResponseWriter, r *http.Request) {
	renterID, err := uuid.Parse(r.Context().Value("userID").(string))
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
	var body model.UpdateBookingForRenterParams
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	body.ID = id
	body.RenterID = renterID
	booking, err := h.query.UpdateBookingForRenter(r.Context(), body)
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
