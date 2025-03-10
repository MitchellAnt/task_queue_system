package queue

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisQueue struct {
	client *redis.Client
}

func NewRedisQueue() *RedisQueue {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &RedisQueue{client: client}
}

// Enqueue Task
func (r *RedisQueue) Enqueue(queueName string, task string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	err := r.client.LPush(ctx, queueName, task).Err()
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %v", err)
	}
	return nil
}

// Dequeue Task
func (r *RedisQueue) Dequeue(queueName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	task, err := r.client.RPop(ctx, queueName).Result()
	if err != nil {
		return "", fmt.Errorf("failed to dequeue task: %v", err)
	}
	return task, nil
}
