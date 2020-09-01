/*
方法: 定义一个结构体,将函数和结构体绑定在一起的东西就是
Go语言中的函数可以和任何类型绑定, 但是一般用于和结构体绑定

格式:
func (接收者名称 接收者类型)函数名称(形参列表)(返回值列表){
	逻辑语句;
}


注意点:
1.方法和函数的区别在于, 函数可以直接调用(通过包名.函数名称), 而方法只能通过绑定的数据类型对应的变量来调用(变量.函数名称)
2.函数名称和方法名称可以重名


*/
package main

import "fmt"

type Animal struct {
	name string
	age  int
}

//地址传递
//指定接收者名称
func (a *Animal) run() {
	a.name = "Lucy"
	a.age = 221
	fmt.Println("run1")
	fmt.Println(a.name, a.age)
}

//值传递
//指定接收者名称
func (a Animal) run3() {
	fmt.Println("run3")
	a.name = "小明"
	a.age = 1000
}

//没有执行接收者名称
func (Animal) run2() {
	fmt.Println("run2")
}

func main() {
	a := Animal{"Tom", 22}
	//如果指定了接收者名称,那么调用方法时会将调用者传递给接收者(可以把接收者看做函数的形参)
	fmt.Println(a)
	a.run()
	fmt.Println(a)
	a.run2()
	fmt.Println(a)
	a.run3()
	fmt.Println(a)
}
