package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/lyj0209/task-scheduler/internal/models"
	"github.com/lyj0209/task-scheduler/pkg/queue"
	"github.com/lyj0209/task-scheduler/examples/new_task"
	// 导入其他任务类型...
)

type Worker struct {
	id    string
	queue queue.Queue
}

func NewWorker(id string, q queue.Queue) *Worker {
	return &Worker{
		id:    id,
		queue: q,
	}
}

func (w *Worker) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			task, err := w.queue.ConsumeTask()
			if err != nil {
				log.Printf("Error consuming task: %v", err)
				continue
			}

			if err := w.executeTask(task); err != nil {
				log.Printf("Error executing task: %v", err)
				// 处理任务执行失败的逻辑
			}
		}
	}
}

func (w *Worker) executeTask(task *models.Task) error {
	switch task.Type {
	case "new_task":
		return new_task.ExecuteNewTask(task)
	// 其他任务类型...
	default:
		return fmt.Errorf("unknown task type: %s", task.Type)
	}
}