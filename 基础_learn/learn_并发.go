/*
Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。
goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
goroutine 语法格式：
	go 函数名( 参数列表 )

同一个程序中的所有 goroutine 共享同一个地址空间



*/
package main

import (
	"fmt"
	"time"
)

func testThread(a int, b int) {
	c := a + b
	fmt.Println("sleep begin")
	time.Sleep(3)
	fmt.Println(c)

}
func main() {
	go testThread(1, 5)
	testThread(2, 500)
}
