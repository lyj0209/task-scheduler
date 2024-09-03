package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lyj0209/task-scheduler/internal/api/handlers"
	"github.com/lyj0209/task-scheduler/internal/scheduler"
)

func SetupRoutes(router *gin.Engine, scheduler *scheduler.Scheduler) {
	taskHandler := handlers.NewTaskHandler(scheduler)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/tasks", taskHandler.CreateTask)
		v1.POST("/tasks/new", taskHandler.CreateNewTask) // 新添加的路由
		// 其他路由...
	}
}