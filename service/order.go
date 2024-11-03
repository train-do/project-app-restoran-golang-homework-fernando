package service

import (
	"database/sql"
	"fmt"

	"github.com/train-do/project-app-restoran-golang-homework-fernando/model"
	"github.com/train-do/project-app-restoran-golang-homework-fernando/repository"
	"github.com/train-do/project-app-restoran-golang-homework-fernando/utils"
	"github.com/train-do/project-app-restoran-golang-homework-fernando/view"
)

func insertOrder(db *sql.DB, user interface{}) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error Tx Begin : ", err)
	}
	err = utils.ValidationFormOrder(Request)
	if err != nil {
		return
	}
	var customerId int
	switch user.(type) {
	case model.Customer:
		user, _ := user.(model.Customer)
		customerId = user.Id
	case model.Admin:
		if Request.FormOrder.CustomerId == 0 {
			view.BadRequest("Field Customer Id is Empty")
			return
		}
		customerId = Request.FormOrder.CustomerId
	case model.Chef:
		view.NotAuthorized()
		return
	}
	repoOrder := repository.RepoOrder{}
	err = repoOrder.Create(tx, customerId)
	if err != nil {
		return
	}
	RepoOrderItem := repository.RepoOrderItem{}
	var totalPrice int
	for _, v := range Request.FormOrder.OrderItem {
		// fmt.Printf("%+v\n", RepoOrderItem.OrderItem)
		// fmt.Println(repoOrder.Order.Id, v.ItemId, v.Qty)
		err = RepoOrderItem.Create(tx, repoOrder.Order.Id, v.ItemId, v.Qty)
		if err != nil {
			return
		}
		totalPrice += (RepoOrderItem.OrderItem.Item.Price * v.Qty)
	}
	err = repoOrder.Update(tx, totalPrice, Request.FormOrder.Discount, Request.FormOrder.Rating)
	if err != nil {
		return
	}
	// fmt.Println(totalPrice, repoOrder.Order.Id)
	err = repoOrder.FindById(tx)
	if err != nil {
		return
	}
	tx.Commit()
	view.Success("Order Success Added", repoOrder.Order)
}

func getOrder(db *sql.DB, user interface{}) {
	repoOrder := repository.RepoOrder{}
	var response []model.Order
	switch user.(type) {
	case model.Customer:
		user, _ := user.(model.Customer)
		response, _ = repoOrder.FindByCustomerId(db, user.Id)
	case model.Admin:
		response, _ = repoOrder.FindAll(db)
	case model.Chef:
		response, _ = repoOrder.FindAll(db)
	}
	view.Success("Success Retrieve Data", response)
}
