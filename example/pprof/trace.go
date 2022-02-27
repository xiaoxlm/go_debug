package pprof

import (
	"fmt"
	"runtime"
)

func TestTrace(ch chan int){
	go JustPrint()

	go Action(ch)
}

func JustPrint() {
	fmt.Println("!!!")
}

func Action(ch chan int) {
	slice := make([]int, 8)

	for i := 0; i < 10*200*200; i++ {
		slice = append(slice, i)
	}
	runtime.GC()
	ch <- 1
}