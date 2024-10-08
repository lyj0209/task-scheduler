package main

import (
    "github.com/lyj0209/task-scheduler/internal/storage/mysql"
    "github.com/lyj0209/task-scheduler/internal/storage/redis"
    "github.com/lyj0209/task-scheduler/internal/worker"
    "github.com/lyj0209/task-scheduler/pkg/queue/kafka"
    "log"
)

func main() {
    mysqlStorage, err := mysql.NewMySQLStorage("root:yourpassword@tcp(localhost:3306)/ecommerce")
    if err != nil {
        log.Fatalf("Failed to connect to MySQL: %v", err)
    }

    redisStorage, err := redis.NewRedisStorage("localhost:6379")
    if err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }

    kafkaQueue, err := kafka.NewKafkaQueue([]string{"kafka:29092"}, "tasks")
    if err != nil {
        log.Fatalf("Failed to create Kafka queue: %v", err)
    }
    defer kafkaQueue.Close()

    worker := worker.NewWorker(mysqlStorage, redisStorage, kafkaQueue)
    worker.Start()
}