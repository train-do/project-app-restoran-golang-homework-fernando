package model

import "time"

type OrderItem struct {
	Id        int       `json:"id"`
	OrderId   int       `json:"orderId,omitempty"`
	ItemId    int       `json:"itemId,omitempty"`
	Qty       int       `json:"qty"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	Item      Item      `json:"item"`
}
