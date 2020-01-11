package profile

import (
	"github.com/CBSktravers/hooli/pkg/profile/models"
)

// Repository manages am underlying storage mechanism for Profile
type Repository interface {
	Create(profile *models.Profile) error
	/*Get(name string, department string) (models.Profile, error)
	List(department string) (models.Profile, error)
	Delete(name string, department string) error
	Update(profile *models.Profile) error*/
}
