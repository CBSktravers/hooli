package handlers

import (
	"github.com/CBSktravers/hooli/profile/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProfileCollection struct {
	C *mgo.Collection
}

func (r *ProfileCollection) Create(profile *models.Profile) error {
	obj_id := bson.NewObjectId()
	profile.Id = obj_id
	//profile.CreatedOn = time.Now()
	err := r.C.Insert(&profile)
	return err
}

func (r *ProfileCollection) GetAll() []models.Profile {
	var profiles []models.Profile
	iter := r.C.Find(nil).Iter()
	result := models.Profile{}
	for iter.Next(&result) {
		profiles = append(profiles, result)
	}
	return profiles
}

/*
func (r *ProfileCollection) GetByDate(date string) (profile models.Profile, err error) {
	err = r.C.Find(bson.M{"date": date}).One(&profile)
	return
}
*/

func (r *ProfileCollection) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
