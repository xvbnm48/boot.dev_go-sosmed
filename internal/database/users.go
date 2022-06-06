package database

import "time"

type User struct {
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}
