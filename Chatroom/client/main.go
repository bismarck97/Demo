package main

import "fmt"

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
			loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
	//根据用户的输入，显示新的提示信息
	if key == 1 {
		//说明用户要登录
		fmt.Println("")
	} else if key == 2 {

	}
}
