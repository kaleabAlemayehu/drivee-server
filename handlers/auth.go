package handlers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	argon "github.com/alexedwards/argon2id"
	"github.com/jackc/pgx/v5/pgtype"
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
	hashedPass, err := argon.CreateHash(params.Password.String, argon.DefaultParams)
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
		Password:       pgtype.Text{String: hashedPass, Valid: true},
		ProfilePicture: profPic,
	})
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "email is already taken")
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.FirstName, user.Email, false)
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

	match, err := argon.ComparePasswordAndHash(body.Password, user.Password.String)
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
	token, err := utils.GenerateJWT(user.ID, user.FirstName, user.Email, true)
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

func (h *handler) HandleRequestReset(w http.ResponseWriter, r *http.Request) {
	// TODO: add rate limiter
	var input dto.ResetRequestInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "We'll send a reset email if the account exists")
		return
	}

	user, err := h.query.GetUserByEmail(r.Context(), input.Email)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "We'll send a reset email if the account exists")
		return
	}
	token, hashed, err := utils.GenerateToken()
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "internal server error")
		return
	}

	var params = model.InsertTokenParams{
		Token:     hashed,
		ExpiresAt: int32(time.Now().Add(time.Hour * 6).Unix()),
		UserID:    user.ID,
	}

	_, err = h.query.InsertToken(r.Context(), params)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "internal server error")
		return
	}

	clientAddr := os.Getenv("CLIENT_ADDR")
	url := fmt.Sprintf("%s/reset-password/?token=%s", clientAddr, token)

	if err := utils.SendEmail(user.Email, user.FirstName, url); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "internal server error")
		return
	}
	utils.SendResponse(w, "success", http.StatusOK, "We'll send a reset email if the account exists")
	return
}

func (h *handler) HandleVerifyToken(w http.ResponseWriter, r *http.Request) {
	var body dto.VerifyTokenInput
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	hash := sha256.Sum256([]byte(body.Token))
	hashed := hex.EncodeToString(hash[:])
	param := model.GetTokenParams{
		Token:     hashed,
		ExpiresAt: int32(time.Now().Unix()),
	}
	token, err := h.query.GetToken(r.Context(), param)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "token expired request a new one")
		return
	}
	now := int32(time.Now().Unix())
	if token.ExpiresAt <= now {
		utils.SendResponse(w, "error", http.StatusBadRequest, "token expired request a new one")
		return
	}

	utils.SendResponse(w, "success", http.StatusOK, token)
	return
}

func (h *handler) HandleResetPassword(w http.ResponseWriter, r *http.Request) {
	var params dto.PasswordResetInput
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}

	args := model.GetTokenParams{
		Token:     params.Token,
		ExpiresAt: int32(time.Now().Unix()),
	}

	_, err := h.query.GetToken(r.Context(), args)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "token expired request a new one")
		return
	}

	deleteParams := model.DeleteTokenParams{
		Token:  params.Token,
		UserID: params.UserID,
	}

	if err := h.query.DeleteToken(r.Context(), deleteParams); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "token expired request a new one")
		return
	}

	hashedPass, err := argon.CreateHash(params.Password, argon.DefaultParams)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to update the new password")
		return

	}

	input := model.UpdateUserPasswordByIDParams{
		ID:       params.UserID,
		Password: pgtype.Text{String: hashedPass, Valid: true},
	}

	if err := h.query.UpdateUserPasswordByID(r.Context(), input); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to update the new password")
		return

	}
	utils.SendResponse(w, "success", http.StatusOK, "successfully updated the new password")

	return
}
