package main

import (
	"fmt"
)

// map是一种无序的基于key-value的数据结构，是引用类型，必须初始化才能使用
//一般用make进行初始化，开辟内存空间
func main() {
	var m1 = make(map[string]int, 10) //估算好map容量，避免在程序运行期间再动态扩容
	m1["age"] = 18
	m1["gender"] = 0
	fmt.Printf("m1: %v\n", m1)
	//判断其中的key是否存在
	value, ok := m1["age"]
	//约定俗成用ok接收bool值
	if !ok {
		fmt.Println("查无此人")
	} else {
		fmt.Printf("value: %v\n", value)
	}

	// 删除键值对：delete(map,key)
	delete(m1, "gender")
	fmt.Printf("m1: %v\n", m1)

	//元素类型为map的切片
	var s1 = make([]map[int]string, 10) //切片初始化
	s1[0] = make(map[int]string, 1)
	s1[0][0] = "沙河"
	s1[0][1] = "海淀"
	fmt.Printf("s1: %v\n", s1)

	//值为切片类型的map
	var y1 = make(map[string][]string, 10)
	y1["北京"] = []string{"海淀", "朝阳"}
	fmt.Printf("y1: %v\n", y1)

}
