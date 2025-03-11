package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
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
