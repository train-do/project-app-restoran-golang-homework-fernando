package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/train-do/project-app-restoran-golang-homework-fernando/database"
	"github.com/train-do/project-app-restoran-golang-homework-fernando/service"
	"github.com/train-do/project-app-restoran-golang-homework-fernando/utils"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	service.RunningApp(db)
}
func init() {
	utils.ClearScreen()
	// x := model.User{
	// 	Customer: model.Customer{
	// 		Name: "Boojang",
	// 	},
	// }
	// service.InsertOrder(&sql.DB{}, x)
	// var tes []struct {
	// 	name  string
	// 	email string
	// }
	// uniqueItems := make([]struct{}, 0)
	// fmt.Println(reflect.TypeOf(tes).Name(), "<<<<")
	// fmt.Println(reflect.TypeOf(uniqueItems).Name(), "<<<<")
}
