package main

import "fmt"

// 接口实现

type animal interface {
	move()
	eat(s string)
}
type cat struct {
	name string
	feet int8
}
type chick struct {
	name string
	feet int8
}

func (c *cat) move() {
	fmt.Printf("%v会上蹦下跳\n", c.name)
}
func (c *cat) eat(s string) {
	fmt.Printf("%v喜欢吃%s\n", c.name, s)
}
func (c *chick) move() {
	fmt.Printf("%v会鸡冻\n", c.name)
}
func (c *chick) eat(s string) {
	fmt.Printf("%v喜欢吃%s\n", c.name, s)
}

// func ani(a animal) {
// 	a.eat("小鱼")
// 	a.move()
// }

func main() {
	c1 := &cat{
		name: "小花",
		feet: 4,
	}
	c2 := &chick{
		name: "大公鸡",
		feet: 2,
	}
	var a1, a2 animal

	a1 = c1
	a1.eat("小鱼")
	a1.move()
	a2 = c2
	a2.eat("饲料")
	a2.move()
	fmt.Printf("a2: %T\n", a2)

}
