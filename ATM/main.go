package main

import "fmt"

func main() {
	s := NewAccount()
	welcome(s)
}

//欢迎界面
func welcome(s *accountList) {
	var command int
	var user *account
	for {
		fmt.Println("========欢迎进入银行ATM系统========")
		fmt.Println("1.登录账号")
		fmt.Println("2.注册开户")
		fmt.Println("3.退出系统")
		fmt.Println("请您选择操作:")
		fmt.Scan(&command)
		switch command {
		case 1:
			//用户登录操作
			user = s.login()
			if user != nil {
				operate(user, s)
			} else {
				welcome(s)
			}

		case 2:
			//用户注册操作
			s.signIn()
		case 3:
			//退出系统
			goto breakWelcome
		default:
			fmt.Println("输入有误，请重新输入")

		}
	}
breakWelcome:
	fmt.Println("感谢您使用银行ATM系统")
}

//用户操作界面
func operate(account *account, userList *accountList) {
	//单用户操作
	var user User
	user = account
	var users Users
	//系统操作
	users = userList
	var command int
	for {
		fmt.Println("========欢迎您进入银行用户操作界面========")
		fmt.Println("1.查询\t2.存款:\t3.取款\t4.转账")
		fmt.Println("5.修改密码\t6.退出\t7.注销当前账号")
		fmt.Println("请您选择操作:")
		fmt.Scan(&command)
		switch command {
		case 1:
			user.Select()
		case 2:
			user.SetMoney()
		case 3:
			user.withdrawingMoney()
		case 4:
			users.transfer(account)
		case 5:
			user.changePassword(userList)
		case 6:
			goto welcome
		case 7:
			users.deleteUser(account)
		}
	}
welcome:
	fmt.Println("欢迎再次使用银行用户操作系统")
	welcome(userList)
}
