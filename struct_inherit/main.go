package main

import "fmt"

// go语言使用结构体可以实现其他编程语言中面向对象的继承
// 可以继承属性以及方法

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%v会动\n", a.name)
}

type Dog struct {
	feet int8
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%v汪汪叫\n", d.name)
}

func main() {
	d1 := &Dog{
		feet: 4,
		Animal: &Animal{
			name: "花花",
		},
	}
	fmt.Printf("d1: %v\n", *d1)
	fmt.Printf("d1.Animal: %v\n", *d1.Animal)
	d1.move()
	d1.wang()

}
