package profile

import (
	"log"

	"github.com/CBSktravers/hooli/pkg/profile/models"
	v "github.com/gobuffalo/validate"
)

// Service manages profile
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
	// validate all infomation is here
	log.Println("Service create called")

	//Validate that all fields that are required are there
	errors := v.Validate(&profile)
	if len(errors.Errors) != 0 {
		log.Println(errors.Errors)
		return errors
	}
	err := s.repo.Create(&profile)
	return err
}
