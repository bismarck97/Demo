package main

import (
	"Demo/ATM-MVC/service"
	"fmt"
	"strings"
)

type accountView struct {
	//定义必要字段
	key        string //接收用户输入...
	cardId     string
	passWord   string
	loop       bool //表示是否循环的显示主菜单
	money      float64
	userName   string
	quotaMoney float64
	flag       string
	//增加一个accountService字段
	accountService *service.AccountService
}
type Viewer interface {
	welcome(*service.AccountService)
	login()
	signIn()
	operate(id string)
	selectAccount(id string)
	setMoney(id string)
	withdrawingMoney(id string)
	transferMoney(id string)
	changePassword(id string)
	deleteAccount(id string)
}

//用户欢迎界面
func (v *accountView) welcome(*service.AccountService) {
	for {
		fmt.Println("========欢迎进入银行ATM系统========")
		fmt.Println("1.登录账号")
		fmt.Println("2.注册开户")
		fmt.Println("3.退出系统")
		fmt.Println("请您选择操作:")
		fmt.Scanln(&v.key)
		switch v.key {
		case "1":
			//用户登录操作
			view.login()
		case "2":
			view.signIn()
		case "3":
			//退出系统
			v.loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
		if !v.loop {
			fmt.Println("欢迎再次使用银行ATM系统")
			break
		}
	}
	return
}

//用户登录页面
func (v *accountView) login() {
	for {
		fmt.Println("请输入卡号：")
		fmt.Scanln(&v.cardId)
		if v.accountService.SelectCardId(v.cardId) {
			//定义计时器
			num := 4
			for {
				fmt.Println("请输入密码：")
				fmt.Scanln(&v.passWord)
				if name, ok := v.accountService.SelectPassWord(v.cardId, v.passWord); ok {
					fmt.Printf("贵宾%s,您好，欢迎进入银行用户操作界面，您的卡号为:%s\n", name, v.cardId)
					//跳到其他函数应该结束这个函数
					view.operate(v.cardId)
					return
				} else {
					num--
					fmt.Printf("输入错误，还有%d次机会\n", num)
				}
				if num == 0 {
					fmt.Println("输入密码错误达到3次，返回ATM系统界面")
					view.welcome(v.accountService)
					//跳到其他函数应该结束这个函数
					return
				}
			}
		} else {
			fmt.Println("卡号输入有误，请重新输入")
		}
	}
}

//用户注册页面
func (v *accountView) signIn() {
	fmt.Println("========欢迎进入银行用户办卡中心界面========")
	var okPassword string
	fmt.Println("请输入您的姓名：")
	fmt.Scanln(&v.userName)
	fmt.Println("请输入您的密码：")
	fmt.Scanln(&v.passWord)
	for {
		fmt.Println("请您再次输入密码：")
		fmt.Scanln(&okPassword)
		if okPassword == v.passWord {
			break
		} else {
			fmt.Println("输入不正确，请重新输入")
		}
	}
	for {
		fmt.Println("请设置当日取现额度")
		fmt.Scanln(&v.quotaMoney)
		if v.quotaMoney < 5000 {
			fmt.Println("每日取现额度不能小于5000，请重新设置")
		} else {
			if userName, cardId, ok := v.accountService.SignIn(v.userName, v.passWord, v.quotaMoney); ok {
				fmt.Printf("贵宾%s，您好，您的账户已开卡成功，您的卡号是:%s\n", userName, cardId)
				view.welcome(v.accountService)
				return
			}
		}
	}
}

//用户操作界面
func (v *accountView) operate(id string) {
	for {
		fmt.Println("========欢迎您进入银行用户操作界面========")
		fmt.Println("1.查询\t2.存款:\t3.取款\t4.转账")
		fmt.Println("5.修改密码\t6.退出\t7.注销当前账号")
		fmt.Println("请您选择操作:")
		fmt.Scanln(&v.key)
		switch v.key {
		case "1":
			//根据id查找对应的数据
			view.selectAccount(id)

		case "2":
			//根据id完成存款操作
			view.setMoney(id)

		case "3":
			view.withdrawingMoney(id)

		case "4":
			view.transferMoney(id)

		case "5":
			view.changePassword(id)

		case "6":
			v.loop = false
		case "7":
			view.deleteAccount(id)
		}
		if !v.loop {
			//跳到其他函数应该结束这个函数
			view.welcome(v.accountService)
			return
		}
	}
}

//显示用户数据
func (v *accountView) selectAccount(id string) {
	fmt.Println("您的账户信息如下：")
	fmt.Println("卡号\t\t姓名\t余额\t取现额度")
	info := v.accountService.SelectUser(id)
	fmt.Println(info)
	return
}

//存款操作
func (v *accountView) setMoney(id string) {
	fmt.Println("请输入您要存入的金额")
	fmt.Scanln(&v.money)
	if money, ok := v.accountService.SetMoney(id, v.money); ok {
		fmt.Println("存款成功，当前余额为:", money)
		return
	}
}

//取款操作
func (v *accountView) withdrawingMoney(id string) {
	fmt.Println("========欢迎您进入银行用户取款页面========")
	money, quotaMoney := v.accountService.QueryMoney(id)
	fmt.Println("当前账户可用余额为：", money)
	if money == 0 {
		fmt.Println("账户当前金额为0，不能完成取款操作")
		return
	}
	for {
		fmt.Println("请输入取款金额")
		fmt.Scanln(&v.money)
		if v.money > money {
			fmt.Println("您的账户余额不足")
		} else if v.money > quotaMoney {
			fmt.Println("您当前取款超过每日取款限额")
		} else {
			if m, ok := v.accountService.WithdrawingMoney(id, v.money); ok {
				fmt.Println("您已成功取款,您的账户余额为：", m)
				return
			}
		}
	}
}

//转账
func (v *accountView) transferMoney(id string) {
	//先判断是否有两个以上账户
	if v.accountService.AccountInquiry() < 2 {
		fmt.Println("当前银行账户不足2个，不能完成转账操作")
		return
	}
	for {
		//设置转账账户
		var transferId string
		fmt.Println("请输入要转账的账户:")
		fmt.Scanln(&transferId)
		if transferId == id {
			fmt.Println("您不能给自己转账，请重新输入")
			//查找银行内是否有此账户
		} else if !v.accountService.SelectCardId(transferId) {
			fmt.Println("卡号输入错误，请重新输入")
		} else {
			//定义变量接收用户姓氏
			var lastName string
			money, quotaMoney := v.accountService.QueryMoney(id)
			//go语言中，中文占用三个字节，先转换为[]rune，截取后，在转回string
			nameRune := []rune(v.accountService.SelectUserName(transferId))
			fmt.Printf("您当前要为:*%s转账\n", string(nameRune[1:]))
			fmt.Println("请您输入姓氏确认：")
			fmt.Scanln(&lastName)
			if strings.HasPrefix(v.accountService.SelectUserName(transferId), lastName) {
				for {
					fmt.Println("请输入要转账的金额")
					fmt.Scanln(&v.money)
					if v.money > money {
						fmt.Println("转账金额超出您的余额，请重新操作")
					} else if v.money > quotaMoney {
						fmt.Println("超出转账金额限额，请重新操作")
					} else {
						if m, ok := v.accountService.TransferMoney(v.money, id, transferId); ok {
							fmt.Println("转账成功，您的当前余额为：", m)
							view.operate(id)
							return
						} else {
							fmt.Println("转账失败")
						}
					}
				}
			} else {
				fmt.Println("姓氏输入不正确，请重新输入")
			}
		}

	}
}

//修改密码
func (v *accountView) changePassword(id string) {
	var newPassword string
	var newPassword2 string
	for {
		fmt.Println("请您输入当前密码:")
		fmt.Scanln(&v.passWord)
		//查询密码是否相等
		if v.accountService.QueryPassWord(v.passWord, id) {
			for {
				fmt.Println("请输入新的密码：")
				fmt.Scanln(&newPassword)
				for {
					fmt.Println("请再次输入新的密码：")
					fmt.Scanln(&newPassword2)
					if newPassword == newPassword2 {
						if v.accountService.ChangePassword(newPassword2, id) {
							fmt.Println("您的密码修改成功，请重新登录")
							view.welcome(v.accountService)
							return
						}
					} else {
						fmt.Println("确认密码输入不正确")
					}
				}
			}
		} else {
			fmt.Println("当前用户密码输入不正确，请重新输入！")
		}
	}
}

//删除用户
func (v *accountView) deleteAccount(id string) {
	fmt.Println("您真的确定要销毁您的账户吗")
	fmt.Println("y/n")
	fmt.Scanln(&v.flag)
	money, _ := v.accountService.QueryMoney(id)
	if v.flag == "y" {
		if money > 0 {
			fmt.Println("您的账户还有钱没有取完，不允许销户")
		} else {
			if v.accountService.DeleteAccount(id) {
				fmt.Println("账户销毁成功，返回首页")
				v.loop = false
			} else {
				fmt.Println("账户销毁失败，返回上一级")
			}
		}
	} else {
		fmt.Println("账户继续保留，返回上一级")
	}
}

//定义接口
var view Viewer

func main() {
	//在主函数中创建一个accountView，并运行显示主菜单
	accountView := accountView{
		key:        "",
		loop:       true,
		userName:   "",
		cardId:     "",
		passWord:   "",
		money:      0.0,
		quotaMoney: 0.0,
		flag:       "",
		//完成对accountView结构体的accountService字段的初始化
		accountService: service.NewAccountService(),
	}
	//显示主菜单
	view = &accountView
	view.welcome(accountView.accountService)
}
