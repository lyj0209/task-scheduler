package main

import (
    "github.com/lyj0209/task-scheduler/internal/scheduler"
    "github.com/lyj0209/task-scheduler/internal/storage/mysql"
    "github.com/lyj0209/task-scheduler/pkg/queue"
    "log"
)

func main() {
    mysqlStorage, err := mysql.NewMySQLStorage("user:password@tcp(localhost:3306)/ecommerce")
    if err != nil {
        log.Fatalf("Failed to connect to MySQL: %v", err)
    }

    queue := queue.NewMemoryQueue()

    scheduler := scheduler.NewScheduler(mysqlStorage, queue)
    scheduler.Start()
}
