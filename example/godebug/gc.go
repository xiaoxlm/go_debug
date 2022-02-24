package godebug

import (
	"fmt"
	"runtime"
	"time"
)

func GC() {
	testGC()
	fmt.Println("force gc")
	runtime.GC()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("force gc done")
}

func testGC() {
	slice := make([]int, 8)

	fmt.Println("slice alloc begin.")
	for i := 0; i < 32*1000*1000; i++ {
		slice = append(slice, i)
	}
	fmt.Println("slice alloc end.")
	fmt.Println("=================")
}