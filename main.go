package main

import (
	"go_debug/example/pprof"
	_ "net/http/pprof"
	"os"
	"runtime/trace"
)

func main() {
	/***GODEBUG***/
	// godebug.Sched() // schedtrace + scheddetail
	// godebug.GC() // gctrace
	// time.Sleep(10 * time.Second)
	/***GODEBUG***/

	/***pprof***/
	//{
	//	// cpu
	//	go func() {
	//		for {
	//			pprof.TestCPU()
	//			time.Sleep(500 * time.Millisecond)
	//		}
	//	}()
	//}
	//{
	//	// memory
	//	go pprof.TestAlloc()
	//}
	//{
	//	// gorountine
	//	go pprof.TestGorountine()
	//}

	//{
	//	// mutex
	//	go pprof.TestMutex()
	//}

	//http.ListenAndServe(":8989", nil)
	/***pprof***/

	/***trace***/
	{
		trace.Start(os.Stderr)
		defer trace.Stop()

		ch := make(chan int)
		go pprof.TestTrace(ch)
		<- ch
	}
	/***trace***/
}
