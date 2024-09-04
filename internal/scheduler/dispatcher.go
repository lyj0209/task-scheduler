package scheduler

import (
    "github.com/lyj0209/task-scheduler/internal/models"
    "github.com/lyj0209/task-scheduler/internal/storage/mysql"
    "github.com/lyj0209/task-scheduler/pkg/queue"
    "log"
    "time"
)

type Scheduler struct {
    mysqlStorage *mysql.MySQLStorage
    queue        queue.Queue
}

func NewScheduler(mysqlStorage *mysql.MySQLStorage, queue queue.Queue) *Scheduler {
    return &Scheduler{
        mysqlStorage: mysqlStorage,
        queue:        queue,
    }
}

func (s *Scheduler) Start() {
    ticker := time.NewTicker(1 * time.Minute)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            s.scheduleTasks()
        }
    }
}

func (s *Scheduler) scheduleTasks() {
    tasks := []*models.Task{
        {Type: "update_order_count", Status: models.TaskStatusPending},
        {Type: "update_hot_products", Status: models.TaskStatusPending},
    }

    for _, task := range tasks {
        err := s.mysqlStorage.CreateTask(task)
        if err != nil {
            log.Printf("Error creating task: %v", err)
            continue
        }

        err = s.queue.PublishTask(task)
        if err != nil {
            log.Printf("Error publishing task: %v", err)
        }
    }
}