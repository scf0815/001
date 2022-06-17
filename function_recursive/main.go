package main

// 递归函数：在运行的过程中调用自己
// 永远不要眼高手低
// 计算n 的阶乘
func f(n int) int {
	if n == 1 {
		return 1
	}
	return n * f(n-1)
}

// 斐波那契数列
func ff(n uint) uint {
	if n == 1 || n == 2 {
		return 1
	}
	return ff(n-1) + ff(n-2)
}

//n个台阶，一次可以走1步或者2步，问有多少种走法
func taijie(n uint) uint {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)

}

func main() {
	// a := f(6)
	// fmt.Printf("a: %v\n", a)
	// b := ff(10)
	// fmt.Printf("b: %v\n", b)

}
