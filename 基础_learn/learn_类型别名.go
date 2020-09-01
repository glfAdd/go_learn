/*


 */
package main

import "fmt"

func main() {
	//int64的别名为bigin
	type bigint int64
	var a bigint = 100

	type (
		myint    int
		mystring string
	)
	var b myint = 200
	var c mystring = "mystring"
	fmt.Println(a, b, c)

}
