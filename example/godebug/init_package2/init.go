package init_package2

import (
	"time"
)

func init() {
	println("init2")
	slice := make([]int, 8)
	for i := 0; i < 32*1000*1000; i++ {
		slice = append(slice, i)
	}
	time.Sleep(2 * time.Second)
}
