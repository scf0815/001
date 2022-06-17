package main

import (
	"fmt"
	"strings"
)

// go语言中字符串用双引号包裹，单引号包裹的叫字符
//字节（byte）=8 bit(8个二进制位)是数量单位
//一个字符‘A’占一个字节，一个汉字在utf-8编码中一般占3个字节
//定义多行字符串，利用反引号``(esc下面)
func main() {
	s := "hello world"
	a := 'a'
	b := 10
	fmt.Printf("s: %v\n", s)
	fmt.Printf("s: %T\n", s)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)
	s2 := `
		世情薄
		人情恶
		与送黄昏花易洛
	`
	fmt.Printf("s2: %v\n", s2)

	s3 := "scf"
	s4 := "394903424"
	s5 := s3 + s4
	fmt.Printf("s5: %v\n", s5)
	//分割
	s6 := strings.Split(s5, "9")
	fmt.Printf("s6: %v\n", s6)
	fmt.Printf("s6: %T\n", s6)
	//包含
	fmt.Printf("strings.Contains(s5, \"f\"): %v\n", strings.Contains(s5, "f"))
	//前缀，后缀：判断字符串是不是以某字符开头或结尾
	fmt.Printf("strings.HasPrefix(s5, \"s\"): %v\n", strings.HasPrefix(s5, "s"))
	fmt.Printf("strings.HasSuffix(s5, \"4\"): %v\n", strings.HasSuffix(s5, "4"))

	//判断字符串中某字符出现的位置
	s7 := "abcdefg"
	fmt.Printf("strings.Index(s7, \"e\"): %v\n", strings.Index(s7, "ec"))
	//拼接
	fmt.Printf("strings.Join(s6, \"+\"): %v\n", strings.Join(s6, "+"))
	//字符串修改及类型转换
	c1 := "白萝卜"      //[白 萝 卜]
	c2 := []rune(c1) //把字符串强制转换成一个rune切片
	c2[0] = '红'      //单引号表示修改字符
	fmt.Printf("string(c2): %v\n", string(c2))
	fmt.Printf("c1: %p\n", &c1)
	fmt.Printf("c2: %p\n", &c2)
	c3 := "黄"
	c4 := '黄'
	fmt.Printf("c3: %v\n", c3)
	fmt.Printf("c4: %v\n", c4)
	// c5 := "hello沙河小王子"

}
