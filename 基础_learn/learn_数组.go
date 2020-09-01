/*
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type

设置长度以后没有初始化的值使用默认值
元素个数不能大于设置的size
如果忽略size, 会根据元素数自动设置大小, 确定大小以后设置range意外的值会index out of range
可以通过索引操作数组
长度不可变

*/
package main

import "fmt"

func array1(a1 [6]float32) float32 {
	var a float32 = a1[0] / a1[4]
	return a
}

func main() {
	//没设置的元素使用默认值
	var a = [6]float32{1.1, 2.332}
	//使用索引
	a[4] = 333
	fmt.Println(a)

	//自动设置长度
	var b = []string{"aa", "bb", "cc"}
	fmt.Println(b)

	//声明数组
	var c [10]int
	c[0] = 1

	//多维数组
	var d [2][4]int
	d[0][2] = 3
	fmt.Println(d)

	var e = [2][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}
	fmt.Println(e)

	//数组当作参数
	//未定义长度的数组只能传给不限制数组长度的函数
	//定义了长度的数组只能传给限制了相同数组长度的函数
	var f float32 = array1(a)
	fmt.Println(f)
}
