package main

import "fmt"

//go支持函数、匿名函数及闭包，函数在go中属于一等公民
// go语言中没有默认参数的概念
// nil代表没有分配内存
func sum(x, y int) (r int) {
	r = x + y
	return
}

//可变长参数，必须放在参数的最后
func f1(x string, y ...int) {
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y) //y的类型是切片

}
func main() {
	r := sum(12, 34)
	fmt.Printf("r: %v\n", r)
	f1("海淀", 1, 3, 4, 7)

}
