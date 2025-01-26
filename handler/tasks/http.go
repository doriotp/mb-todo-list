package tasks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/todo-list/models"
)

type taskHandler struct {
	tskService taskService
}

func New(tskService taskService) *taskHandler {
	return &taskHandler{tskService: tskService}
}

func (th *taskHandler) CreateTask(c *gin.Context) {
	var (
		task models.Task
	)
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to retrieve user claims"})
		return
	}

	userId := int(claims.(jwt.MapClaims)["user_id"].(float64))
	task.UserId = userId

	createTaskResponse, customErr := th.tskService.CreateTask(task)
	if customErr != nil {
		c.JSON(customErr.Code, gin.H{"message": customErr.Message})
		return
	}

	c.JSON(http.StatusCreated, createTaskResponse)
}

func (th *taskHandler) GetUserTasks(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to retrieve user claims"})
		return
	}

	userId := int(claims.(jwt.MapClaims)["user_id"].(float64))
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	tasks, customErr := th.tskService.GetUserTasks(userId, page, size)
	if customErr != nil {
		c.JSON(customErr.Code, gin.H{"message": customErr.Message})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (th *taskHandler) GetTaskById(c *gin.Context) {
	taskId, _ := strconv.Atoi(c.Param("id"))
	tasks, customErr := th.tskService.GetTaskById(taskId)
	if customErr != nil {
		c.JSON(customErr.Code, gin.H{"message": customErr.Message})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (th *taskHandler) DeleteTaskById(c *gin.Context) {
	taskId, _ := strconv.Atoi(c.Param("id"))
	customErr := th.tskService.DeleteTaskById(taskId)
	if customErr != nil {
		c.JSON(customErr.Code, gin.H{"message": customErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}

func (th *taskHandler) UpdateTaskCompletionStatus(c *gin.Context) {
	taskID, _ := strconv.Atoi(c.Param("id"))

	tasks, customErr := th.tskService.UpdateTaskCompletionStatus(taskID)
	if customErr != nil {
		c.JSON(customErr.Code, gin.H{"message": customErr.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (th *taskHandler) GetUserCompletedTasks(c *gin.Context) {
	var (
		isCompleted = true
	)

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to retrieve user claims"})
		return
	}

	userId := int(claims.(jwt.MapClaims)["user_id"].(float64))
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	CompletedTasks, customErr := th.tskService.GetUserCompletedTasks(isCompleted, userId, page, size)
	if customErr != nil {
		c.JSON(customErr.Code, gin.H{"message": customErr.Message})
		return
	}

	c.JSON(http.StatusOK, CompletedTasks)
}

func (th *taskHandler) UpdateTaskById(c *gin.Context) {
	var (
		task models.Task
	)

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	taskID, _ := strconv.Atoi(c.Param("id"))

	tasks, customErr := th.tskService.UpdateTaskById(task, taskID)
	if customErr != nil {
		c.JSON(customErr.Code, gin.H{"message": customErr.Message})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
