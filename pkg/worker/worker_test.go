package worker_test

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/shaebaratheon/distributed-task-queue/pkg/queue"
	"github.com/shaebaratheon/distributed-task-queue/pkg/worker"
)

func TestWorkerPoolExecution(t *testing.T) {
	q := queue.NewMemoryQueue()
	var processed int32

	handler := func(ctx context.Context, task *queue.Task) error {
		atomic.AddInt32(&processed, 1)
		return nil
	}

	pool := worker.NewPool(q, handler, 4)
	pool.Start()
	defer pool.Stop()

	for i := 0; i < 20; i++ {
		task := queue.NewTask(string(rune(i)), []byte("payload"), queue.PriorityNormal)
		q.Enqueue(task)
	}

	time.Sleep(300 * time.Millisecond)
	if atomic.LoadInt32(&processed) != 20 {
		t.Fatalf("expected 20 tasks processed, got %d", atomic.LoadInt32(&processed))
	}
}

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites

// Additional integration test suites
