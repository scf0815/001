package main

import "fmt"

// 分金币作业
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin1() {
	for _, v := range users {
		sum := 0
		for i := 0; i < len(v); i++ {
			switch {
			case v[i] == 'e' || v[i] == 'E':
				sum += 1
			case v[i] == 'i' || v[i] == 'I':
				sum += 2
			case v[i] == 'o' || v[i] == 'O':
				sum += 3
			case v[i] == 'u' || v[i] == 'U':
				sum += 4
			default:
				continue
			}
		}
		distribution[v] = sum
	}
	count := 0
	for i, v := range distribution {
		fmt.Printf("%s:%d\n", i, v)
		count += v
	}
	left := coins - count
	fmt.Println("剩下金币为：", left)
}

// 第二种更为规范简洁
func dispatchCoin2() {
	for _, name := range users {
		for _, c := range name {
			switch c {
			case 'e', 'E':
				distribution[name]++
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			default:
				continue
			}
		}
	}

	for i, v := range distribution {
		fmt.Printf("%s:%d\n", i, v)
	}
	fmt.Printf("剩余金币为: %v\n", coins)

}

func main() {
	// dispatchCoin1()
	dispatchCoin2()

}
