package main

import "fmt"

//使用switch可以方便对大量的值进行条件判断,起到简化大量判断的作用
//每个switch只能由一个default分支

func switch1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效输入")

	}

}

//case后面也可以跟多个判断变量的取值
func switch2() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Printf("n: %v\n", n)

	}
}

//分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量
func switch3() {
	age := 30
	switch { //该位置的switch后面没有判断变量
	case age < 25:
		fmt.Println("好好学习")
	case age >= 25 && age < 35:
		fmt.Println("好好工作")
	case age >= 35 && age < 60:
		fmt.Println("劳逸结合")
	default:
		fmt.Println("活在当下")

	}
}

//fallthrough语法可以执行满足条件的case的下一个case,是为了兼容C语言的case设计的

func switch4() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")

	}
}

//goto+label可以实现跳出多层for循环
func switch5() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto label
			}
			fmt.Printf("%v-%c\n", i, j)

		}
	}
label:
	fmt.Println("over")
}

func main() {
	switch5()

}
