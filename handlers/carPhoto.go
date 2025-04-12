package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/kaleabAlemayehu/drivee-server/model"
)
func (h *handler) HandleGetAllCarPhotos(w http.ResponseWriter, r *http.Request) {
	photos, err := h.query.GetCarPhotos(h.ctx)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(photos); err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *handler) HandleGetCarPhoto(w http.ResponseWriter, r *http.Request) {
