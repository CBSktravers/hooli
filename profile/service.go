package profile

import (
	"github.com/CBSktravers/hooli/profile/models"
)

// Service is an implementation of Service
type Service struct {
	repo Repository
}

func (s *Service) Create(profile models.Profile) error {

	return nil
}
