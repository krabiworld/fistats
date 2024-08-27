package fistorage

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(addr, password string, db int) *Redis {
	return &Redis{client: redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})}
}

func (r *Redis) Increment(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return r.client.Incr(ctx, key).Err()
}

func (r *Redis) GetAll() (map[string]uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	keys, err := r.client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	result := make(map[string]uint64, len(keys))
	for _, key := range keys {
		value, err := r.client.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		result[key], err = strconv.ParseUint(value, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (r *Redis) DeleteAll() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return r.client.FlushDB(ctx).Err()
}

func (r *Redis) Close() error {
	return r.client.Close()
}
