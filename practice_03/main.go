package main

import (
	"fmt"
	"os"
)

// 学生管理系统
// 利用结构体获得数据，执行方法
type Student struct {
	id   int
	name string
}

// 创造结构体：一个学生管理者 查看、增加、修改、删除学生
type StudentMgr struct {
	allStudent map[int]Student //map的key用student的id来表示，每个student的id是不重复的
}

var smr StudentMgr

func (s StudentMgr) showStudent() {
	fmt.Println("所有学生信息如下：")
	for _, stu := range s.allStudent {
		fmt.Printf("学号: %v 姓名: %v\n", stu.id, stu.name)
	}
}
func (s StudentMgr) addStudent() {
	var (
		stuID   int
		stuName string
	)
	//获取用户输入
	fmt.Print("请输入学号：") //print来保证不换行
	fmt.Scanln(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	//1.根据用户输入的内容创建一个新的学生
	new := Student{
		id:   stuID,
		name: stuName,
	}
	//2.把新的学生放入map中
	s.allStudent[new.id] = new
	fmt.Println("添加成功！")
}
func (s StudentMgr) modifyStudent() {
	//获取学生的学号
	var stuID int
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	//展示该学号对应的学生信息，如果无，查无此人
	stuName, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你要修改的学生信息为：学号：%v,姓名：%v\n", stuName.id, stuName.name)
	//输入修改后的学生名
	fmt.Print("请输入学生的新名字：")
	var newName string
	fmt.Scanln(&newName)
	//更新学生的姓名
	stuName.name = newName
	// 注意要修改map里面的学生信息
	s.allStudent[stuID] = stuName
	fmt.Println("修改成功！")
}
func (s StudentMgr) deleteStudent() {
	// 1.请用户输入要删除的学生id
	var stuID int
	fmt.Print("请输入要删除的学生学号：")
	fmt.Scanln(&stuID)
	// 2.去map找有无该学生，无就查无此人
	stuName, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	// 3. 有的话删除
	fmt.Printf("你要删除的学生信息为：学号：%v,姓名：%v\n", stuName.id, stuName.name)
	delete(s.allStudent, stuID)
	fmt.Println("删除成功！")
}

// 菜单函数
func showMenu() {
	fmt.Println("Welcome to Student Managent System!")
	fmt.Println(`
		1.查看所有学生
		2.添加学生
		3.修改学生
		4.删除学生
		5.退出
	`)
}

func main() {
	smr = StudentMgr{ //修改的全局变量
		allStudent: make(map[int]Student, 100),
	}
	for {
		showMenu()
		fmt.Print("请输入序号：")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.modifyStudent()
		case 4:
			smr.deleteStudent()
		case 5:
			fmt.Println("您已经退出系统")
			os.Exit(1) //直接退出
		default:
			fmt.Println("非常抱歉，不在业务范围之内")
		}
	}
}
