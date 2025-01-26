package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/todo-list/models"
	"github.com/todo-list/utils"
)

type userHandler struct {
	usrService userService
}

func New(usrService userService) *userHandler {
	return &userHandler{usrService: usrService}
}

func (uH *userHandler) Register(c *gin.Context) {

	var (
		user models.User
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uH.usrService.CreateUser(user)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	resp := models.RegisterResponse{
		Name:       user.Name,
		Email:      user.Email,
		Country:    user.Country,
		Occupation: user.Occupation,
		Phone:      user.Phone,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"data":    resp,
	})
}

func (uH *userHandler) Login(c *gin.Context) {
	var (
		loginRequest models.LoginRequest
	)

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resp, err := uH.usrService.Login(loginRequest)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (uH *userHandler) ForgotPassword(c *gin.Context) {
	var (
		fpr models.ForgotPasswordRequest
	)

	if err := c.ShouldBindJSON(&fpr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uH.usrService.ForgotPassword(fpr)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message": "email sent successfully"})
}

func (uH *userHandler) ResetPassword(c *gin.Context) {
	var (
		rpr models.ResetPasswordRequest
	)

	if err := c.ShouldBindJSON(&rpr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token := c.Query("token")

	claims, err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return 
	}

	id := int(claims["user_id"].(float64))

	customErr := uH.usrService.ResetPassword(rpr, id)
	if customErr != nil {
		c.JSON(customErr.Code, gin.H{"message": customErr.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message": "password reset successfully"})

}

func (uH *userHandler) Logout(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "user logged out succesfullly"})

}

func (uH *userHandler) GetCurrentUser(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to retrieve user claims"})
		return
	}

	id := int(claims.(jwt.MapClaims)["user_id"].(float64))

	user, err := uH.usrService.GetCurrentUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, user)
}

func (uH *userHandler) UpdateUserDetailsById(c *gin.Context) {
	var (
		user models.User
	)

	// Check user credentials and generate a JWT token
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	updatedUserDetails, err := uH.usrService.UpdateUserDetailsById(user, id)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return 
	}

	c.JSON(http.StatusOK, updatedUserDetails)
}
