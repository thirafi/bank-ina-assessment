package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"bank-ina-assessment/model"
	"bank-ina-assessment/repository/database"
	"bank-ina-assessment/usecase"
)

type TaskController interface {
	GetTasksController(c *gin.Context) error
	GetTaskController(c *gin.Context) error
	CreateTaskController(c *gin.Context) error
	DeleteTaskController(c *gin.Context) error
	UpdateTaskController(c *gin.Context) error
}

type taskController struct {
	taskUsecase    usecase.TaskUsecase
	taskRepository database.TaskRepository
}

func NewTaskController(
	taskUsecase usecase.TaskUsecase,
	taskRepository database.TaskRepository,
) *taskController {
	return &taskController{
		taskUsecase,
		taskRepository,
	}
}

func (b *taskController) GetTasksController(c *gin.Context) {
	user, _ := c.Get("user")
	userID := user.(model.User).ID
	tasks, err := b.taskUsecase.GetListTasks(userID)
	if err.NotNil() {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"tasks":  tasks,
	})
}

func (b *taskController) GetTaskController(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user, _ := c.Get("user")
	userID := user.(model.User).ID
	task, err := b.taskUsecase.GetTask(userID, uint(id))

	if err.NotNil() {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"tasks":  task,
	})
}

// create new task
func (b *taskController) CreateTaskController(c *gin.Context) {
	task := model.Task{}
	c.Bind(&task)
	user, _ := c.Get("user")
	task.UserID = user.(model.User).ID

	if err := b.taskUsecase.CreateTask(&task); err.NotNil() {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new task",
		"task":    task,
	})
}

// delete task by id
func (b *taskController) DeleteTaskController(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user, _ := c.Get("user")
	userID := user.(model.User).ID
	if err := b.taskUsecase.DeleteTask(userID, uint(id)); err.NotNil() {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete task",
	})
}

// update task by id
func (b *taskController) UpdateTaskController(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	task := model.Task{}
	c.Bind(&task)
	task.ID = uint(id)
	user, _ := c.Get("user")
	task.UserID = user.(model.User).ID
	if err := b.taskUsecase.UpdateTask(&task); err.NotNil() {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update task",
	})
}
