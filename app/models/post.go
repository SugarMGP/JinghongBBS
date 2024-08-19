package models

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	User      uint      `json:"user_id"`
	CreatedAt time.Time `json:"time"`
}
