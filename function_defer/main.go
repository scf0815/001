package main

import "fmt"

// 资源清理 文件关闭 解锁以及记录时间等
// go语言中的return不是原子操作，在底层分为两步来执行
// 1 返回值赋值
// 2 真正的RET返回
// 函数中如何存在defer,那么defer在1,2之间执行

func f1() int {
	x := 5
	defer func() {
		x++ //修改的x不是返回值
	}() //立即执行函数
	return x //返回值为5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //返回值为6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //返回值为5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 //返回值为5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {
	// r := f1()
	// fmt.Printf("r: %v\n", r)
	// r = f2()
	// fmt.Printf("r: %v\n", r)
	// r = f3()
	// fmt.Printf("r: %v\n", r)
	// r = f4()
	// fmt.Printf("r: %v\n", r)
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1

}
