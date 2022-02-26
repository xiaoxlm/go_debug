package pprof

import (
	"runtime"
	"sync"
	"time"
)

var (
	lock sync.Mutex
	items = make(map[int]int)
)

func TestMutex() {
	runtime.SetMutexProfileFraction(1)
	for i:=0; i< 1000*1000*1000; i++ {
		go func(num int) {
			lock.Lock()
			time.Sleep(1 * time.Second)
			items[num] = num
			lock.Unlock()
		}(i)
	}

}
