package main

import (
	"encoding/json"
	"fmt"
)

// JSON（JavaScript Object Notation）是一种轻量级的数据交换格式
// 1.把go语言中的结构体变量-->json格式的字符串:序列化
// 2.json格式的字符串-->go语言中能够识别的结构体变量：反序列化

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"` //首字母不大写可能无法引用，但有些输出要求是必须小写，反引号可以在指定的包调用修改
}

func main() {
	p1 := Person{
		Name:   "tom",
		Age:    10,
		Gender: "male",
	}
	b, err := json.Marshal(p1) //序列化使用marshal
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	// fmt.Printf("%v\n", string(b))
	fmt.Printf(" %s\n", b)
	// 反序列化
	str := `{"name":"理想","age":18,"gender":"male"}`
	var p2 Person
	json.Unmarshal([]byte(str), &p2) //指针为了修改原来变量
	fmt.Printf("p2: %v\n", p2)

}
