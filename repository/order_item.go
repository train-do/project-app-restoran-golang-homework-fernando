package repository

import (
	"database/sql"
	"fmt"

	"github.com/train-do/project-app-restoran-golang-homework-fernando/model"
	"github.com/train-do/project-app-restoran-golang-homework-fernando/view"
)

type RepoOrderItem struct {
	OrderItem model.OrderItem
}

func (r *RepoOrderItem) Create(tx *sql.Tx, orderId int, itemId int, qty int) error {
	query := `select id, price from "Item" where id = $1;`
	err := tx.QueryRow(query, itemId).Scan(&r.OrderItem.Item.Id, &r.OrderItem.Item.Price)
	if err != nil {
		fmt.Println("Query Create Rows OrderItem/Item:", err)
		view.BadRequest("Invalid Item Order")
		tx.Rollback()
		return err
	}
	query = `insert into "OrderItem" (order_id, item_id, quantity, status, created_at) values ($1, $2, $3, 'ordered', NOW()) returning id;`
	err = tx.QueryRow(query, orderId, itemId, qty).Scan(&r.OrderItem.Id)
	if err != nil {
		fmt.Println("Query Create Rows OrderItem:", err)
		view.BadRequest("Bad Request : " + err.Error())
		tx.Rollback()
		return err
	}
	return nil
}
func (r *RepoOrderItem) Update(db *sql.DB) error {
	// fmt.Println(r.OrderItem.Status, r.OrderItem.Id)
	query := `update "OrderItem" set status=$1 where id=$2 returning id;`
	err := db.QueryRow(query, r.OrderItem.Status, r.OrderItem.Id).Scan(&r.OrderItem.Id)
	if err != nil {
		fmt.Println("Query Update Rows OrderItem:", err)
		view.BadRequest("Invalid Item Order")
		return err
	}
	return nil
}
func (r *RepoOrderItem) Delete(db *sql.DB) error {
	// fmt.Println(r.OrderItem.Status, r.OrderItem.Id)
	query := `update "OrderItem" set status=$1 where id=$2 returning id;`
	err := db.QueryRow(query, r.OrderItem.Status, r.OrderItem.Id).Scan(&r.OrderItem.Id, &r.OrderItem.OrderId, &r.OrderItem.ItemId, &r.OrderItem.Qty, &r.OrderItem.Status, &r.OrderItem.CreatedAt)
	if err != nil {
		fmt.Println("Query Update Rows OrderItem:", err)
		view.BadRequest("Invalid Item Order")
		return err
	}
	return nil
}
func (r *RepoOrderItem) FindById(db *sql.DB) {
	query := `select o.id, o.order_id, o.quantity, o.status, o.created_at, i.id, i.name, i.price, i.created_at from "OrderItem" o join "Item" i on o.item_id = i.id where o.id=$1;`
	err := db.QueryRow(query, r.OrderItem.Id).Scan(&r.OrderItem.Id, &r.OrderItem.OrderId, &r.OrderItem.Qty, &r.OrderItem.Status, &r.OrderItem.CreatedAt, &r.OrderItem.Item.Id, &r.OrderItem.Item.Name, &r.OrderItem.Item.Price, &r.OrderItem.Item.CreatedAt)
	if err != nil {
		fmt.Println("Query FindById Rows OrderItem:", err)
		view.BadRequest("Invalid Item Order")
		return
	}
}
