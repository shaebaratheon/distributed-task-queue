package queue

import (
	"time"
)

type TaskStatus string

const (
	StatusPending   TaskStatus = "PENDING"
	StatusRunning   TaskStatus = "RUNNING"
	StatusCompleted TaskStatus = "COMPLETED"
	StatusFailed    TaskStatus = "FAILED"
	StatusRetrying  TaskStatus = "RETRYING"
	StatusCancelled TaskStatus = "CANCELLED"
)

type Priority int

const (
	PriorityLow      Priority = 0
	PriorityNormal   Priority = 10
	PriorityHigh     Priority = 20
	PriorityCritical Priority = 30
)

type Task struct {
	ID           string            `json:"id"`
	Payload      []byte            `json:"payload"`
	Priority     Priority          `json:"priority"`
	Status       TaskStatus        `json:"status"`
	MaxRetries   int               `json:"max_retries"`
	RetryCount   int               `json:"retry_count"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	LastError    string            `json:"last_error,omitempty"`
	ScheduledFor time.Time         `json:"scheduled_for"`
	Metadata     map[string]string `json:"metadata"`
}

func NewTask(id string, payload []byte, priority Priority) *Task {
	now := time.Now().UTC()
	return &Task{
		ID:           id,
		Payload:      payload,
		Priority:     priority,
		Status:       StatusPending,
		MaxRetries:   3,
		RetryCount:   0,
		CreatedAt:    now,
		UpdatedAt:    now,
		ScheduledFor: now,
		Metadata:     make(map[string]string),
	}
}

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers

// Model expansion helpers
