package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *handler) HandleGetAllPayment(w http.ResponseWriter, r *http.Request) {
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
