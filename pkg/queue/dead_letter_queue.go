package queue

import (
	"sync"
	"time"
)

type DeadLetterQueue struct {
	mu     sync.RWMutex
	failed []*Task
}

func NewDeadLetterQueue() *DeadLetterQueue {
	return &DeadLetterQueue{
		failed: make([]*Task, 0),
	}
}

func (dlq *DeadLetterQueue) Push(t *Task) {
	dlq.mu.Lock()
	defer dlq.mu.Unlock()
	t.Status = StatusFailed
	t.UpdatedAt = time.Now().UTC()
	dlq.failed = append(dlq.failed, t)
}

func (dlq *DeadLetterQueue) List() []*Task {
	dlq.mu.RLock()
	defer dlq.mu.RUnlock()
	res := make([]*Task, len(dlq.failed))
	copy(res, dlq.failed)
	return res
}

func (dlq *DeadLetterQueue) Size() int {
	dlq.mu.RLock()
	defer dlq.mu.RUnlock()
	return len(dlq.failed)
}

func (dlq *DeadLetterQueue) Replay(taskID string, targetQueue *MemoryQueue) bool {
	dlq.mu.Lock()
	defer dlq.mu.Unlock()

	for i, t := range dlq.failed {
		if t.ID == taskID {
			dlq.failed = append(dlq.failed[:i], dlq.failed[i+1:]...)
			t.RetryCount = 0
			t.Status = StatusPending
			t.LastError = ""
			t.ScheduledFor = time.Now().UTC()
			targetQueue.Enqueue(t)
			return true
		}
	}
	return false
}

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry

// Dead letter queue persistence & replay telemetry
