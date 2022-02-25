package main

import (
	_"go_debug/example/godebug/init_package1"
	_ "go_debug/example/godebug/init_package2"
	"time"
)

func main() {
	// godebug.Sched()
	//godebug.GC()

	time.Sleep(60 * 10 * time.Second)
}
