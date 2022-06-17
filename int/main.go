package main

import "fmt"

//整型，查看类型：%T
func main() {
	// 十进制
	var a1 = 10
	fmt.Printf("a1: %d\n", a1) //%d代表十进制,%o代表八进制,%x代表十六进制
	fmt.Printf("a1: %o\n", a1)
	//八进制
	var a2 = 077
	fmt.Printf("a2: %d\n", a2) //0开头是八进制，0x开头是十六进制，从0至f
	//十六进制
	var a3 = 0x1238237823
	fmt.Printf("a3: %d\n", a3)
	fmt.Printf("a3: %b\n", a3) //转换为二进制

	a4 := int8(9)
	fmt.Printf("a4: %v\n", a4)
	fmt.Printf("a4: %T\n", a4)

}
