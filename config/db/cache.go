package db

import (
	"time"
)

var cache Cache

type Cache interface {
	Write(key string, value []byte, duration time.Duration) error
	Get(key string) ([]byte, error)
	IsCached(key string) bool
	Delete(key string) error
	Close() error
	DeleteExpired()
}

func GetCache() Cache {
	return cache
}

const (
	DB_LEVEL = "leveldb"
	DB_REDIS = "redis"
)
