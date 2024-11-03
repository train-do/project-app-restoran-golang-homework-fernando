package service

import (
	"database/sql"

	"github.com/train-do/project-app-restoran-golang-homework-fernando/model"
	"github.com/train-do/project-app-restoran-golang-homework-fernando/utils"
	"github.com/train-do/project-app-restoran-golang-homework-fernando/view"
)

var Request model.Body

func init() {
	utils.DecodeFromJSON(&Request)
}
func RunningApp(db *sql.DB) {
	user := Login(db)
	switch Request.Endpoint {
	case "addOrder":
		insertOrder(db, user)
	case "getOrder":
		getOrder(db, user)
	case "updateOrderItem":
		updateStatusOrderItem(db, user)
	case "deleteOrderItem":
		deleteOrderItem(db, user)
	default:
		view.NotFound()
	}
}
