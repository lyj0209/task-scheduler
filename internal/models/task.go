package models

type TaskStatus string

const (
    TaskStatusPending   TaskStatus = "pending"
    TaskStatusRunning   TaskStatus = "running"
    TaskStatusCompleted TaskStatus = "completed"
    TaskStatusFailed    TaskStatus = "failed"
)

type Task struct {
    ID     int         `json:"id"`
    Type   string      `json:"type"`
    Status TaskStatus  `json:"status"`
    Result interface{} `json:"result,omitempty"`
}