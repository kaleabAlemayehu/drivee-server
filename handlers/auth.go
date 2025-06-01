package handlers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	argon "github.com/alexedwards/argon2id"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/kaleabAlemayehu/drivee-server/dto"
	"github.com/kaleabAlemayehu/drivee-server/model"
	"github.com/kaleabAlemayehu/drivee-server/utils"
	"github.com/kaleabAlemayehu/identicon"
)

func (h *handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var params model.InsertUserParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "all inputs are needed")
		return
	}
	// hash the password
	hashedPass, err := argon.CreateHash(params.Password, argon.DefaultParams)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "internal server error")
		return
	}
	img := identicon.New7X7().Render([]byte(fmt.Sprintf(params.Email, params.FirstName)))
	profPic, err := utils.Upload(img, fmt.Sprintf("%s%s", params.FirstName, rand.Text()))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "internal server error")
		return
	}

	// insert into the table
	user, err := h.query.InsertUser(r.Context(), model.InsertUserParams{
		FirstName:      params.FirstName,
		Email:          params.Email,
		Password:       hashedPass,
		ProfilePicture: profPic,
	})
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "email is already taken")
		return
	}
	// generate jwt token and attack to response
	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"name":  user.FirstName,
		"email": user.Email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().AddDate(0, 0, 7).Unix(),
	})
	token, err := tokenStr.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "internal server error")
		return
	}

	var res dto.AuthResponse = dto.AuthResponse{
		ID:             user.ID,
		Email:          user.Email,
		FirstName:      user.FirstName,
		MiddleName:     user.MiddleName.String,
		ProfilePicture: user.ProfilePicture,
		Token:          token,
	}

	if err := utils.SendResponse(w, "success", http.StatusCreated, res); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send data")
		return
	}
}

func (h *handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var body dto.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "all inputs are needed")
		return
	}

	user, err := h.query.GetUserByEmail(r.Context(), body.Email)
	if err != nil {
		log.Println(err)
		utils.SendResponse(w, "error", http.StatusBadRequest, "invalid email or password")
		return
	}

	match, err := argon.ComparePasswordAndHash(body.Password, user.Password)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "invalid email or password")
		return
	}

	if !match {
		log.Println("password don't match")
		utils.SendResponse(w, "error", http.StatusBadRequest, "invalid email or password")
		return
	}
	var date int
	if body.RememberMe {
		date = 30

	} else {

		date = 7
	}
	// return new token and send it will the response
	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"name":  user.FirstName,
		"email": user.Email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().AddDate(0, 0, date).Unix(),
	})
	token, err := tokenStr.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "internal server error")
		return
	}

	var res dto.AuthResponse = dto.AuthResponse{
		ID:             user.ID,
		Email:          user.Email,
		FirstName:      user.FirstName,
		MiddleName:     user.MiddleName.String,
		Token:          token,
		ProfilePicture: user.ProfilePicture,
	}

	if err := utils.SendResponse(w, "success", int(http.StatusOK), res); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send data")
		return
	}

}
