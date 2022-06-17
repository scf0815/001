package main

import "fmt"

// 批量声明常量时，如果有一行没有赋值，那就和上一行的值保持一致
const (
	n1 = 100
	n2
	n3
)

const (
	a1 = iota
	a2
	a3
)

// 插队
const (
	b1 = iota
	b2 = 100
	b3 = iota
	b4
)

// 多个常量在同一行
const (
	c1, c2 = iota + 1, iota + 2
	c3, c4 = iota + 2, iota + 3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	// 代表1左移10位，10000000000二进制=1024十进制
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Printf("KB: %v\n", KB)
	fmt.Printf("MB: %v\n", MB)
	fmt.Printf("GB: %v\n", GB)
	fmt.Printf("TB: %v\n", TB)
	fmt.Printf("PB: %v\n", PB)

}
