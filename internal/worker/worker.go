package worker

import (
    "github.com/lyj0209/task-scheduler/internal/models"
    "github.com/lyj0209/task-scheduler/internal/storage/mysql"
    "github.com/lyj0209/task-scheduler/internal/storage/redis"
    "github.com/lyj0209/task-scheduler/pkg/queue"
    "log"
    "time"
)

type Worker struct {
    mysqlStorage *mysql.MySQLStorage
    redisStorage *redis.RedisStorage
    queue        queue.Queue
}

func NewWorker(mysqlStorage *mysql.MySQLStorage, redisStorage *redis.RedisStorage, queue queue.Queue) *Worker {
    return &Worker{
        mysqlStorage: mysqlStorage,
        redisStorage: redisStorage,
        queue:        queue,
    }
}

func (w *Worker) Start() {
    for {
        task, err := w.queue.ConsumeTask()
        if err != nil {
            log.Printf("Error consuming task: %v", err)
            time.Sleep(5 * time.Second)
            continue
        }

        if task == nil {
            time.Sleep(5 * time.Second)
            continue
        }

        log.Printf("Processing task: %+v", task)

        err = w.executeTask(task)
        if err != nil {
            log.Printf("Error executing task: %v", err)
            task.Status = models.TaskStatusFailed
        } else {
            task.Status = models.TaskStatusCompleted
        }

        err = w.mysqlStorage.UpdateTask(task)
        if err != nil {
            log.Printf("Error updating task: %v", err)
        }
    }
}

func (w *Worker) executeTask(task *models.Task) error {
    switch task.Type {
    case "update_order_count":
        return w.updateOrderCount(task)
    case "update_hot_products":
        return w.updateHotProducts(task)
    default:
        return nil
    }
}

func (w *Worker) updateOrderCount(task *models.Task) error {
    count, err := w.mysqlStorage.GetOrderCount24h()
    if err != nil {
        return err
    }

    err = w.redisStorage.SetOrderCount24h(count)
    if err != nil {
        return err
    }

    task.Result = map[string]interface{}{
        "order_count_24h": count,
    }
    return nil
}

func (w *Worker) updateHotProducts(task *models.Task) error {
    hotProducts, err := w.mysqlStorage.GetHotProducts(10)
    if err != nil {
        return err
    }

    err = w.redisStorage.UpdateHotProducts(hotProducts)
    if err != nil {
        return err
    }

    task.Result = map[string]interface{}{
        "hot_products": hotProducts,
    }
    return nil
}