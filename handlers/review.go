package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/kaleabAlemayehu/drivee-server/model"
	"github.com/kaleabAlemayehu/drivee-server/utils"
)

func (h *handler) HandleGetAllReviews(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	reviews, err := h.query.ListReviews(h.ctx, id)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	if reviews == nil {
		log.Println("there is no review to list")
		utils.SendResponse(w, "success", http.StatusNoContent, "No review found")
		return
	}
	if err := utils.SendResponse(w, "success", http.StatusOK, reviews); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send review")
		return
	}
}

func (h *handler) HandleGetReview(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	review, err := h.query.GetReview(h.ctx, id)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	if err := utils.SendResponse(w, "success", http.StatusOK, review); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send review")
		return
	}
}

func (h *handler) HandleInsertReview(w http.ResponseWriter, r *http.Request) {
	reviewerID, err := uuid.Parse(r.Context().Value("userID").(string))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusForbidden, "unauthorized write")
		return
	}
	var body model.InsertReviewParams
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	body.ReviewerID = reviewerID

	review, err := h.query.InsertReview(h.ctx, body)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}

	if err := utils.SendResponse(w, "success", http.StatusCreated, review); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusInternalServerError, "unable to send created review")
		return
	}
}

func (h *handler) HandleUpdateReview(w http.ResponseWriter, r *http.Request) {
	reviewerID, err := uuid.Parse(r.Context().Value("userID").(string))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusForbidden, "unauthorized write")
		return
	}
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	var body model.UpdateReviewParams
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	body.ID = id
	body.ReviewerID = reviewerID
	review, err := h.query.UpdateReview(h.ctx, body)
	if err != nil {
		log.Println(err.Error())
		utils.SendResponse(w, "error", http.StatusBadRequest, "bad request")
		return
	}
	if err := json.NewEncoder(w).Encode(review); err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
