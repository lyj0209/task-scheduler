package storage

import (
	"github.com/lyj0209/task-scheduler/internal/models"
)

type Storage interface {
	CreateTask(task *models.Task) error
	GetPendingTasks() ([]*models.Task, error)
	UpdateTask(task *models.Task) error
	// 其他方法...
}