package main

import (
	"go_debug/example/pprof"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	/***GODEBUG***/
	// godebug.Sched()
	// godebug.GC()
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

	{
		// mutex
		go pprof.TestMutex()
	}

	http.ListenAndServe(":8989", nil)
	/***pprof***/
}
