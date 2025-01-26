package users

import (
	customerrors "github.com/todo-list/customErrors"
	"github.com/todo-list/models"
)

type userService interface {
	CreateUser(models.User) *customerrors.Error
	Login(user models.LoginRequest) (*models.LoginResponse, *customerrors.Error) 
	ForgotPassword(fpr models.ForgotPasswordRequest) (*customerrors.Error)
	ResetPassword(fpr models.ResetPasswordRequest, id int) (*customerrors.Error)
	GetCurrentUser(id int)(*models.User, error) 
	UpdateUserDetailsById(user models.User, id int)(*models.User,*customerrors.Error )
}
