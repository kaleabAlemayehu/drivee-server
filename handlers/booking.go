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
	log.Printf("uid:%v\n", r.Context().Value("userID"))
	ownerID, err := uuid.Parse(r.Context().Value("userID").(string))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	bookings, err := h.query.ListBookingsForOwner(h.ctx, ownerID)
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
	log.Printf("uid:%v\n", r.Context().Value("userID"))
	renterID, err := uuid.Parse(r.Context().Value("userID").(string))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	bookings, err := h.query.ListBookingsForRenter(h.ctx, renterID)
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

func (h *handler) HandleGetBooking(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	booking, err := h.query.GetBooking(h.ctx, id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(booking)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *handler) HandleInsertBooking(w http.ResponseWriter, r *http.Request) {
	var body model.InsertBookingParams
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	booking, err := h.query.InsertBooking(h.ctx, body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(booking); err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *handler) HandleUpdateBooking(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	var body model.UpdateBookingParams
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	body.ID = id
	booking, err := h.query.UpdateBooking(h.ctx, body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(booking); err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	return
}
