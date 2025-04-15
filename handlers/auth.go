package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	argon "github.com/alexedwards/argon2id"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kaleabAlemayehu/drivee-server/model"
)

type registerResponse struct {
	ID         uuid.UUID `json:"id"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name"`
	LastName   string    `json:"last_name"`
	Token      string    `json:"token"`
}

func (h *handler) HandleRegister(w http.ResponseWriter, r *http.Request) {

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
	hashedPass, err := argon.CreateHash(params.Password, argon.DefaultParams)
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
	// generate jwt token and attack to response
	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   owner.ID,
		"name":  owner.FirstName,
		"email": owner.Email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().AddDate(0, 0, 7).Unix(),
	})
	token, err := tokenStr.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var res registerResponse = registerResponse{
		ID:         owner.ID,
		Email:      owner.Email,
		FirstName:  owner.FirstName,
		MiddleName: owner.MiddleName.String,
		LastName:   owner.LastName,
		Token:      token,
	}

	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "unable to send data", http.StatusInternalServerError)
		return
	}
}
