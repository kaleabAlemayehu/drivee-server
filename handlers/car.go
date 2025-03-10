package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

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
