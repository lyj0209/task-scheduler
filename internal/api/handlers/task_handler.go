package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/distributed-task-scheduler/examples/new_task"
	"github.com/yourusername/distributed-task-scheduler/internal/scheduler"
)

type TaskHandler struct {
	scheduler *scheduler.Scheduler
}

func NewTaskHandler(s *scheduler.Scheduler) *TaskHandler {
	return &TaskHandler{scheduler: s}
}

func (h *TaskHandler) CreateNewTask(c *gin.Context) {
	var request struct {
		Param1 string `json:"param1" binding:"required"`
		Param2 int    `json:"param2" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := new_task.NewNewTask(request.Param1, request.Param2)

	if err := h.scheduler.SubmitTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}