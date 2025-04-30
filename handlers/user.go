package handlers

import (
	"net/http"

	"encoding/json"
	"log"

	"github.com/google/uuid"

	model "github.com/kaleabAlemayehu/drivee-server/model"
	"github.com/kaleabAlemayehu/drivee-server/utils"
)

func (h *handler) HandleGetAllUser(w http.ResponseWriter, r *http.Request) {
	owners, err := h.query.ListUser(r.Context())
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to fetch user list")
		return
	}
	if err := utils.SendResponse(w, "success", http.StatusOK, owners); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send user list")
		return
	}
}

func (h *handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println("unable to parse uuid that sent")
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	owner, err := h.query.GetUser(r.Context(), id)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to fetch owner")
		return
	}
	if err := utils.SendResponse(w, "success", http.StatusOK, owner); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send owner")
		return
	}
}

func (h *handler) HandleInsertUser(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.Context().Value("userID").(string))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	var params model.UpdateUserParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	params.ID = id
	res, err := h.query.UpdateUser(r.Context(), params)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request body")
		return
	}
	if err := utils.SendResponse(w, "success", http.StatusOK, res); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "internal server error")
		return
	}
}
