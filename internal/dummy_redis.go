package internal

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type DummyRedis struct {
	client *redis.Client
	ctx    context.Context
}

func NewDummyRedis(host string, port int) (*DummyRedis, error) {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		return nil, err
	}

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get PONG: %s", err)
	}
	if pong != "PONG" {
		return nil, errors.New("failed to get PONG")
	}

	return &DummyRedis{
		client: client,
		ctx:    ctx,
	}, nil
}

func (dr *DummyRedis) TestConnection() error {
	if err := dr.client.Set(dr.ctx, "dummy-key", "DUMMY_VALUE", 0).Err(); err != nil {
		return fmt.Errorf("failed to get dummy-key: %w", err)
	}

	val, err := dr.client.Get(dr.ctx, "dummy-key").Result()
	if err != nil {
		return fmt.Errorf("failed to get dummy-key: %w", err)
	}
	fmt.Println("dummy-key", val)

	val2, err := dr.client.Get(dr.ctx, "dummy-key2").Result()
	if err == redis.Nil {
		fmt.Println("dummy-key2 does not exist, ok")
	} else if err != nil {
		return fmt.Errorf("failed to get dummy-key2: %w", err)
	} else {
		fmt.Println("dummy-key2, not ok :/", val2)
	}

	return nil
}
