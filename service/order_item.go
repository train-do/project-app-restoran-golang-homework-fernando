package service

import (
	"database/sql"
	"strings"

	"github.com/train-do/project-app-restoran-golang-homework-fernando/model"
	"github.com/train-do/project-app-restoran-golang-homework-fernando/repository"
	"github.com/train-do/project-app-restoran-golang-homework-fernando/utils"
	"github.com/train-do/project-app-restoran-golang-homework-fernando/view"
)

func updateStatusOrderItem(db *sql.DB, user interface{}) {
	err := utils.ValidationFormOrderItem(Request)
	if err != nil {
		return
	}
	if strings.ToLower(Request.FormOrderItem.Status) == "cancel" {
		// Karena endpoint update tapi status soft delete
		view.InternalServerError()
		return
	}
	switch user.(type) {
	case model.Customer:
		view.NotAuthorized()
		return
	}
	repoOrderItem := repository.RepoOrderItem{
		OrderItem: model.OrderItem{
			Id:     Request.FormOrderItem.OrderItemId,
			Status: Request.FormOrderItem.Status,
		},
	}
	repoOrderItem.Update(db)
	repoOrderItem.FindById(db)
	view.Success("Update Status Succes", repoOrderItem.OrderItem)
}
func deleteOrderItem(db *sql.DB, user interface{}) {
	switch user.(type) {
	case model.Customer:
		view.NotAuthorized()
		return
	case model.Chef:
		view.NotAuthorized()
		return
	}
	if strings.ToLower(Request.FormOrderItem.Status) != "cancel" {
		// Karena endpoint delete tapi value status selain cancel
		view.InternalServerError()
		return
	}
	repoOrderItem := repository.RepoOrderItem{
		OrderItem: model.OrderItem{
			Id:     Request.FormOrderItem.OrderItemId,
			Status: Request.FormOrderItem.Status,
		},
	}
	repoOrderItem.Update(db)
	repoOrderItem.FindById(db)
	view.Success("Delete Status Succes", repoOrderItem.OrderItem)
}
