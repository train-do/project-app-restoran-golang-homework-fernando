package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strings"

	"github.com/train-do/project-app-restoran-golang-homework-fernando/model"
	"github.com/train-do/project-app-restoran-golang-homework-fernando/view"
)

func ValidationFormOrder(req model.Body) error {
	form := req.FormOrder.OrderItem
	if len(req.FormOrder.OrderItem) <= 0 || reflect.TypeOf(form).Name() != "" {
		view.BadRequest("Field is Empty")
		return errors.New("Invalid Form")
	}
	return nil
}
func ValidationFormOrderItem(req model.Body) error {
	form := req.FormOrderItem
	status := strings.ToLower(form.Status)
	if form.OrderItemId == 0 {
		view.BadRequest("Field Item Order is Empty")
		return errors.New("Invalid Form")
	}
	if status == "" {
		view.BadRequest("Field Status is Empty")
		return errors.New("Invalid Form")
	}
	if status != "cancel" && status != "completed" && status != "process" {
		view.BadRequest("Invalid Status Value")
		return errors.New("Invalid Form")
	}
	return nil
}
func DecodeFromJSON(req interface{}) {
	// fmt.Printf("%s\n%+v\n", reflect.TypeOf(req).Name(), req)
	file, err := os.Open("body.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	// fmt.Printf("%v\n%v\n", reflect.TypeOf(req).Name(), req)
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(req); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	// fmt.Printf("After Decode :%s\n%+v\n", reflect.TypeOf(req).Name(), req)
}
func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
