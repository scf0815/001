package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 计算一个字符串中汉字的数量
// 1 依次拿到字符串中的字符
// 2 判断每个字符是否为汉字
// 3 把汉字出现的次数累加得到最终结果
func main() {
	s1 := "helloworld北京欢迎你"
	sum := 0
	for _, v := range s1 {
		if unicode.Is(unicode.Han, v) {
			sum += 1
		}
	}
	fmt.Println("汉字数量为：", sum)

	// how do you do单词出现的次数
	s2 := "how do you do"
	s3 := strings.Split(s2, " ")
	fmt.Printf("s3: %v\n", s3)
	m1 := make(map[string]int, 10)
	fmt.Printf("m1: %v\n", m1)
	for _, v := range s3 {
		_, ok := m1[v]
		if !ok {
			m1[v] = 1
		} else {
			m1[v]++
		}

	}
	fmt.Printf("m1: %v\n", m1)
	// 回文判断：一个字符串从左向右读和从右向左读是一样的，比如说：上海在海上

	s4 := "上海自来水来自海上"
	// 由于1个汉字占3个字节，需要把字符串里面的汉字拿出来放到一个rune[]里面
	r := make([]rune, 0, len(s4))
	for _, v := range s4 {
		r = append(r, v)
	}
	fmt.Printf("len(r): %v\n", len(r))
	for i := 0; i < len(r); i++ {
		if r[i] != r[len(r)-i-1] {
			fmt.Println("不是回文")
			break
		} else {
			continue
		}

	}
	fmt.Println("是回文")

}
