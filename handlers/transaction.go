package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *handler) HandleGetAllTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.query.ListTransactions(h.ctx)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(transactions); err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

