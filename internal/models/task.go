package models

import (
	"time"
)

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusFailed    TaskStatus = "failed"
)

type Task struct {
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	Type              string      `json:"type"`
	Payload           interface{} `json:"payload"`
	Status            TaskStatus  `json:"status"`
	Priority          int         `json:"priority"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	ScheduledAt       time.Time   `json:"scheduled_at"`
	EstimatedDuration int         `json:"estimated_duration"` // in seconds
}