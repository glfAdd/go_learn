/*
长度可变




*/
package main

import (
	"fmt"
)

func main() {
	//创建长度位4的整型切片
	a := []int{1, 2, 3, 4}
	fmt.Println(a)
	fmt.Println(len(a))
	//创建长度为4的数组
	e := [4]int{1, 2, 3, 4}
	e[0] = 100
	fmt.Println(e)

	fmt.Println("-------------------- 通过切片创建的新切片, 底层相同")
	b := a[:]
	i := a[:2]
	i[0] = 33
	fmt.Println(b, i)

	fmt.Println("-------------------- make 创建切片")
	//创建整型且切片, 长度和容量都是5
	c := make([]int, 5)
	fmt.Println(c)
	//长度2, 容量5. 长度不能大于容量
	d := make([]int, 2, 10)
	fmt.Println(d)

	fmt.Println("-------------------- nil切片")
	var f []int

	fmt.Println("-------------------- 空切片")
	g := make([]int, 0)
	h := []int{}
	fmt.Println(f, g, h)

	fmt.Println("-------------------- append")
	j := make([]int, 1, 2)
	//append 返回新的切片
	//在切片的容量小于 1000 个元素时，总是会成倍地增加容量。一旦元素个数超过 1000，容量的增长因子会设为 1.25，也就是会每次增加 25%的容量(随着语言的演化，这种增长算法可能会有所改变)
	//append首先使用可用容量, 没有可用容量时创建新的底层数组
	//切片之间共用一个底层数组
	//创建切片时设置切片的容量和长度一样, 就可以强制让新切片的第一个 append 操作创建新的底层数组，与原有的底层数组分离
	k := append(j, 15)
	m := append(k, 25)
	n := append(m, 35, 45, 55, 65)
	o := append(n, 75, 1, 1, 1, 1, 1)
	//两个切片拼接
	p := append(o, []int{2, 2, 2}...)
	q := append(k[:0], []int{5, 5, 5}...)
	j[0] = 1
	m[0] = 2
	fmt.Println(j, k, m, n, o, p, q)

	fmt.Println("-------------------- 遍历")
	//range返回两个值: 当前迭代到的索引位置, 该位置对应元素值的一份副本
	for index, value := range a {
		fmt.Println(index, value)
		value = 1000
	}
	for index, value := range a {
		fmt.Println(index, value)
		value = 1000
	}
	for index := 1; index < 3; index++ {
		fmt.Println(index, a[index])
	}

	fmt.Println("-------------------- 拷贝")
	//func copy(dst, src []Type) int
	//把切片src中的元素拷贝到切片dst中, 返回值为拷贝成功的元素个数.
	//如果src比dst长就截断, 如果src比dst短则只拷贝src那部分, 其他部分保持不变
	b1 := []string{"a", "b", "c"}
	b2 := []string{"e", "f"}
	b3 := copy(b2, b1)
	fmt.Println(b1, b2, b3)

	b4 := []string{"1", "2", "3", "4", "5"}
	b5 := copy(b4, b2)
	fmt.Println(b4, b5)

	fmt.Println("-------------------- 切片在函数见传递")
	//在 64 位架构的机器上，一个切片需要 24 字节的内存：指针字段需要 8 字节，长度和容量字段分别需要 8 字节
	//由于与切片关联的数据包含在底层数组里，不属于切片本身，所以将切片复制到任意函数的时候，对底层数组大小都不会有影响。复制时只会复制切片本身，不会涉及底层数组
	//地址指针保存的是底层数组的指针

	fmt.Println("-------------------- 删除元素")
	c1 := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Println(c1)
	//使用索引删除元素
	c2 := c1[1:3]
	fmt.Println(c2)

	//删除开头3个元素
	c3 := append(c1[:0], c1[3:]...)
	fmt.Println(c3)

	//删除中间的元素
	c4 := append(c1[:2], c1[4:]...)
	fmt.Println(c4)

	//为什么少了第一个元素? [d e f g h f g h]
	fmt.Println(c1)
}
