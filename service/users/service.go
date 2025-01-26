package users

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	customerrors "github.com/todo-list/customErrors"
	"github.com/todo-list/models"
	"github.com/todo-list/utils"
)

type userService struct {
	usrStore userStore
}

func New(usrStore userStore) *userService {
	return &userService{usrStore: usrStore}
}

func (us *userService) CreateUser(user models.User) *customerrors.Error {
	if user.Email == "" || user.Password == "" || user.Name == "" {
		return customerrors.New(http.StatusBadRequest, "invalid input")
	}

	_, err := us.usrStore.GetUserByEmail(user.Email)
	if err != nil {
		return customerrors.New(http.StatusInternalServerError, err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return customerrors.New(http.StatusInternalServerError, err.Error())
	}

	user.Password = string(hashedPassword)

	err = us.usrStore.CreateUser(user)
	if err != nil {
		return customerrors.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (us *userService) Login(loginRequest models.LoginRequest) (*models.LoginResponse, *customerrors.Error) {
	if loginRequest.Email == "" || loginRequest.Password == "" {
		return nil, customerrors.New(http.StatusBadRequest, "invalid email or password")
	}

	userInfo, err := us.usrStore.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return nil, customerrors.New(http.StatusInternalServerError, err.Error())
	}

	if userInfo == nil || bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(loginRequest.Password)) != nil {
		return nil, customerrors.New(http.StatusBadRequest, "user does not exist")
	}

	token, err := utils.GenerateToken(userInfo.ID)
	if err != nil {
		return nil, customerrors.New(http.StatusInternalServerError, err.Error())
	}

	resp := models.LoginResponse{
		Email: loginRequest.Email,
		Token: token,
	}

	return &resp, nil

}

func (us *userService) ForgotPassword(fpr models.ForgotPasswordRequest) *customerrors.Error {
	if fpr.Email == "" {
		return customerrors.New(http.StatusBadRequest, "invalid email or password")
	}
	userInfo, err := us.usrStore.GetUserByEmail(fpr.Email)
	if err != nil {
		return customerrors.New(http.StatusInternalServerError, err.Error())
	}

	if userInfo == nil {
		return customerrors.New(http.StatusBadRequest, "user not found")
	}

	token, err := utils.GenerateToken(userInfo.ID)
	if err != nil {
		return customerrors.New(http.StatusInternalServerError, err.Error())
	}

	err = utils.SendEmail(token, fpr.Email)
	if err != nil {
		return customerrors.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (us *userService) ResetPassword(fpr models.ResetPasswordRequest, id int) *customerrors.Error {

	if fpr.NewPassword == "" || fpr.ConfirmNewPassword == "" {
		return customerrors.New(http.StatusBadRequest, "invalid password or new password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(fpr.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return customerrors.New(http.StatusInternalServerError, err.Error())
	}

	err = us.usrStore.UpdatePasswordById(string(hashedPassword), id)
	if err != nil {
		return customerrors.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (us *userService) GetCurrentUser(id int) (*models.User, error) {

	user, err := us.usrStore.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (us *userService) UpdateUserDetailsById(user models.User, id int) (*models.User, *customerrors.Error) {

	exisitingUser, err := us.usrStore.GetUserById(id)
	if err != nil {
		return nil, customerrors.New(http.StatusInternalServerError, err.Error())
	}

	if exisitingUser == nil {
		return nil, customerrors.New(http.StatusBadRequest, "user not found")
	}

	updatedUserDetails, err := us.usrStore.UpdateUserDetailsById(user, id)
	if err != nil {
		return nil, customerrors.New(http.StatusInternalServerError, err.Error())
	}

	return updatedUserDetails, nil
}
