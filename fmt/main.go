package main

import "fmt"

// %t代表布尔值 %c：字符 %q:字符（串）转成值
// 宽度标识符
// %f:默认宽度，默认精度
// %9f:宽度为9，默认精度
// %.2f：默认宽度，精度为2
// %9.2f：宽度为9，精度为2
// 进行了四舍五入

// go语言中有Scan\Scanf\Scanln获取用户输入
func main() {
	m := make(map[string]int, 10)
	m["china"] = 1
	fmt.Printf("m: %#v\n", m) //#v显示语法
	fmt.Printf("m: %v\n", m)
	n := 99
	fmt.Printf("n: %v%%\n", n) //%%代表加入%，可用于显示百分比
	x := 3.145592653
	fmt.Printf("x: %5f\n", x)
	fmt.Printf("x: %.2f\n", x)
	fmt.Printf("x: %5.2f\n", x)

	var (
		name   string
		age    int
		gender string
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &gender)
	fmt.Println(name, age, gender)
}
