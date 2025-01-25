package models

import "time"

type Task struct {
	Id          int
	UserId      int
	Title       string
	Description string
	IsCompleted string
	CreatedAt   time.Time
}
