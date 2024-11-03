package model

type Login struct {
	Username string
	Password string
}
type FormOrder struct {
	CustomerId int
	OrderId    int
	Discount   string
	Rating     int
	OrderItem  []struct {
		ItemId int
		Qty    int
	}
}
type FormOrderItem struct {
	OrderItemId int
	Status      string
}
type Body struct {
	Endpoint      string
	User          Login
	FormOrder     FormOrder
	FormOrderItem FormOrderItem
}
