package godebug

import (
	"fmt"
	"sync"
)

func Sched() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			var counter int
			for i := 0; i < 1e10; i++ {
				counter++
			}
			wg.Done()
		}(&wg)
	}
	wg.Wait()
	fmt.Println("schedule done")
}