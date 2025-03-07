package handlers

import (
	"context"
	"io"
	"net/http"

	"encoding/json"
	"log"

	"github.com/alexedwards/argon2id"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	model "github.com/kaleabAlemayehu/drivee-server/model"
)

type handler struct {
	query *model.Queries
	ctx   context.Context
}

func NewHandler(ctx context.Context, conn *pgx.Conn) *handler {
	query := model.New(conn)
	return &handler{query: query, ctx: ctx}
}

func (h *handler) HandleGetAllOwner(w http.ResponseWriter, r *http.Request) {

	owners, err := h.query.ListOwner(h.ctx)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "unable to fetch owner list", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&owners)
	if err != nil {
		log.Println("unable to send owner list")
		http.Error(w, "unable to send owner list", http.StatusInternalServerError)
		return
	}
}

func (h *handler) HandleGetOwner(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println("unable to parse uuid that sent")
		http.Error(w, "bad request", http.StatusBadRequest)
		return

	}
	owner, err := h.query.GetOwner(h.ctx, id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "unable to fetch owner", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&owner)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "unable to send owner", http.StatusInternalServerError)
		return
	}
}

func (h *handler) HandleInsertOwner(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	var params model.InsertOwnerParams
	err = json.Unmarshal(body, &params)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	// hash the password
	hashedPass, err := argon2id.CreateHash(params.Password, argon2id.DefaultParams)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusBadRequest)
		return
	}
	// insert into the table
	owner, err := h.query.InsertOwner(h.ctx, model.InsertOwnerParams{
		FirstName:     params.FirstName,
		MiddleName:    params.MiddleName,
		LastName:      params.LastName,
		Email:         params.Email,
		Password:      hashedPass,
		PhoneNumber:   params.PhoneNumber,
		AccountNumber: params.AccountNumber,
		BankName:      params.BankName,
	})
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(&owner)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "unable to send data", http.StatusInternalServerError)
		return
	}

}
