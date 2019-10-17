package types

import (
	"sync"
)

var (
	once sync.Once
)

func GetInstance(f func(), instance interface{}) interface{} {
	once.Do(f)
	return instance
}
