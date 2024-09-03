// cmd/scheduler/main.go

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lyj0209/task-scheduler/internal/api"
	"github.com/lyj0209/task-scheduler/internal/scheduler"
	"github.com/lyj0209/task-scheduler/internal/storage/mysql"
	"github.com/lyj0209/task-scheduler/pkg/queue/kafka"
	"github.com/lyj0209/task-scheduler/pkg/discovery/etcd"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化组件
	mysqlStorage, err := mysql.NewMySQLStorage("connection_string")
	if err != nil {
		log.Fatalf("Failed to initialize MySQL storage: %v", err)
	}

	kafkaQueue, err := kafka.NewKafkaQueue("kafka_brokers")
	if err != nil {
		log.Fatalf("Failed to initialize Kafka queue: %v", err)
	}

	etcdClient, err := etcd.NewEtcdClient("etcd_endpoints")
	if err != nil {
		log.Fatalf("Failed to initialize etcd client: %v", err)
	}

	// 创建调度器
	sched := scheduler.NewScheduler(mysqlStorage, kafkaQueue, etcdClient)

	// 创建API服务器
	router := gin.Default()
	api.SetupRoutes(router, sched)

	// 创建上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动调度器
	go func() {
		if err := sched.Start(ctx); err != nil {
			log.Printf("Scheduler stopped: %v", err)
		}
	}()

	// 启动API服务器
	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Printf("API server stopped: %v", err)
		}
	}()

	// 等待中断信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	log.Println("Shutting down...")
	cancel()
	// 执行清理操作
}