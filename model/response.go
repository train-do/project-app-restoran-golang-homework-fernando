package model

type ResponseSuccess struct {
	StatusCode int
	Message    string
	Data       interface{}
}
type ResponseError struct {
	StatusCode int
	Message    string
}
