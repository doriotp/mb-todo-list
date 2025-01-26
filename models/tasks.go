package models

import "time"

type Task struct {
    Id          int       `json:"id,omitempty"`
    UserId      int       `json:"user_id,omitempty"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    IsCompleted string    `json:"is_completed"`
    CreatedAt   time.Time `json:"created_at"`
}

type CreateTaskResponse struct{
	Id          int       `json:"id"`
    UserId      int       `json:"user_id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    IsCompleted string    `json:"is_completed"`
    CreatedAt   time.Time `json:"created_at"`
}
