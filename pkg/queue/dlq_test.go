package queue_test

import (
	"testing"
	"github.com/shaebaratheon/distributed-task-queue/pkg/queue"
)

func TestDeadLetterQueue(t *testing.T) {
	dlq := queue.NewDeadLetterQueue()
	task := queue.NewTask("task_fail_1", []byte("bad data"), queue.PriorityLow)
	dlq.Push(task)

	if dlq.Size() != 1 {
		t.Fatalf("expected 1 dead letter task, got %d", dlq.Size())
	}

	q := queue.NewMemoryQueue()
	if !dlq.Replay("task_fail_1", q) {
		t.Fatalf("failed to replay task")
	}

	if dlq.Size() != 0 {
		t.Fatalf("expected DLQ empty after replay")
	}
}

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests

// Dead letter queue unit tests
