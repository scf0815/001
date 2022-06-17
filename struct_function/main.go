package main

import "fmt"

// 构造函数：约定俗成用new开头
type Person struct {
	name   string
	age    int
	gender string
}

func newPerson(age int, name, gender string) *Person {
	return &Person{
		name:   name,
		age:    age,
		gender: gender,
	}
	// 构造函数返回值是结构体还是结构体指针时需要考虑，尽量使用指针类型，不消耗太多内存
}
func main() {
	p1 := newPerson(18, "tom", "male")
	fmt.Printf("p1: %v\n", *p1)
	p2 := newPerson(20, "jerry", "female")
	fmt.Printf("p2: %v\n", *p2)

}
