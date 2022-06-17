package main

import "fmt"

//数组：存放元素的容器
//必须指定存放的元素类型及容量
//长度是数组类型的一部分
//数组长度为常量，一旦定义，不能修改
func main() {
	var a1 [3]int
	fmt.Printf("a1: %v\n", a1)
	//数组初始化
	a1 = [3]int{1, 2, 3}
	fmt.Printf("a1: %v\n", a1)
	a2 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("a2: %T\n", a2)
	//指定索引
	a3 := [...]int{1: 2, 5: 7}
	fmt.Printf("a3: %v\n", a3)

	//数组遍历
	//方法1
	var b = [...]string{"北京", "上海", "苏州"}
	for i := 0; i < len(b); i++ {
		fmt.Printf("b[%v]: %v\n", i, b[i])
	}
	//方法2:for range
	for index, value := range b {
		fmt.Println(index, value)

	}

	//多维数组
	//go语言中数组元素之间用空格，而非逗号
	//[[1,2] [3,4] [5,6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Printf("a11: %v\n", a11)
	//多维数组的遍历
	for _, v1 := range a11 {
		fmt.Printf("v1: %v\n", v1)
		for _, v2 := range v1 {
			fmt.Printf("v2: %v\n", v2)
		}
	}
	//数组是值类型，赋值和传参会复制整个数组，因此改变副本的值，不会改变本身的值
	//复制过去的值，而非指针，等同于word文件的拷贝粘贴
	c1 := [...]int{1, 2, 3}
	c2 := c1
	c2[0] = 100
	fmt.Printf("c1: %p\n", &c1)
	fmt.Printf("c2: %p\n", &c2)

	d := [...]int{1, 3, 5, 7, 8}
	//求数组中所有元素的和
	sum := 0
	for i := 0; i < len(d); i++ {
		sum += d[i]
	}
	fmt.Printf("sum: %v\n", sum)
	//求数组中两元素和为指定值的下标，比如和为8时，下表为（0，3）和（1，2）
	goal := 8
	for i := 0; i < len(d); i++ {
		for j := i + 1; j < len(d); j++ {
			if d[i]+d[j] == goal {
				fmt.Printf("(%v,%v)\n", i, j)
			}
		}
	}

}
