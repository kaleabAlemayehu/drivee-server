package handlers

import (
	"encoding/json"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kaleabAlemayehu/drivee-server/dto"
	"github.com/kaleabAlemayehu/drivee-server/model"
	"github.com/kaleabAlemayehu/drivee-server/utils"
	"log"
	"net/http"
	"os"
)

func (h *handler) HandleGoogleAuth(w http.ResponseWriter, r *http.Request) {
	var params dto.GoogleBody
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	// Initialize verifier (should be a global instance in real app)
	verifier := utils.NewGoogleTokenVerifier(os.Getenv("GOOGLE_CLIENT_ID"))

	// Verify token
	claims, err := verifier.VerifyToken(params.Token)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}

	user, err := h.query.GetUserByEmail(r.Context(), claims.Email)
	if err != nil {
		var lastName pgtype.Text
		if claims.LastName != "" {
			lastName = pgtype.Text{
				String: claims.LastName,
				Valid:  true,
			}
		} else {
			lastName = pgtype.Text{
				String: claims.LastName,
				Valid:  false,
			}

		}

		params := model.InsertUserSSOParams{
			FirstName:      claims.FirstName,
			Email:          claims.Email,
			ProfilePicture: claims.Picture,
			LastName:       lastName,
		}
		user, err := h.query.InsertUserSSO(r.Context(), params)
		if err != nil {
			log.Println(err.Error())
			utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to signup the user")
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
		return
	} else {
		if user.Password.String != "" || user.Password.Valid {
			log.Println("user already exist with password auth")
			utils.SendResponse(w, "error", http.StatusBadRequest, "email already exist use email and password instead.")
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
		return
	}
}
