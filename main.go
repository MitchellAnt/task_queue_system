package main

import (
	"fmt"
	"log"

	"github.com/LetsGetStartedWithBub/task-queue-system/internal/queue"
)

func main() {
	// Redis queue instance
	redisQueue := queue.NewRedisQueue()

	// Queue name
	queueName := "email_tasks"

	// 📨 Task ko queue me dalna (Enqueue)
	fmt.Println("Enqueuing task: Send email to user@example.com")
	err := redisQueue.Enqueue(queueName, "Send email to user@example.com")
	if err != nil {
		log.Fatalf("Error enqueuing task: %v", err)
	}

	// 📤 Task ko queue se nikalna (Dequeue)
	task, err := redisQueue.Dequeue(queueName)
	if err != nil {
		log.Fatalf("Error dequeuing task: %v", err)
	}

	fmt.Println("Dequeued Task:", task)
}
