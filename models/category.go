package models

import (
	"time"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	//ProductIDs []string  `json: "-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
