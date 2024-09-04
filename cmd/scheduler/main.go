package main

import (
    "github.com/lyj0209/task-scheduler/internal/scheduler"
    "github.com/lyj0209/task-scheduler/internal/storage/mysql"
    "github.com/lyj0209/task-scheduler/pkg/queue/kafka"
    "log"
    "os"
    "fmt"
)

func main() {
    mysqlHost := os.Getenv("MYSQL_HOST")
    mysqlPort := os.Getenv("MYSQL_PORT")
    mysqlUser := os.Getenv("MYSQL_USER")
    mysqlPassword := os.Getenv("MYSQL_PASSWORD")
    mysqlDatabase := os.Getenv("MYSQL_DATABASE")
    kafkaBrokers := os.Getenv("KAFKA_BROKERS")
    mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)
    mysqlStorage, err := mysql.NewMySQLStorage(mysqlDSN)
    if err != nil {
        log.Fatalf("Failed to connect to MySQL: %v", err)
    }

    kafkaQueue, err := kafka.NewKafkaQueue([]string{kafkaBrokers}, "tasks")
    if err != nil {
        log.Fatalf("Failed to create Kafka queue: %v", err)
    }
    defer kafkaQueue.Close()

    scheduler := scheduler.NewScheduler(mysqlStorage, kafkaQueue)
    scheduler.Start()


    // mysqlStorage, err := mysql.NewMySQLStorage("root:yourpassword@tcp(mysql:3306)/ecommerce")
    // if err != nil {
    //     log.Fatalf("Failed to connect to MySQL: %v", err)
    // }

    // kafkaQueue, err := kafka.NewKafkaQueue([]string{"kafka:29092"}, "tasks")
    // if err != nil {
    //     log.Fatalf("Failed to create Kafka queue: %v", err)
    // }
    // defer kafkaQueue.Close()

    // scheduler := scheduler.NewScheduler(mysqlStorage, kafkaQueue)
    // scheduler.Start()
}