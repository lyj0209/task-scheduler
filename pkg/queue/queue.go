package queue

import "github.com/lyj0209/task-scheduler/internal/models"

type Queue interface {
    PublishTask(task *models.Task) error
    // 其他必要的方法...
}