package dto

import (
	"github.com/kaleabAlemayehu/drivee-server/model"
)

type CarResponse struct {
	Cars  []model.ListCarsRow `json:"cars"`
	Total int64               `json:"total"`
}
