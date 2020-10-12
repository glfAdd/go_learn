/*
http://maoqide.live/post/golang/golang-data-race-detector/

竞争检测器

在同一个线程中执行多个 goroutine, 有数据安全问题

当两个 goroutine 同时访问同一个变量并且至少有一个是写入时，就会发生 data race(数据竞争)


命令
	$ go test -race mypkg    // to test the package
	$ go run -race mysrc.go  // to run the source file
	$ go build -race mycmd   // to build the command
	$ go install -race mypkg // to install the package


原子函数能够以很底层的加锁机制来同步访问整型变量和指针






*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go printCount()
	go printCount()
	wg.Wait()
}

func printCount() {
	defer wg.Done()
	var i int
	for i = 0; i < 20000; i++ {
		//runtime.Gosched()
		fmt.Println(i)
		i++
	}
}
