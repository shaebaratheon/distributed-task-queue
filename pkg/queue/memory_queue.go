package queue

import (
	"context"
	"errors"
	"sort"
	"sync"
	"time"
)

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrQueueEmpty   = errors.New("queue is empty")
	ErrQueueClosed  = errors.New("queue is closed")
)

type MemoryQueue struct {
	mu     sync.RWMutex
	tasks  map[string]*Task
	ready  []*Task
	notify chan struct{}
	closed bool
}

func NewMemoryQueue() *MemoryQueue {
	return &MemoryQueue{
		tasks:  make(map[string]*Task),
		ready:  make([]*Task, 0),
		notify: make(chan struct{}, 1),
	}
}

func (q *MemoryQueue) Enqueue(t *Task) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.closed {
		return ErrQueueClosed
	}

	q.tasks[t.ID] = t
	q.ready = append(q.ready, t)
	q.sortReady()

	select {
	case q.notify <- struct{}{}:
	default:
	}
	return nil
}

func (q *MemoryQueue) Dequeue(ctx context.Context) (*Task, error) {
	for {
		q.mu.Lock()
		if q.closed {
			q.mu.Unlock()
			return nil, ErrQueueClosed
		}

		now := time.Now().UTC()
		var candidateIdx = -1
		for i, t := range q.ready {
			if t.Status == StatusPending && !t.ScheduledFor.After(now) {
				candidateIdx = i
				break
			}
		}

		if candidateIdx != -1 {
			task := q.ready[candidateIdx]
			q.ready = append(q.ready[:candidateIdx], q.ready[candidateIdx+1:]...)
			task.Status = StatusRunning
			task.UpdatedAt = time.Now().UTC()
			q.mu.Unlock()
			return task, nil
		}

		q.mu.Unlock()

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-q.notify:
		case <-time.After(50 * time.Millisecond):
		}
	}
}

func (q *MemoryQueue) UpdateStatus(taskID string, status TaskStatus, errMsg string) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	task, exists := q.tasks[taskID]
	if !exists {
		return ErrTaskNotFound
	}

	task.Status = status
	task.UpdatedAt = time.Now().UTC()
	task.LastError = errMsg
	return nil
}

func (q *MemoryQueue) sortReady() {
	sort.SliceStable(q.ready, func(i, j int) bool {
		if q.ready[i].Priority != q.ready[j].Priority {
			return q.ready[i].Priority > q.ready[j].Priority
		}
		return q.ready[i].CreatedAt.Before(q.ready[j].CreatedAt)
	})
}

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives

// Queue synchronization primitives
