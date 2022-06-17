package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 3
	//算数运算符
	fmt.Printf("%v+%v=%v\n", a, b, a+b)
	fmt.Printf("%v-%v=%v\n", a, b, a-b)
	fmt.Printf("%v*%v=%v\n", a, b, a*b)
	fmt.Printf("%v/%v=%v\n", a, b, a/b)
	fmt.Printf("%v %v=%v\n", a, b, a%b)
	a += 1
	fmt.Printf("a: %v\n", a)
	a++
	fmt.Printf("a: %v\n", a)
	//关系运算符
	//go语言是强类型，只有相同类型的变量才能比较
	fmt.Printf("(a == b): %v\n", (a == b))
	fmt.Printf("(a != b): %v\n", (a != b))
	//逻辑运算符：并且（and）：&&，或者(or)：||,not取反:!
	//位运算符：针对的二进制数
	//5的二进制：101,2的二进制：10
	//&：按位与(两位均为1才为1)
	fmt.Println(5 & 2) //000
	//按位或：|（两位有1个1就为1）
	fmt.Println(5 | 2) //7，对应二进制为111
	//按位异或：^(两位不一样就为1)
	//<<:将二进制位左移指定位数
	fmt.Println(5 << 1) //101左移1位是1010，左移几位就是在二进制位后加几个0
	//>>:将二进制位右移指定位数
	fmt.Println(5 >> 1) //101右移1位是10，右移几位就是去掉末尾几位数
}
