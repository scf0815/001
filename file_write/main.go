package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 打开文件写内容
func Write1() {
	file, err := os.OpenFile("D:/Desktop/123.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666) //没有文件可以创建
	if err != nil {
		fmt.Println("failed", err)
	}
	defer file.Close()
	str := "hello 沙河"
	file.Write([]byte(str)) //写入字节切片数据
	file.WriteString("hello 小王子")
}

func Write2() {
	file, err := os.OpenFile("D:/Desktop/123.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666) //0666代表文件权限
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello沙河\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}

func Write3() {
	str := "hello 沙河"
	err := ioutil.WriteFile("D:/Desktop/123.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func main() {
	Write3()
}
