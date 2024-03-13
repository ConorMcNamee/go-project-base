package database

import (
	"context"
	"gobaseproject/config"
	"gobaseproject/models"
	"time"
)

var app *config.AppConfig

func SetDB(a *config.AppConfig) {
	app = a
}

func GetUserByID(id int64) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `select id, first_name, last_name, email, password, created_at, updated_at from users where id = $1`

	var u models.User
	row := app.DB.SQL.QueryRowContext(ctx, query, id)

	if err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.Admin, &u.Age); err != nil {
		return u, err
	}

	return u, nil
}

func CreateNewUser() {}
