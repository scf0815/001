package main

import "fmt"

//go语言中只有for循环
func main() {

	//九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%v×%v=%v\t", i, j, i*j)
		}
		fmt.Println("")
	}
	//变种1
	i := 5
	for ; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
	//变种2
	i = 5
	for i < 10 {
		fmt.Printf("i: %v\n", i)
		i++
	}
	//for range(键值循环)：遍历数组、切片、字符串、map以及channel
	//数组、切片、字符串返回索引和值;map返回键和值;channel只返回channel里面的值
	s := "hello,北航"
	for i, j := range s {
		fmt.Printf("i: %v\tj: %c\n", i, j) //%c代表输出单个字符
	} //一个汉字一般占3个字节，故其索引也是占3位

	//流程控制之跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break //跳出for循环，剩下的循环均不再执行
			//continue 跳出此次循环，执行下一次循环
		}
		fmt.Printf("i=%v\n", i)
	}
	fmt.Println("over")

}
