package model

import "time"

type Order struct {
	Id         int         `json:"id"`
	CustomerId int         `json:"customerId,omitempty"`
	TotalPrice int         `json:"totalPrice"`
	Discount   int         `json:"discount"`
	Rating     int         `json:"rating"`
	CreatedAt  time.Time   `json:"createdAt"`
	OrderItem  []OrderItem `json:"orderItem"`
}
