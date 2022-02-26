package pprof

import (
	"fmt"
	"time"
)

type Obj struct {
	Name string
	Arr []int
}

func Alloc() {
	o := allocObject()
	time.Sleep(500 * time.Millisecond)
	Alloc()
	fmt.Println(o)
}

func allocObject() *Obj {
	slice := make([]int, 8)

	for i := 0; i < 10*200*200; i++ {
		slice = append(slice, i)
	}

	return &Obj{
		Arr: slice,
	}
}

func UseObj(o *Obj) {
	o.Name = "xxx"
}