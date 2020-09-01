package main

import "fmt"

type USB interface {
	start()
	end()
}
type Computer struct {
	name string
}

func (c Computer) start() {
	fmt.Println(c.name, "被打开了")
}
func (c Computer) end() {
	fmt.Println(c.name, "被关闭了")
}

type Phone struct {
	name string
}

func (ph Phone) start() {
	fmt.Println(ph.name, "被打开了")
}
func (ph Phone) end() {
	fmt.Println(ph.name, "被关闭了")
}

func main() {
	var in USB
	var cm Computer = Computer{"三星"}
	Option(cm)
	var ph Phone = Phone{"苹果"}
	Option(ph)
}
func Option(in USB) {
	in.start()
	in.end()
}
