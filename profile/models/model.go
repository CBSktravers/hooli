package models

// Profile is object for profiles in mongodatabase
// Consiting of name of the profile and Department
// Last feild is keys for mapping
type Profile struct {
	//ID         bson.ObjectId     `bson:"_id,omitempty" json:"id"`
	Name       string            `bson:"name" json:"name"`
	Department string            `bson:"department" json:"department"`
	Keys       map[string]string `bson:"keys" json:"keys"`
	//CreatedOn  time.Time     `json:"createdon,omitempty"`
}
