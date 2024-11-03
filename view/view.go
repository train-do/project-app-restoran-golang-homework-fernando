package view

import (
	"encoding/json"
	"fmt"

	"github.com/train-do/project-app-restoran-golang-homework-fernando/model"
)

func Success(msg string, data interface{}) {
	response := model.ResponseSuccess{
		StatusCode: 200,
		Message:    msg,
		Data:       data,
	}
	jsonData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		fmt.Println("JSON Marshal Error")
	}
	fmt.Println(string(jsonData))
}
func NoAuthentication() {
	response := model.ResponseError{
		StatusCode: 401,
		Message:    "No Authentication",
	}
	jsonData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		fmt.Println("JSON Marshal Error")
	}
	fmt.Println(string(jsonData))
}
func NotAuthorized() {
	response := model.ResponseError{
		StatusCode: 401,
		Message:    "Not Authorized",
	}
	jsonData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		fmt.Println("JSON Marshal Error")
	}
	fmt.Println(string(jsonData))
}
func NotFound() {
	response := model.ResponseError{
		StatusCode: 404,
		Message:    "404 Not Found",
	}
	jsonData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		fmt.Println("JSON Marshal Error")
	}
	fmt.Println(string(jsonData))
}
func BadRequest(msg string) {
	response := model.ResponseError{
		StatusCode: 400,
		Message:    msg,
	}
	jsonData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		fmt.Println("JSON Marshal Error")
	}
	fmt.Println(string(jsonData))
}
func InternalServerError() {
	response := model.ResponseError{
		StatusCode: 500,
		Message:    "Internal Server Error",
	}
	jsonData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		fmt.Println("JSON Marshal Error")
	}
	fmt.Println(string(jsonData))
}
