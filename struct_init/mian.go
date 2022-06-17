package main

import "fmt"

// go语言中利用struct来实现面向对象
// 自定义类型和类型别名
// type后面跟的是类型
// 声明字符：rune本质上是int32，官方自定义的内置函数
type myInt int     //自定义类型
type yourInt = int //类型别名

var m myInt
var n yourInt
var x rune

type Person struct {
	name   string
	age    int
	gender string
}

func f1(x Person) {
	x.age = 20
	fmt.Printf("x: %v\n", x)
}
func f2(x *Person) {
	x.age = 100               //语法糖，自动根据指针找到对应的变量
	fmt.Printf("x: %v\n", *x) //根据内存地址可以找到原来的变量，这样可以修改原来的变量

}

func main() {
	var p Person
	p.age = 18
	f1(p)
	fmt.Printf("p: %v\n", p)
	f2(&p)
	fmt.Printf("p: %v\n", p)

	// m = 100
	// fmt.Printf("m: %T\n", m)
	// n = 100
	// fmt.Printf("n: %T\n", n)
	// x = '中'
	// fmt.Printf("x: %T\n", x)
	// fmt.Printf("x: %c\n", x)
	// zhouyun := Person{
	// 	"zhoulin",
	// 	18,
	// 	[]string{"hejiu", "tangtou"},
	// 	"男",
	// }
	// fmt.Printf("zhouyun: %v\n", zhouyun)
	// 匿名结构体 多用于临时场景
	// var s struct {
	// 	x string
	// 	y int
	// }
	// s.x = "嘿嘿嘿"
	// s.y = 100
	// fmt.Printf("s: %v\n", s)
	//结构体是值类型
	// go语言中函数传参永远是拷贝,修改的也只是副本

}
