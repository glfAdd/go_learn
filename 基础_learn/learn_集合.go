/*
1. Map是一种无序的键值对的集合, 通过key来快速检索数据
2. 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对

声明变量，默认 map 是 nil
var map_variable map[key_data_type]value_data_type

使用 make 函数
map_variable := make(map[key_data_type]value_data_type)


*/
package main

import "fmt"

func main() {
	//创建集合
	var a map[string]string
	a = make(map[string]string)
	a["name"] = "Tom"
	a["tel"] = "010-2233298"
	fmt.Println(a["name"], a["tel"])

	//判断元素是否存在, 如果没有找到b为空字符串, c为false
	b, c := a["haha"]
	fmt.Println(b, c)

	//根据key删除元素
	delete(a, "name")
	fmt.Println(a)

	var d = make(map[string]string)
	d["aa"] = "aa"
	fmt.Println(d)
}
