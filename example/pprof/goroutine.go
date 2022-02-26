package pprof

import (
	"net/http"
	"time"
)

func Gorountine() {
	for{
		go action()
		time.Sleep(100 * time.Millisecond)
	}
}

func action() {
	for {
		http.Get("https://www.baidu.com")
		time.Sleep(1 * time.Second)
	}


}