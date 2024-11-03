package repository

import (
	"database/sql"
	"fmt"

	"github.com/train-do/project-app-restoran-golang-homework-fernando/model"
)

type RepoUser struct {
	User model.User
}

func (r *RepoUser) FindByUsernamePassword(db *sql.DB) (interface{}, error) {
	query := `select u.id from "User" u where u.username = $1 and u.password = $2; `
	// fmt.Printf("%+v *******\n", *r)
	err := db.QueryRow(query, r.User.Username, r.User.Password).Scan(&r.User.Id)
	if err != nil {
		fmt.Println("Error Query Rows :", err)
		return model.User{}, err
	}
	query = `select * from "Admin" where user_id = $1; `
	err = db.QueryRow(query, r.User.Id).Scan(&r.User.Admin.Id, &r.User.Admin.UserId, &r.User.Admin.Name, &r.User.Admin.CreatedAt)
	if err == nil {
		return r.User.Admin, err
	}
	query = `select * from "Chef" where user_id = $1; `
	err = db.QueryRow(query, r.User.Id).Scan(&r.User.Chef.Id, &r.User.Chef.UserId, &r.User.Chef.Name, &r.User.Chef.CreatedAt)
	if err == nil {
		return r.User.Chef, err
	}
	query = `select * from "Customer" where user_id = $1; `
	err = db.QueryRow(query, r.User.Id).Scan(&r.User.Customer.Id, &r.User.Customer.UserId, &r.User.Customer.Name, &r.User.Customer.CreatedAt)
	if err == nil {
		return r.User.Customer, err
	}
	return r.User, nil
}
