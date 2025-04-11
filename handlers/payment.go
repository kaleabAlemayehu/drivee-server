package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (h *handler) HandleGetAllPayments(w http.ResponseWriter, r *http.Request) {
	payments, err := h.query.ListPayments(h.ctx)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if payments == nil {
		log.Println("there is not booking to list")
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("[]"))
		return
	}
	if err := json.NewEncoder(w).Encode(payments); err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *handler) HandleGetPayment(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	payment, err := h.query.GetPayment(h.ctx, id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(payment); err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
