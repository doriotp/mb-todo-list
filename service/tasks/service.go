package tasks

import (
	"net/http"
	"time"

	customerrors "github.com/todo-list/customErrors"
	"github.com/todo-list/models"
)

type taskService struct {
	tskStore taskStore
	// usrStore userStore
}

func New(tskStore taskStore) *taskService {
	return &taskService{tskStore: tskStore}
}

func (ts *taskService) CreateTask(task models.Task) *customerrors.Error {
	if task.Title == "" || task.Description == "" {
		return &customerrors.Error{Code: http.StatusBadRequest, Message: "invalid input"}
	}

	// userInfo, err := ts.usrStore.GetUserById(task.UserId)
	// if err != nil {
	// 	return &customerrors.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	// }

	// if userInfo == nil {
	// 	return &customerrors.Error{Code: http.StatusBadRequest, Message:"user does not exist"}
	// }

	createdAt := time.Now() 

	task.CreatedAt=createdAt 

	err := ts.tskStore.CreateTask(task) 
	if err!=nil{
		return &customerrors.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return nil

}

func (ts *taskService)GetUserTasks(userId, page, size int) (*models.Task, *customerrors.Error){

	tasks, err := ts.tskStore.GetUserTasks(userId,page, size )
	if err!=nil{
		return nil, &customerrors.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return tasks, nil
}

func (ts *taskService)GetTaskById(id int) (*models.Task, *customerrors.Error){
	tasks, err := ts.tskStore.GetTaskById(id)
	if err!=nil{
		return nil, &customerrors.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return tasks, nil
}

func (ts *taskService)DeleteTaskById(id int) *customerrors.Error{
	err := ts.tskStore.DeleteTaskById(id)
	if err!=nil{
		return &customerrors.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return nil
}

func (ts *taskService)UpdateTaskCompletionStatus(taskId int) (*models.Task, *customerrors.Error){

	err := ts.tskStore.UpdateTaskCompletionStatus(taskId) 
	if err!=nil{
		return nil, &customerrors.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	tasks, err := ts.tskStore.GetTaskById(taskId)
	if err!=nil{
		return nil, &customerrors.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return tasks, nil

}

func(ts *taskService)GetUserCompletedTasks(isCompleted bool, userId, page, size int) (*models.Task, *customerrors.Error){
	CompletedTasks , err := ts.tskStore.GetUserCompletedTasks(isCompleted, userId, page, size)
	if err!=nil{
		return nil, &customerrors.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return CompletedTasks, nil
}

func(ts *taskService)UpdateTaskById(task models.Task, id int) (*models.Task, *customerrors.Error){
	tasks, err := ts.tskStore.UpdateTaskById(task, id)
	if err!=nil{
		return nil, &customerrors.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return tasks, nil
}

















