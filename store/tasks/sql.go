package tasks

import (
	"database/sql"

	"github.com/todo-list/models"
)

type taskStore struct {
	DB *sql.DB
}

func New(db *sql.DB) *taskStore {
	return &taskStore{DB: db}
}

func (ts *taskStore) CreateTask(task models.Task) error {
	_, err := ts.DB.Exec(`INSERT INTO tasks (id,userId,title,description,
	isCompleted,createdAt) VALUES ($1,$2,$3,$4,$5,$6)`, task.Id, task.UserId, task.Title, task.Description,
		task.IsCompleted, task.CreatedAt)
	return err
}

func (ts *taskStore) GetUserTasks(userId, page, size int) (*models.Task, error) {
	var (
		task models.Task
	)

	offset := page * size
	if err := ts.DB.QueryRow("SELECT * FROM tasks WHERE userId=$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3", userId, size, offset).
		Scan(&task.Id, &task.UserId, &task.Title, &task.Description, &task.IsCompleted); err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &task, nil
}

func (ts *taskStore) GetTaskById(id int) (*models.Task, error) {
	var (
		task models.Task
	)

	if err := ts.DB.QueryRow("SELECT * FROM tasks WHERE id=$1", id).
		Scan(&task.Id, &task.UserId, &task.Title, &task.Description, &task.IsCompleted); err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &task, nil
}

func (ts *taskStore) UpdateTaskById(task models.Task, id int) (*models.Task, error) {
	_, err := ts.DB.Exec("UPDATE SET title=$2, description=$3, isCompleted=$4 WHERE id=$5",
		task.Title, task.Description, task.IsCompleted)
	if err != nil {
		return nil, err
	}

	return &task, err
}

func (ts *taskStore) DeleteTaskById(id int) error {
	_, err := ts.DB.Exec("DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

func (ts *taskStore) UpdateTaskCompletionStatus(taskId int) error {
	// Update the task's 'isCompleted' status in the database
	_, err := ts.DB.Exec("UPDATE tasks SET isCompleted = true WHERE id = $1", taskId)
	if err != nil {
		return err
	}
	return nil
}

func (ts *taskStore) GetUserCompletedTasks(isCompleted bool, userId, page, size int) (*models.Task, error) {
	var (
		task models.Task
	)

	offset := page*size

	if err := ts.DB.QueryRow("SELECT * FROM tasks WHERE userId=$1 AND isCompleted=$2 ORDER BY created_at DESC LIMIT $3 OFFSET $4", userId, isCompleted,size, offset).
		Scan(&task.Id, &task.UserId, &task.Title, &task.Description, &task.IsCompleted); err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &task, nil
}


