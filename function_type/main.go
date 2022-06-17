package main

import "fmt"

// 函数类型

func f1() int {
	return 5
}

//函数也可以作为参数的类型
func f2(x func() int) int {
	ret := x()
	return ret
}

//函数还可以作为返回值
func ff(a, b int) int {
	return a + b
}
func f3(x func() int) func(int, int) int {
	return ff

}

//匿名函数:调用时需要一个变量接收

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2(f1)
	fmt.Printf("b: %v\n", b)
	c := f3(f1)
	fmt.Printf("c: %T\n", c)
	// 函数内部没有办法声明带名字的函数
	var m = func(x, y int) {
		fmt.Println(x + y)
	}
	m(10, 20)
	// 如果只是调用一次，还可以简写成立即执行函数
	func(x, y int) {
		fmt.Println("x+y=", x+y)

	}(100, 200) //相当于定义完之后就直接调用

}
