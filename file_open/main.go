package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//table键自动补全

func Read1() {
	fileObj, err := os.Open("./main.go") //.表示当前路径
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()
	//读文件
	// var tmp=make([]byte,128)//指定读的长度
	tmp := [128]byte{}
	//循环读
	for {
		n, err := fileObj.Read(tmp[:])
		if err != nil {
			fmt.Println("read file failed", err)
			return
		}
		fmt.Printf("读了多少字节 %v\n", n)
		fmt.Printf("string(tmp[:n]): %v\n", string(tmp[:n]))
		if n < 128 {
			return
		}
	}

}

// 利用bufio读文件,可以实现一行一行的读
func Read2() {
	fileObj, err := os.Open("./main.go") //.表示当前路径
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()
	// 创建一个从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n') //注意是字符,按照一行一行的读
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

// ioutil.ReadFile读取整个文件
// io/ioutil包的ReadFile方法能够读取完整的文件，只需要将文件名作为参数传入。
func Read3() {
	content, err := ioutil.ReadFile("D:/Desktop/123.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

func main() {

	Read3()
}
