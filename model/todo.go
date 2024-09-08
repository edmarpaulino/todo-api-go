package model

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"createdAt"`
}
