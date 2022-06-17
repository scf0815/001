package main

import "fmt"

// 创建指针类型结构体
// 可以通过new关键字对结构体进行实例化，得到的是结构体的地址
type Person struct {
	name   string
	age    int
	gender string
}

func main() {
	var p = new(Person)
	fmt.Printf("p: %T\n", p)
	fmt.Printf("p: %#v\n", p)
	fmt.Printf("p: %p\n", p)
	fmt.Printf("p: %v\n", p)  //p的值就是一个内存地址，同时他还有自己的内存
	fmt.Printf("p: %p\n", &p) //相当于取p的地址，p是指针类型，也有自己的地址
	// go语言支持对结构体指针直接使用.来访问结构体的成员
	// 结构体初始化1
	p.age = 18
	p.gender = "male"
	p.name = "tom"
	fmt.Printf("p: %v\n", p)
	// 结构体初始化2(key:value)
	p1 := Person{
		name:   "jerry",
		age:    12,
		gender: "female",
	}
	fmt.Printf("p1: %#v\n", p1)
	// 结构体初始化3(value,省去key)，注意写的顺序
	p2 := Person{
		"lamda",
		10,
		"male",
	}
	fmt.Printf("p2: %v\n", p2)
	// 一个结构体占用的内存是连续的
	fmt.Printf("p2.age: %p\n", &p2.age)
	fmt.Printf("p2.gender: %p\n", &p2.gender)
	fmt.Printf("p2.name: %p\n", &p2.name)

}
