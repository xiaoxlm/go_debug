package main

import (
	"go_debug/example/godebug"
	"time"
)

func main() {
	//debug.Sched()
	godebug.GC()

	time.Sleep(60 * 10 * time.Second)
}