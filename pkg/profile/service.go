package profile

import (
	"log"

	"github.com/CBSktravers/hooli/pkg/profile/models"
)

// Service manages cloud credentials
type Service interface {
	Create(profile models.Profile) error
}

// DefaultService is an implementation of Service
type DefaultService struct {
	repo Repository
}

// NewDefaultService creates and returns a pointer to a DefaultService
func NewDefaultService(repo Repository) *DefaultService {
	return &DefaultService{repo: repo}
}

// Create makes a new Profile into database service
func (s *DefaultService) Create(profile models.Profile) error {
	// All logic to create profile and call repo to perform task
	log.Println("Service create called")
	s.repo.Create(&profile)
	return nil
}
