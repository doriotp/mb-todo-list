package tasks

import "github.com/todo-list/models"

type taskStore interface {
	CreateTask(task models.Task) (int,error)
	GetUserTasks(userId, page, size int) ([]models.Task, error) 
	GetTaskById(id int) (*models.Task, error)
	DeleteTaskById(id int) error
	UpdateTaskCompletionStatus(taskId int) error
	GetUserCompletedTasks(isCompleted bool, userId, page, size int) ([]models.Task, error) 
	UpdateTaskById(task models.Task, id int) (*models.Task, error)
}

// type userStore interface {
// 	GetUserById(id int) (*models.User, error)
// }
