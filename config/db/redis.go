package db

import (
	"github.com/go-redis/redis"
	"sync"
	"time"
)

type RedisCache struct {
	db *redis.Client
	mu sync.RWMutex
}

func NewRedisCache(addr string, password string, db int) (Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:       addr,
		PoolSize:   100,
		MaxRetries: 2,
		Password:   password, // no password set
		DB:         db,       // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	cache = &RedisCache{
		db: client,
	}
	return cache, nil
}

func (r *RedisCache) Write(key string, value []byte, d time.Duration) error {
	err := r.db.Set(key, value, d).Err()
	return err
}

func (r *RedisCache) Get(key string) ([]byte, error) {
	data, err := r.db.Get(key).Bytes()
	return data, err
}

func (r *RedisCache) IsCached(key string) bool {
	exist := r.db.Exists(key)
	return exist.Val() > 0
}

func (r *RedisCache) Delete(key string) error {
	err := r.db.Del(key).Err()
	return err
}

func (r *RedisCache) Close() error {
	err := r.db.Close()
	return err
}

func (r *RedisCache) DeleteExpired() {
	//
}
