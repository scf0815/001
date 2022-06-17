package main

import (
	"fmt"
	"sort"
)

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1
	var a3 = make([]int, 3)
	copy(a3, a1)
	fmt.Printf("a1: %v,a2: %v,a3: %v\n", a1, a2, a3)
	a1[0] = 100 //copy等同于将word里面的复制文字
	fmt.Printf("a1: %v,a2: %v,a3: %v\n", a1, a2, a3)
	//要从切片a中删除索引为index的元素，操作方法是：a=append(a[:index],a[index+1:]...)
	a1 = append(a1[:1], a1[2:]...)
	fmt.Printf("a1: %v\n", a1)

	//切片不保存具体的值
	//底层数组都是占用一块连续的内存

	x1 := [...]int{1, 3, 4}
	s1 := x1[:]
	fmt.Println(s1, len(s1), cap(s1))
	s1 = append(x1[:1], x1[2:]...)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Printf("x1: %v\n", x1) //[1,4,4]相当与把第二个值变成4

	var a = make([]int, 5, 10)
	fmt.Printf("a: %v\n", a) //注意只要定义了长度，就一定存在相应数量的元素存在
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Printf("a: %v\n", a)
	fmt.Printf("cap(a): %v\n", cap(a))
	sort.Ints(a) //内置函数进行排序
	fmt.Printf("a: %v\n", a)

	c1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15}
	d1 := c1[:]
	//去掉切片里面的5
	d1 = append(d1[:2], d1[3:]...)
	fmt.Printf("d1: %v\n", d1)
	fmt.Printf("c1: %v\n", c1)

}
