package db

import (
	"github.com/syndtr/goleveldb/leveldb"
	"sync"
	"time"
)

type CacheTime struct {
	cachedTime time.Time
	duration   time.Duration
}

type LevelCache struct {
	db        *leveldb.DB
	timeTable map[string]CacheTime
	mu        sync.RWMutex
}

func NewLevelDbCache(filepath string) (Cache, error) {
	db, err := leveldb.OpenFile(filepath, nil)
	if err != nil {
		return nil, err
	}
	cache = &LevelCache{
		db:        db,
		timeTable: make(map[string]CacheTime),
	}
	return cache, nil
}

func (level *LevelCache) Write(key string, value []byte, d time.Duration) error {

	err := level.db.Put([]byte(key), value, nil)
	level.timeTable[key] = CacheTime{
		time.Now(),
		d,
	}
	return err
}

func (level *LevelCache) Get(key string) ([]byte, error) {
	level.mu.RLock()
	defer level.mu.RUnlock()

	_, isCached := level.timeTable[key]
	diff := level.timeTable[key].cachedTime.Add(level.timeTable[key].duration).Sub(time.Now()).Nanoseconds()
	if isCached && diff < 0 {
		delete(level.timeTable, key)
		_ = level.Delete(key)
	}
	data, err := level.db.Get([]byte(key), nil)
	return data, err
}

func (level *LevelCache) IsCached(key string) bool {
	level.mu.RLock()
	defer level.mu.RUnlock()

	_, isCached := level.timeTable[key]
	diff := level.timeTable[key].cachedTime.Add(level.timeTable[key].duration).Sub(time.Now()).Nanoseconds()
	if isCached && diff < 0 {
		delete(level.timeTable, key)
		_ = level.Delete(key)
	}

	_, isCached = level.timeTable[key]
	return isCached
}

func (level *LevelCache) Delete(key string) error {
	err := level.db.Delete([]byte(key), nil)
	return err
}

func (level *LevelCache) Close() error {
	err := level.db.Close()
	return err
}

func (level *LevelCache) DeleteExpired() {
	level.mu.Lock()
	defer level.mu.Unlock()

	for key, cacheTime := range level.timeTable {
		if cacheTime.cachedTime.Add(cacheTime.duration).Sub(time.Now()).Nanoseconds() < 0 {
			delete(level.timeTable, key)
			_ = level.Delete(key)
		}
	}
}
