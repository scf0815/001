package main

import "fmt"

// 空接口是指没有定义任何方法的接口类型。
// 因此任何类型都可以视为实现了空接口。
// 也正是因为空接口类型的这个特性，空接口类型的变量可以存储任意类型的值。
// 通常我们在使用空接口类型时不必使用type关键字声明，可以像下面的代码一样直接使用interface{}。
// var x interface{} // 声明一个空接口类型变量x
// 使用空接口实现可以接收任意类型的函数参数
// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

// 类型断言
func assign(a interface{}) {
	// str, ok := a.(string)
	// if !ok {
	// 	fmt.Println("类型断言错误")
	// } else {
	// 	fmt.Printf("str: %T\n", str)
	// }
	switch t := a.(type) {
	case string:
		fmt.Printf("type:%T value: %v\n", t, a)
	case int:
		fmt.Printf("type:%T value: %v\n", t, a)
	case bool:
		fmt.Printf("type:%T value: %v\n", t, a)
	}

}
func main() {

	// 使用空接口实现可以保存任意值的字典。
	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	studentInfo["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(studentInfo)
	fmt.Printf("studentInfo: %T\n", studentInfo)
	var x interface{}
	x = "helloworld"
	show(x)
	show(studentInfo)

	assign("100")
	assign("hello")

}
