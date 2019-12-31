package handlers

import (
	"github.com/CBSktravers/hooli/profile/models"
)

type (
	// For Get - /profiles
	ProfilesResource struct {
		Data []models.Profile `json:"data"`
	}
	// For Post/Put - /profiles
	ProfileResource struct {
		Data models.Profile `json:"data"`
	}
)
