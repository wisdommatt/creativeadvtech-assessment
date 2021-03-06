package users

import "time"

type User struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	FirstName   string    `json:"firstName" bson:"firstName,omitempty"`
	LastName    string    `json:"lastName" bson:"lastName,omitempty"`
	Email       string    `json:"email" bson:"email,omitempty"`
	Password    string    `json:"-" bson:"password,omitempty"`
	TimeAdded   time.Time `json:"timeAdded" bson:"timeAdded,omitempty"`
	LastUpdated time.Time `json:"-" bson:"lastUpdated,omitempty"`
}
