package repository

import (
	"database/sql"
	"fmt"

	"github.com/train-do/project-app-restoran-golang-homework-fernando/model"
	"github.com/train-do/project-app-restoran-golang-homework-fernando/view"
)

type RepoOrder struct {
	Order model.Order
}

func (r *RepoOrder) Create(tx *sql.Tx, customerId int) error {
	query := `insert into "Order" (customer_id, total_price, discount, rating, created_at) values ($1, 0, 0, 0, NOW()) returning id;`
	err := tx.QueryRow(query, customerId).Scan(&r.Order.Id)
	if err != nil {
		fmt.Println("Query Create Rows Order:", err)
		view.BadRequest("Bad Request : " + err.Error())
		tx.Rollback()
		return err
	}
	return nil
}
func (r *RepoOrder) Update(tx *sql.Tx, totalPrice int, promo string, rating int) error {
	var discount int
	switch promo {
	case "disc10":
		discount = 10000
	case "disc20":
		discount = 20000
	case "disc30":
		discount = 30000
	default:
		discount = 0
	}
	query := `update "Order" set total_price=$1, discount=$2, rating=$3 where id=$4 returning id;`
	err := tx.QueryRow(query, (totalPrice - discount), discount, rating, r.Order.Id).Scan(&r.Order.Id)
	if err != nil {
		fmt.Println("Query Update Rows Order:", err)
		view.BadRequest("Bad Request : " + err.Error())
		tx.Rollback()
		return err
	}
	return nil
}
func (r *RepoOrder) FindById(tx *sql.Tx) error {
	query := `select o.id, o.customer_id, o.total_price, o.discount, o.rating, o.created_at, oi.id, oi.quantity, oi.status, oi.created_at, i.id, i."name", i.price, i.created_at
	from "Order" o
	join "OrderItem" oi on oi.order_id = o.id
	join "Item" i on oi.item_id = i.id
	where o.id = $1;`
	rows, err := tx.Query(query, r.Order.Id)
	if err != nil {
		fmt.Println("Query Find Rows Order:", err)
		tx.Rollback()
		return err
	}
	for rows.Next() {
		var orderItem model.OrderItem
		err := rows.Scan(&r.Order.Id, &r.Order.CustomerId, &r.Order.TotalPrice, &r.Order.Discount, &r.Order.Rating, &r.Order.CreatedAt, &orderItem.Id, &orderItem.Qty, &orderItem.Status, &orderItem.CreatedAt, &orderItem.Item.Id, &orderItem.Item.Name, &orderItem.Item.Price, &orderItem.Item.CreatedAt)
		if err != nil {
			fmt.Println("Query Find Next Order:", err)
			tx.Rollback()
			return err
		}
		r.Order.OrderItem = append(r.Order.OrderItem, orderItem)
	}
	return nil
}
func (r *RepoOrder) FindAll(db *sql.DB) ([]model.Order, error) {
	var totalOrder []int
	query := `select o.id from "Order" o order by o.id;`
	var orders []model.Order
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Query FindAll Rows Order:", err)
		return []model.Order{}, err
	}
	for rows.Next() {
		var temp int
		err := rows.Scan(&temp)
		if err != nil {
			fmt.Println("Query FindAll Next Order:", err)
			return []model.Order{}, err
		}
		totalOrder = append(totalOrder, temp)
	}
	for _, v := range totalOrder {
		r.Order = model.Order{}
		query := `select o.id, o.customer_id, o.total_price, o.discount, o.rating, o.created_at, oi.id, oi.quantity, oi.status, oi.created_at, i.id, i."name", i.price, i.created_at
		from "Order" o
		join "OrderItem" oi on oi.order_id = o.id
		join "Item" i on oi.item_id = i.id
		where o.id = $1 order by o.id;`
		rows, err := db.Query(query, v)
		if err != nil {
			fmt.Println("Query Find Rows Order:", err)
			return []model.Order{}, err
		}
		for rows.Next() {
			var orderItem model.OrderItem
			err := rows.Scan(&r.Order.Id, &r.Order.CustomerId, &r.Order.TotalPrice, &r.Order.Discount, &r.Order.Rating, &r.Order.CreatedAt, &orderItem.Id, &orderItem.Qty, &orderItem.Status, &orderItem.CreatedAt, &orderItem.Item.Id, &orderItem.Item.Name, &orderItem.Item.Price, &orderItem.Item.CreatedAt)
			if err != nil {
				fmt.Println("Query Find Next Order:", err)
				return []model.Order{}, err
			}
			r.Order.OrderItem = append(r.Order.OrderItem, orderItem)
		}
		orders = append(orders, r.Order)
	}
	return orders, nil
}
func (r *RepoOrder) FindByCustomerId(db *sql.DB, customerId int) ([]model.Order, error) {
	var totalOrder []int
	query := `select o.id from "Order" o where o.customer_id=$1 order by o.id;`
	var orders []model.Order
	rows, err := db.Query(query, customerId)
	if err != nil {
		fmt.Println("Query FindAll Rows Order:", err)
		return []model.Order{}, err
	}
	for rows.Next() {
		var temp int
		err := rows.Scan(&temp)
		if err != nil {
			fmt.Println("Query FindAll Next Order:", err)
			return []model.Order{}, err
		}
		totalOrder = append(totalOrder, temp)
	}
	for _, v := range totalOrder {
		r.Order = model.Order{}
		query := `select o.id, o.customer_id, o.total_price, o.discount, o.rating, o.created_at, oi.id, oi.quantity, oi.status, oi.created_at, i.id, i."name", i.price, i.created_at
		from "Order" o
		join "OrderItem" oi on oi.order_id = o.id
		join "Item" i on oi.item_id = i.id
		where o.id = $1;`
		rows, err := db.Query(query, v)
		if err != nil {
			fmt.Println("Query Find Rows Order:", err)
			return []model.Order{}, err
		}
		for rows.Next() {
			var orderItem model.OrderItem
			err := rows.Scan(&r.Order.Id, &r.Order.CustomerId, &r.Order.TotalPrice, &r.Order.Discount, &r.Order.Rating, &r.Order.CreatedAt, &orderItem.Id, &orderItem.Qty, &orderItem.Status, &orderItem.CreatedAt, &orderItem.Item.Id, &orderItem.Item.Name, &orderItem.Item.Price, &orderItem.Item.CreatedAt)
			if err != nil {
				fmt.Println("Query Find Next Order:", err)
				return []model.Order{}, err
			}
			r.Order.OrderItem = append(r.Order.OrderItem, orderItem)
		}
		orders = append(orders, r.Order)
	}
	return orders, nil
}
