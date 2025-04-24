package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/kaleabAlemayehu/drivee-server/model"
	"github.com/kaleabAlemayehu/drivee-server/utils"
)

func (h *handler) HandleGetAllCarPhotos(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	photos, err := h.query.GetCarPhotos(h.ctx, id)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}

	if err := utils.SendResponse(w, "success", http.StatusOK, photos); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send data")
		return
	}
}

func (h *handler) HandleGetCarPhoto(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	photo, err := h.query.GetCarPhoto(h.ctx, id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(photo); err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *handler) HandleInsertCarPhoto(w http.ResponseWriter, r *http.Request) {
	var body model.InsertCarPhotoParams
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	photo, err := h.query.InsertCarPhoto(h.ctx, body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(photo); err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *handler) HandleUpdateCarPhoto(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	var body model.UpdateCarPhotoParams
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	body.ID = id

	photo, err := h.query.UpdateCarPhoto(h.ctx, body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(photo); err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusBadRequest)
		return
	}
}
