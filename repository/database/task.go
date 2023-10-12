package database

import (
	"bank-ina-assessment/config"
	"bank-ina-assessment/model"
	"bank-ina-assessment/util"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *model.Task) util.Error
	GetTasks(userId uint) (tasks []model.Task, err util.Error)
	GetTask(userId, id uint) (task model.Task, err util.Error)
	UpdateTask(task *model.Task) util.Error
	DeleteTask(task *model.Task) util.Error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (u *taskRepository) CreateTask(task *model.Task) (err util.Error) {
	if errDb := config.DB.Create(task).Error; errDb != nil {
		err.StatusInternalServerError(errDb, "Error to create task")
		return
	}
	return
}

func (u *taskRepository) GetTasks(userId uint) (tasks []model.Task, err util.Error) {
	if errDb := config.DB.Find(&tasks).Error; errDb != nil {
		err.StatusInternalServerError(errDb, "Error to get tasks")
		return
	}
	return
}

func (u *taskRepository) GetTask(userId, id uint) (task model.Task, err util.Error) {
	task.ID = id
	if errDb := config.DB.First(&task).Error; errDb != nil {
		if errDb == gorm.ErrRecordNotFound {
			err.StatusNotFound(errDb, "Task not found")
			return
		}
		err.StatusInternalServerError(errDb, "Error to get task")
		return
	}
	return
}

func (u *taskRepository) UpdateTask(task *model.Task) (err util.Error) {
	if errDb := config.DB.Updates(task).Error; errDb != nil {
		err.StatusInternalServerError(errDb, "Error update task")
		return err
	}
	return
}

func (u *taskRepository) DeleteTask(task *model.Task) (err util.Error) {
	if errDb := config.DB.Delete(task).Error; errDb != nil {
		err.StatusInternalServerError(errDb, "Error delete task")
		return err
	}
	return
}
