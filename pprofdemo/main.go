package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

// 查看哪里分配了大量堆的空间
// go tool pprof http://localhost:39090/debug/pprof/heap
// top --> 可以看到main占用了很多空间
// list main.main 可以看到代码指定位置占用很多空间
func main() {
	arr := make([]string, 0, 100000)
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			// log.Println(1)
			// 分配堆内存
			arr = append(arr, "aaaaaaaa")
		}
	}()
	http.ListenAndServe(":39090", nil)
}
