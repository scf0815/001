package main

import "fmt"

// 接口就像一个协议一样，能实现某种功能的东西
// 同一个结构体实现多个接口
// Sayer 接口
type Sayer interface {
	Say()
}

// Mover 接口
type Mover interface {
	Move()
}
type Dog struct {
	Name string
}

// Dog既可以实现Sayer接口，也可以实现Mover接口。
// 实现Sayer接口
func (d Dog) Say() {
	fmt.Printf("%s会叫汪汪汪\n", d.Name)
}

// 实现Mover接口
func (d Dog) Move() {
	fmt.Printf("%s会动\n", d.Name)
}

// 结构体可以实现嵌套
type dongzuo interface {
	Mover
	Sayer
}

func main() {
	var d = Dog{Name: "旺财"}

	var s Sayer = d
	var m Mover = d

	s.Say()  // 对Sayer类型调用Say方法
	m.Move() // 对Mover类型调用Move方法

}
