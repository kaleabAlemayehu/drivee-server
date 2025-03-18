package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *handler) HandleGetAllBookings(w http.ResponseWriter, r *http.Request) {
	cars, err := h.query.ListBookings(h.ctx)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if cars == nil {
		log.Println("there is not booking to list")
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("[]"))
		return
	}

	err = json.NewEncoder(w).Encode(cars)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
