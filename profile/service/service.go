package service

import (
	"github.com/CBSktravers/hooli/profile/models"
)

// Service is an implementation of Service
type Service struct {
	repo Repository
}

// NewDefaultService creates and returns a pointer to a DefaultService
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(profile models.Profile) error {

	return nil
}
