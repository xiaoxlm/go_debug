package pprof

import (
	"fmt"
	"net/http"
)

func TestCPU() {
	fmt.Println("run")
	wasteCPU()
	businessLogic()
	fmt.Println("run done")
}

func businessLogic() {
	http.Get("https://www.baidu.com")
}

func wasteCPU() {
	for i:=0; i<10000; i++ {
		for j:=0; j<10000; j++ {

		}
	}
}
