package service

import (
	"database/sql"
	"fmt"

	"github.com/train-do/project-app-restoran-golang-homework-fernando/model"
	"github.com/train-do/project-app-restoran-golang-homework-fernando/repository"
)

func Login(db *sql.DB) interface{} {
	repoUser := repository.RepoUser{
		User: model.User{
			Username: Request.User.Username,
			Password: Request.User.Password,
		},
	}
	user, err := repoUser.FindByUsernamePassword(db)
	if err != nil {
		fmt.Println(err)
	}
	return user
}
