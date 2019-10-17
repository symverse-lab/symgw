package db

import (
	"time"
)

type NonCache struct{}

func NewNonCache() Cache {
	cache = &NonCache{}
	return cache
}

func (non *NonCache) Write(key string, value []byte, d time.Duration) error {
	return nil
}

func (non *NonCache) Get(key string) ([]byte, error) {
	return nil, nil
}

func (non *NonCache) IsCached(key string) bool {
	return false
}

func (non *NonCache) Delete(key string) error {
	return nil
}

func (non *NonCache) Close() error {
	return nil
}

func (non *NonCache) DeleteExpired() {

}
