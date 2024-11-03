package model

import "time"

type Item struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}
