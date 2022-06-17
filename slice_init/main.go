package main

import "fmt"

//切片本质上就是可变长度的数组
//切片长度：len(),容量：cap()
func main() {
	var s1 = []int{}
	s1 = []int{1, 2, 3, 5, 7, 9, 10, 11, 13, 15}
	fmt.Printf("len(s1):%v,cap(s1):%v\n", len(s1), cap(s1))
	s2 := s1[4:]
	fmt.Printf("len(s2):%v,cap(s2):%v\n", len(s2), cap(s2))
	s3 := s1[:4]
	fmt.Printf("len(s3):%v,cap(s3):%v\n", len(s3), cap(s3))
	s4 := s1[4:7]
	fmt.Printf("len(s4):%v,cap(s4):%v\n", len(s4), cap(s4))
	//切片的容量：底层数组从切片的第一个元素到底层数组最后一个元素的数量
	s5 := s1
	s5[0] = 100
	fmt.Printf("s5: %v\n", s5)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s1: %p\n", &s1)
	fmt.Printf("s5: %p\n", &s5)

	//切片是引用类型，均指向了底层的数组

}
