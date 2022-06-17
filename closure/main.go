package main

import "fmt"

// 闭包
// 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单地说，闭包=函数+引用环境
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

func f3(f func(int, int), x, y int) func() {
	temp := func() {
		f(x, y)
	}
	return temp
}

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// 函数分析不要慌:关键字+函数名+参数+返回值
// 匿名函数中修改了base的值
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

//变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。在f的生命周期内，变量x也一直有效
// 底层原理
// 1 函数可以作为返回值
// 2 函数内部查找变量的顺序，先在自己内部找，找不到往外部找
// 闭包=函数+外部变量的引用
func main() {

	// f := adder(10) //x=10,f相当于adder的返回值，也是一个函数
	// fmt.Printf("f: %T\n", f)
	// fmt.Println(f(10)) //20
	// fmt.Println(f(20)) //40
	// fmt.Println(f(30)) //70
	// ret := f3(f2, 100, 200)
	// fmt.Printf("ret: %T\n", ret)
	// f1(ret)
	// 原因：想把f2传入到f1中，但f1的参数！=f2的返回值，没法直接传递，需要构造一个参数，里面匿名函数返回类型是f1的参数类型

	x1, x2 := calc(10)
	fmt.Println(x1(1), x2(2))
	fmt.Println(x1(3), x2(4))
	fmt.Println(x1(5), x2(6))

}
