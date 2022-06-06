package database

import "time"

type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserEmail string    `json:"useremail"`
	Text      string    `json:"text"`
}
