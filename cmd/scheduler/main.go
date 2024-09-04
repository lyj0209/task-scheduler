package main

import (
    "github.com/lyj0209/task-scheduler/internal/scheduler"
    "github.com/lyj0209/task-scheduler/internal/storage/mysql"
    "github.com/lyj0209/task-scheduler/pkg/queue/kafka"
    "log"
)

func main() {
    mysqlStorage, err := mysql.NewMySQLStorage("root:yourpassword@tcp(localhost:3306)/ecommerce")
    if err != nil {
        log.Fatalf("Failed to connect to MySQL: %v", err)
    }

    kafkaQueue, err := kafka.NewKafkaQueue([]string{"localhost:9092"}, "tasks")
    if err != nil {
        log.Fatalf("Failed to create Kafka queue: %v", err)
    }
    defer kafkaQueue.Close()

    scheduler := scheduler.NewScheduler(mysqlStorage, kafkaQueue)
    scheduler.Start()
}