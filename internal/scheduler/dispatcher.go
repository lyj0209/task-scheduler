package scheduler

import (
	"context"
	"time"

	"github.com/lyj0209/task-scheduler/internal/models"
	"github.com/lyj0209/task-scheduler/pkg/queue"
	"github.com/lyj0209/task-scheduler/pkg/discovery"
)

type Scheduler struct {
	queue     queue.Queue
	discovery discovery.Discovery
	storage   storage.Storage
}

func NewScheduler(q queue.Queue, d discovery.Discovery, s storage.Storage) *Scheduler {
	return &Scheduler{
		queue:     q,
		discovery: d,
		storage:   s,
	}
}

func (s *Scheduler) Start(ctx context.Context) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := s.dispatchTasks(); err != nil {
				// Log error
			}
		}
	}
}

func (s *Scheduler) dispatchTasks() error {
	tasks, err := s.storage.GetPendingTasks()
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if err := s.queue.PublishTask(task); err != nil {
			// Log error and continue
			continue
		}
		task.Status = models.TaskStatusRunning
		if err := s.storage.UpdateTask(task); err != nil {
			// Log error
		}
	}

	return nil
}

func (s *Scheduler) SubmitTask(task *models.Task) error {
	task.Status = models.TaskStatusPending
	return s.storage.CreateTask(task)
}