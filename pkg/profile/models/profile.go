package models

import (
	v "github.com/gobuffalo/validate"
)

// Profile is object containing all feils for profile in mongodb
type Profile struct {
	//ID         bson.ObjectId     `bson:"_id,omitempty" json:"id"`
	Name       string            `bson:"name" json:"name"`
	Department string            `bson:"department" json:"department"`
	Keys       map[string]string `bson:"keys" json:"keys"`
}

// IsValid is called by gobuffalo to check if models is fully populated
func (p *Profile) IsValid(errors *v.Errors) {
	if p.Name == "" {
		errors.Add("name", "Name must not be blank!")
	}
	if p.Department == "" {
		errors.Add("department", "department must not be blank!")
	}
	if len(p.Keys) == 0 {
		errors.Add("key", "Keys must not be blank!")
	}
}
