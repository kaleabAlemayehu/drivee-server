package handlers

import (
	"io"
	"net/http"
	"reflect"

	"encoding/json"
	"log"

	"github.com/alexedwards/argon2id"
	"github.com/google/uuid"

	model "github.com/kaleabAlemayehu/drivee-server/model"
)

func (h *handler) HandleGetAllUser(w http.ResponseWriter, r *http.Request) {

	owners, err := h.query.ListUser(h.ctx)
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

func (h *handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println("unable to parse uuid that sent")
		http.Error(w, "bad request", http.StatusBadRequest)
		return

	}
	owner, err := h.query.GetUser(h.ctx, id)
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

func (h *handler) HandleInsertUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	var params model.InsertUserParams
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
	owner, err := h.query.InsertUser(h.ctx, model.InsertUserParams{
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

func (h *handler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println("unable to get id parameter")
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("unable to get body for the request")
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	var Body model.UpdateUserParams
	err = json.Unmarshal(body, &Body)
	if err != nil {
		log.Println("unable to unmarshal body from the request bytes")
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	Body.ID = id
	v := reflect.ValueOf(Body)
	for i := range v.NumField() {
		if v.Field(i).Interface() == "" {
			log.Println("unable to unmarshal body from the request bytes")
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
	}
	res, err := h.query.UpdateUser(h.ctx, Body)
	if err != nil {
		log.Println("the query is not excuted")
		http.Error(w, "bad request body", http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		log.Println("unable to encode and send response")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
