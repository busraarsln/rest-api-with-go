package models

import (
	"encoding/json"
	"io"
	"time"
)

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

func (u *User) FromJSON(data io.Reader) error {
	de := json.NewDecoder(data)
	return de.Decode(u)
}

// ToJSON converts the collection to json
func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}
