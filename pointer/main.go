package main

import "fmt"

//pointer
//go语言中不存在指针操作
//&：取地址   *:根据地址取值
//内存地址是利用16进制的数保存的
func main() {
	a := 8
	b := &a
	fmt.Printf("a: %p\n", &a)
	//b存的是a的内存地址值，但他本身也有自己的内存地址
	fmt.Printf("b: %p\n", &b)

	//new内置函数申请一个内存地址,int默认申请0的内存地址
	var c = new(int)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("*c: %v\n", *c)
	*c = 100
	fmt.Printf("*c: %v\n", *c)
	fmt.Printf("c: %v\n", c)
	//make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，
	//而且返回类型是他们类型的本身，不是他们的指针类型，因为他们是引用类型。

}
