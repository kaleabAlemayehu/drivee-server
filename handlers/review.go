package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *handler) HandleGetAllReviews(w http.ResponseWriter, r *http.Request) {
	reviews, err := h.query.ListReviews(h.ctx)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if reviews == nil {
		log.Println("there is no review to list")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(reviews); err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
