package main

import "fmt"

// 方法：一种作用于特定类型变量函数，这种特定类型变量叫接收者（Receiver）,接收者类似于其它语言中的self
// 定义格式func(接收者变量 接收者类型)方法名（参数列表）（返回参数）{}
// 方法相当于动作，结构体定义的是属性
// 结构体是值类型，复制时不会改变其原本变量，除非复制其内存地址
// 方法里面只取决于接收者是值类型还是指针类型，一般接收者为指针类型
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
}

// 接收者变量一般用类型名的首字母小写例如：p Person
// 什么时候使用指针类型接收者
// 1，需要修改接收者的值
// 2，接收者是拷贝代价比较大的对象
// 3，保持一致性，如果有某个方法使用了指针接收者，那其他的方法也是
func (p Person) sing() {
	fmt.Printf("%v在唱歌\n", p.name)
}

func (p *Person) guonian() {
	p.age++
	fmt.Printf("age: %v\n", p.age)
}
func main() {
	p := Person{
		"tom",
		19,
		"male",
	}
	p1 := newPerson(18, "jerry", "female")
	p1.sing()
	p.sing()
	p1.guonian()

}
