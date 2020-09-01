/*
接口内只实现了函数的声明,没有函数实现

Go语言中实现一个接口, 不需要做额外的声明, 只要某种数据类型绑定了所有接口中的方法就是实现了这个接口
实现接口之后的特性:
	只要某种数据类型实现了接口, 那么就可以使用接口变量保存这种数据类型
	只要某种数据类型实现了接口, 那么就可以使用接口变量调用接口中声明的方法


接口声明的格式
type 接口类型名 interface{
	方法名1( 参数列表1 ) 返回值列表1
	方法名2( 参数列表2 ) 返回值列表2
}
1. 接口类型名：使用 type 将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加 er，如有写操作的接口叫 Writer，有字符串功能的接口叫 Stringer，有关闭功能的接口叫 Closer 等。
2. 方法名：当方法名首字母是大写时，且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
3. 参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以被忽略

例如
type Writer interface {
	//方法名 Write, 请求参数p byte类型数组, 返回值err error类型
	Write(p []byte) (n int, err error)
}


1.接口中只能声明方法, 不能实现方法, 方法不能同名
2.不能声明变量
3.只有实现了接口中声明的所有方法, 才算实现了接口, 才能使用接口变量保存
4.在实现接口的时候, 方法名称,形参列表,返回值列表必须一模一样
4.在实现接口的时候, 可以通过接口变量来保存这个结构体变量
5.当这个接口变量保存了这个结构体变量,那么可以通过调用接口的方法来实现结构体调用这个方法. 但是接口变量不能调用接口中没有的方法,也不能调用结构体的属性,要想调用接口以外的方法和变量,我们只能通过接口类型转换来实现
5.接口和结构体一样, 可以嵌套
6.接口和结构体一样, 嵌套时不能嵌套自己(自己搞自己)
7.可以将集接口变量赋值给子集接口变量,不可以将子集接口变量赋值给超集接口变量(无论实际的数据类型是否已经实现了超集的所有方法)






*/
package main

import "fmt"

type Usb interface {
	open()
	close()
}

type Phone struct {
	name string
}

func (p Phone) open() {
	fmt.Println("open", p.name)
}
func (p Phone) close() {
	fmt.Println("close", p.name)
}

type Computer struct {
	name string
}

func (c Computer) open() {
	fmt.Println("open", c.name)
}

func (c Computer) close() {
	fmt.Println("close", c.name)
}

func Call(u Usb) {
	u.open()
	u.close()
}

func main() {
	var p Phone
	p.name = "手机"
	p.open()
	p.close()
	Call(p)
	fmt.Println("--------------------")
	c := Computer{"电脑"}
	c.open()
	c.close()
	Call(c)

	fmt.Println("-------------------- 空接口")
	//空接口可以保存任何类型的数据
	//数组和字典一般是用来保存相同类型数据的,但是我们可以利用空接口使它们保存不同的类型的数据
	var arr [5]interface{}
	arr[0] = 1
	arr[2] = "aa"
	arr[3] = false
	fmt.Println(arr)

	fmt.Println("-------------------- 接口类型转换方式1")
	//value,ok := 接口变量名称.(具体数据类型)
	//将接口转为指定的类型, 并返回两个值
	//value为转换后的结构体或接口
	a1 := Phone{"手机2"}
	var b1 Usb = a1
	//x: 转换后的类型
	//y: 转换的类型和接口保存的类型不一样时ok为false
	x, y := b1.(Phone)
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("-------------------- 接口类型转换方式2")
	//https://www.jianshu.com/p/3735fe9e3f25
	//没看懂
	//a2 := Phone{"手机 a2"}
	//switch expr {
	//
	//}

	fmt.Println("-------------------- 接口类型转换方式3")
	//定义空接口保存结构体
	var a3 interface{} = Phone{"手机a3"}
	//把空接口转为具体接口
	if x, y := a3.(Usb); y {
		fmt.Println(x)
	}
	fmt.Println("-------------------- 接口类型转换方式4")
	//定义空接口保存结构体
	var a4 interface{} = Phone{"手机a4"}
	//把空接口转为结构体
	if x, y := a4.(Phone); y {
		fmt.Println(x)
	}
}
