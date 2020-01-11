package models

// Profile is object containing all feils for profile in mongodb
type Profile struct {
	//ID         bson.ObjectId     `bson:"_id,omitempty" json:"id"`
	Name       string            `bson:"name" json:"name"`
	Department string            `bson:"department" json:"department"`
	Keys       map[string]string `bson:"keys" json:"keys"`
	//CreatedOn  time.Time     `json:"createdon,omitempty"`
	//CreatedBy Edit list?
}
