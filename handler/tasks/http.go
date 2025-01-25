package tasks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/todo-list/models"
	"github.com/todo-list/utils"
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

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		return
	}

	claims, err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	userId := int(claims["user_id"].(float64))
	task.UserId=userId

	customErr := th.tskService.CreateTask(task)
	if err != nil {
		c.JSON(customErr.Code, gin.H{"message": customErr.Message})
	}

	c.JSON(http.StatusCreated, task)
}

func (th *taskHandler) GetUserTasks(c *gin.Context) {
	token := c.Request.Header["Authorization"]

	claims, err := utils.VerifyToken(token[0])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	userId := int(claims["user_id"].(float64))
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	tasks, customErr := th.tskService.GetUserTasks(userId, page, size)
	if err != nil {
		c.JSON(customErr.Code, gin.H{"message": customErr.Message})
	}

	c.JSON(http.StatusOK, tasks)
}

func (th *taskHandler) GetTaskById(c *gin.Context) {
	taskId, _ := strconv.Atoi(c.Param("id"))
	tasks, customErr := th.tskService.GetTaskById(taskId)
	if customErr != nil {
		c.JSON(customErr.Code, gin.H{"message": customErr.Message})
	}

	c.JSON(http.StatusOK, tasks)
}

func (th *taskHandler) DeleteTaskById(c *gin.Context) {
	taskId, _ := strconv.Atoi(c.Param("id"))
	customErr := th.tskService.DeleteTaskById(taskId)
	if customErr != nil {
		c.JSON(customErr.Code, gin.H{"message": customErr.Message})
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}

func (th *taskHandler) UpdateTaskCompletionStatus(c *gin.Context) {
	taskID, _ := strconv.Atoi(c.Param("id"))

	token := c.GetHeader("authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		return
	}

	_, err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	tasks, customErr := th.tskService.UpdateTaskCompletionStatus(taskID)
	if customErr != nil {
		c.JSON(customErr.Code, gin.H{"error": customErr.Error()})
	}

	c.JSON(http.StatusOK, tasks)
}

// 	tasks, err := ts.tskStore.GetTaskById(taskId)
// 	if err != nil {
// 		return nil, &customerrors.Error{Code: http.StatusInternalServerError, Message: err.Error()}
// 	}

// 	return tasks, nil

// }

func (th *taskHandler) GetUserCompletedTasks(c *gin.Context) {
	var (
		isCompleted = true
	)

	token := c.Request.Header["Authorization"]

	claims, err := utils.VerifyToken(token[0])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	userId := int(claims["user_id"].(float64))
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	CompletedTasks, customErr := th.tskService.GetUserCompletedTasks(isCompleted, userId, page, size)
	if err != nil {
		c.JSON(customErr.Code, gin.H{"message": customErr.Message})
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
	}

	c.JSON(http.StatusOK, tasks)
}
