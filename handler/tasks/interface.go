package tasks

import (
	customerrors "github.com/todo-list/customErrors"
	"github.com/todo-list/models"
)

type taskService interface {
	CreateTask(task models.Task) *customerrors.Error
	GetUserTasks(userId, page, size int) (*models.Task, *customerrors.Error)
	GetTaskById(id int) (*models.Task, *customerrors.Error)
	DeleteTaskById(id int) *customerrors.Error
	UpdateTaskCompletionStatus(taskId int)(*models.Task, *customerrors.Error)
	GetUserCompletedTasks(isCompleted bool, userId, page, size int) (*models.Task, *customerrors.Error)
	UpdateTaskById(task models.Task, id int) (*models.Task, *customerrors.Error)
}
