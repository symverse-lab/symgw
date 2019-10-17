package common

import (
	"reflect"
	"sync"
)

type CacheSaveEvent struct{ key []byte }

type caseList []reflect.SelectCase

type Feed struct {
	once sync.Once

	mu        sync.Mutex
	sendCases caseList
	removeSub chan interface{}
}

func (f *Feed) init() {
	f.removeSub = make(chan interface{})
	//f.sendLock = make(chan struct{}, 1)
	//f.sendLock <- struct{}{}
	f.sendCases = caseList{{Chan: reflect.ValueOf(f.removeSub), Dir: reflect.SelectRecv}}
}

func (f *Feed) Subscribe(inputCh interface{}) {

}
