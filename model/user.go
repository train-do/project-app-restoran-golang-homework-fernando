package model

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	Admin     Admin     `json:"admin,omitempty"`
	Chef      Chef      `json:"chef,omitempty"`
	Customer  Customer  `json:"customer,omitempty"`
}
