package common

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type selectCaseList []reflect.SelectCase

func TestReflectChan(t *testing.T) {
	var sendCh1 = make(chan int)    // channel to use (for send or receive)
	var sendCh2 = make(chan string) // channel to use (for send or receive)

	go func(c chan<- int) {
		for i := 0; i < 8; i++ {
			c <- i
		}
		close(c)
	}(sendCh1)

	go func(c chan<- string) {
		var _string = "abcdefghi"
		for _, ca := range _string {
			c <- string(ca)
		}
		close(c)
	}(sendCh2)

	var selectCase = selectCaseList{}

	fmt.Println(reflect.ValueOf(sendCh1))
	selectCase = append(selectCase, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(sendCh1)})
	selectCase = append(selectCase, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(sendCh2)})

	go func() {
		for {
			chosen, recv, recvOk := reflect.Select(selectCase) // <--- here
			if recvOk {
				fmt.Println(chosen, recv, recvOk)
			}
		}
	}()

	time.Sleep(time.Second * 10)
}
