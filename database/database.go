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

func GetUserByEmail(email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `select id, first_name, last_name, email, password, created_at, updated_at from users where email = $1`

	var u models.User
	row := app.DB.SQL.QueryRowContext(ctx, query, email)

	if err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.Admin, &u.Age); err != nil {
		return u, err
	}

	return u, nil
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

func CreateNewUser(u models.User) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into users (first_name, last_name, email, password, age, admin, joined_at, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := app.DB.SQL.ExecContext(ctx, query, u.FirstName, u.LastName, u.Email, u.Password, u.Age, u.Admin, u.Joined_at, u.Created_at, u.Updated_at)

	if err != nil {
		panic(err)
	}

	// MUST RETURN USER ID
	return 0, nil
}
