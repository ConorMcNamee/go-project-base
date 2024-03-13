package models

import "time"

type User struct {
	ID         int
	FirstName  string
	LastName   string
	Email      string
	Password   string
	Age        int
	Admin      bool
	Joined_at  time.Time
	Created_at time.Time
	Updated_at time.Time
}
