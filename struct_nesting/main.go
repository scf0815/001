package main

import "fmt"

// 匿名字段
// 嵌套结构体

type Person struct {
	name    string
	age     int
	address string
}

type company struct {
	boss Person
	time int
}

func main() {
	p1 := Person{
		name:    "黎明",
		age:     38,
		address: "北京",
	}
	c := company{
		boss: p1,
		time: 18,
	}
	fmt.Printf("%#v\n", c)
	c2 := company{
		boss: Person{
			name:    "黎明",
			age:     38,
			address: "北京",
		},
		time: 100,
	}
	fmt.Printf("c2: %#v\n", c2)
	fmt.Printf("c2.boss.address: %v\n", c2.boss.address)
	// 想要直接访问c2.address,可以定义匿名嵌套结构体

}
