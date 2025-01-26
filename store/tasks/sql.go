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

func (ts *taskStore) CreateTask(task models.Task) (int,error) {
	var (
		id int
	)
	err := ts.DB.QueryRow(`INSERT INTO tasks (user_id, title, description, is_completed, created_at)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		task.UserId, task.Title, task.Description, task.IsCompleted, task.CreatedAt).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ts *taskStore) GetUserTasks(userId, page, size int) ([]models.Task, error) {
	var tasks []models.Task

	offset := page * size
	rows, err := ts.DB.Query("SELECT * FROM tasks WHERE user_id=$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3", userId, size, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed to avoid memory leaks

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.Id, &task.UserId, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	// Check if there was any error during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ts *taskStore) GetTaskById(id int) (*models.Task, error) {
	var (
		task models.Task
	)

	if err := ts.DB.QueryRow("SELECT * FROM tasks WHERE id=$1", id).
		Scan(&task.Id, &task.UserId, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt); err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &task, nil
}

func (ts *taskStore) UpdateTaskById(task models.Task, id int) (*models.Task, error) {
	_, err := ts.DB.Exec("UPDATE tasks SET title=$1, description=$2,  is_completed=$3 WHERE id=$4",
		task.Title, task.Description, task.IsCompleted, id)
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
	_, err := ts.DB.Exec("UPDATE tasks SET is_completed = true WHERE id = $1", taskId)
	if err != nil {
		return err
	}
	return nil
}

func (ts *taskStore) GetUserCompletedTasks(isCompleted bool, userId, page, size int) ([]models.Task, error) {
	var tasks []models.Task

	offset := page * size
	rows, err := ts.DB.Query("SELECT * FROM tasks WHERE user_id=$1 AND is_completed=$2 ORDER BY created_at DESC LIMIT $3 OFFSET $4", userId, isCompleted,size, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed to avoid memory leaks

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.Id, &task.UserId, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	// Check if there was any error during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}


