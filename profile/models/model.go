package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Profile struct {
	Id         bson.ObjectId     `bson:"_id,omitempty" json:"id"`
	Name       string            `bson:"name" json:"name"`
	Department string            `bson:"department" json:"department"`
	Keys       map[string]string `bson:"keys" json:"keys"`
	//CreatedOn  time.Time     `json:"createdon,omitempty"`
}
