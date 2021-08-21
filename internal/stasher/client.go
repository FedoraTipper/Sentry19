package stasher

import (
	"context"
	"fmt"
	"github.com/FedoraTipper/AntHive/internal/models"
	"github.com/go-redis/redis/v8"
)

type Stasher struct {
	redisClient *redis.Client
}

func (s *Stasher) init() {
}

func (s *Stasher) NewRedisClient(redisEndpoint, username, password string, selectedDB int) error {
	client := redis.NewClient(&redis.Options{
		Addr:     redisEndpoint,
		Username: username,
		Password: password,
		DB:       selectedDB,
	})

	s.redisClient = client

	return s.redisTestConnection(client)
}

func (s *Stasher) redisTestConnection(client *redis.Client) error {
	ctx := context.Background()

	err := client.Ping(ctx).Err()

	if err != nil {
		return err
	}

	return client.Set(ctx, "anthive_conn_test", "test", 5).Err()
}

func (s *Stasher) StashInterface(key string, miner *models.Miner) error {
	ctx := context.Background()

	err := s.redisClient.Set(ctx, key, miner, -1).Err()

	if err != nil {
		return err
	}

	return nil
}

func (s *Stasher) GetInterface(key string) (string, error) {
	ctx := context.Background()

	i, err := s.redisClient.Get(ctx, key).Result()

	fmt.Println(i)

	if err != nil {
		return "", err
	}

	return i, nil
}