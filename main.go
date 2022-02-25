package main

import (
	"go_debug/example/godebug"
	"time"
)

func main() {
	godebug.Sched()
	//godebug.GC()
	time.Sleep(10 * time.Second)
}
