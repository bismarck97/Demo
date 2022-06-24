package main

import "fmt"

//定义用户账户结构体，用来管理单个用户的操作
type account struct {
	cardID     string  //用户卡号
	userName   string  //用户名
	passWord   string  //用户密码
	money      float64 //账号余额
	quotaMoney float64 //每次取现额度
}

//构造函数
func newAccount(cardID, userName, PassWord string, money, quotaMoney float64) *account {
	return &account{
		cardID:     cardID,
		userName:   userName,
		passWord:   PassWord,
		money:      money,
		quotaMoney: quotaMoney,
	}
}

// User 单个用户操作接口
type User interface {
	Select()
	SetMoney()
	withdrawingMoney()
	changePassword(s *accountList)
}

// Select 查询当前账户信息
func (user *account) Select() {
	fmt.Println("========欢迎您进入银行用户详情页面========")
	fmt.Println("您的账户信息如下：")
	fmt.Println("卡号：", user.cardID)
	fmt.Println("姓名：", user.userName)
	fmt.Println("余额：", user.money)
	fmt.Println("取现额度：", user.quotaMoney)
}

// SetMoney 用户存款
func (user *account) SetMoney() {
	fmt.Println("请输入您要存入的金额")
	var money float64
	fmt.Scan(&money)
	user.money += money
	fmt.Println("存款成功，您的账户当前余额为：", user.money)
}

//用户取款
func (user *account) withdrawingMoney() {
	fmt.Println("========欢迎您进入银行用户取款页面========")
	var money float64
	fmt.Printf("当前账户可用余额为：%f\n", user.money)
	if user.money == 0 {
		fmt.Println("账户当前金额为0，不能完成取款操作")
		return
	}
	for {
		fmt.Println("请输入取款金额")
		fmt.Scan(&money)
		if money > user.money {
			fmt.Println("您的账户余额不足")
		} else if money > user.quotaMoney {
			fmt.Println("您当前取款超过当此限额")
		} else {
			user.money -= money
			fmt.Println("您已成功取款,您的账户余额为：", user.money)
			return
		}
	}
}

//用户修改密码
func (user *account) changePassword(s *accountList) {
	var password string
	var newPassword string
	var newPassword2 string
	for {
		fmt.Println("请您输入当前账户密码")
		fmt.Scan(&password)
		if password == user.passWord {
			for {
				fmt.Println("请输入新的密码")
				fmt.Scan(&newPassword)
				for {
					fmt.Println("请再次输入确认密码")
					fmt.Scan(&newPassword2)
					if newPassword == newPassword2 {
						user.passWord = newPassword2
						fmt.Println("密码修改成功，请重新登录")
						goto ok
					} else {
						fmt.Println("确认密码输入不正确")
					}
				}
			}
		} else {
			fmt.Println("当前用户密码输入不正确，请重新输入！")
		}
	}
ok:
	//调用登录界面
	welcome(s)
}
