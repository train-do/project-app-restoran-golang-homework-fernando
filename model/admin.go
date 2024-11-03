package model

import "time"

type Admin struct {
	Id        int       `json:"id,omitempty"`
	UserId    int       `json:"userId,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
