package usecase

import (
	"bank-ina-assessment/model"
	"bank-ina-assessment/repository/database"
	"bank-ina-assessment/util"
	"fmt"
)

type TaskUsecase interface {
	CreateTask(task *model.Task) util.Error
	GetTask(userId, id uint) (task model.Task, err util.Error)
	GetListTasks(userId uint) (tasks []model.Task, err util.Error)
	UpdateTask(task *model.Task) (err util.Error)
	DeleteTask(userId, id uint) (err util.Error)
}

type taskUsecase struct {
	taskRepository database.TaskRepository
}

func NewTaskUsecase(taskRepo database.TaskRepository) *taskUsecase {
	return &taskUsecase{taskRepository: taskRepo}
}

func (b *taskUsecase) CreateTask(task *model.Task) (err util.Error) {

	if err = b.taskRepository.CreateTask(task); err.NotNil() {
		return err
	}

	return
}

func (b *taskUsecase) GetTask(userId, id uint) (task model.Task, err util.Error) {
	task, err = b.taskRepository.GetTask(userId, id)
	if err.NotNil() {
		fmt.Println("GetTask: Error getting task from database")
		return
	}
	return
}

func (b *taskUsecase) GetListTasks(userID uint) (tasks []model.Task, err util.Error) {
	tasks, err = b.taskRepository.GetTasks(userID)
	if err.NotNil() {
		fmt.Println("GetListTasks: Error getting tasks from database")
		return
	}
	return
}

func (b *taskUsecase) UpdateTask(task *model.Task) (err util.Error) {
	err = b.taskRepository.UpdateTask(task)
	if err.NotNil() {
		fmt.Println("UpdateTask : Error updating task, err: ", err)
		return
	}

	return
}

func (b *taskUsecase) DeleteTask(userId, id uint) (err util.Error) {
	task := model.Task{}
	task.ID = id
	task.UserID = userId
	err = b.taskRepository.DeleteTask(&task)
	if err.NotNil() {
		fmt.Println("DeleteTask : error deleting task, err: ", err)
		return
	}

	return
}
