package main

import (
	"fmt"
	"math"
)

//默认go语言中的浮点数都是float64类型
// 布尔值无法参与数值运算，也无法与其他数据类型进行转换
func main() {
	fmt.Printf("math.Pi: %v\n", math.Pi)
	fmt.Printf("math.MaxFloat32: %v\n", math.MaxFloat32)
	fmt.Printf("math.MaxFloat64: %v\n", math.MaxFloat64)
	f := float32(2.45)
	fmt.Printf("f: %T\n", f)
}
