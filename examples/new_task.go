package new_task

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/lyj0209/task-scheduler/internal/models"
)

type NewTaskPayload struct {
	Param1 string `json:"param1"`
	Param2 int    `json:"param2"`
}

func NewNewTask(param1 string, param2 int) *models.Task {
	payload := NewTaskPayload{
		Param1: param1,
		Param2: param2,
	}

	payloadBytes, _ := json.Marshal(payload)

	return &models.Task{
		Name:              "New Task",
		Type:              "new_task",
		Payload:           string(payloadBytes),
		Priority:          2,
		EstimatedDuration: 180, // 3 minutes
	}
}

func ExecuteNewTask(task *models.Task) error {
	var payload NewTaskPayload
	err := json.Unmarshal([]byte(task.Payload.(string)), &payload)
	if err != nil {
		return err
	}

	// 实际的任务执行逻辑
	fmt.Printf("Executing new task with param1: %s and param2: %d\n", 
		payload.Param1, payload.Param2)
	time.Sleep(time.Duration(task.EstimatedDuration) * time.Second)

	fmt.Println("New task completed")
	return nil
}