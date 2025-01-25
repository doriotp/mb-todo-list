package users

import "github.com/todo-list/models"

type userStore interface {
	CreateUser(models.User) (error)
	GetUserByEmail(string) (*models.User, error)
	UpdatePasswordById(password string, id int) (error)
	GetUserById(id int)(*models.User, error)
	UpdateUserDetailsById(user models.User, id int) (*models.User, error)
}
