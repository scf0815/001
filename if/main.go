package main

import "fmt"

//if条件判断
func main() {
	age := 19
	if age >= 18 {
		fmt.Println("你已经成年了")
	} else {
		fmt.Println("你该写暑假作业了")
	}

	//注意作用域问题，此时的age变量只在if语句里面使用有效
	//这样写可以减少内存的应用
	if age := 20; age > 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}
}
