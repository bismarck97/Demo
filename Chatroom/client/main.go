package main

import (
	"fmt"
	"os"
)

//定义两个变量，一个表示用户id 一个表示用户密码
var userId int
var userPwd string

func main() {
	//接收用户选择
	var key int
	//判断是否还继续显示菜单
	var loop = true
	for loop {
		fmt.Println("---------------欢迎登录多人聊天系统---------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("请选择(1-3)")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			//loop = false
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
	//根据用户的输入，显示新的提示信息
	if key == 1 {
		//说明用户要登录
		fmt.Println("请输入用户的ID")
		fmt.Scanln(&userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanln(&userPwd)
		//先把登录的函数，写到另外一个文件，比如login.go
		login(userId, userPwd)
		//if err != nil {
		//	fmt.Println("登录失败")
		//} else {
		//	fmt.Println("登录成功")
		//}
	} else if key == 2 {
		fmt.Println("进行用户注册逻辑")
	}
}
