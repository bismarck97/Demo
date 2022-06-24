package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//定义用户切片结构体，用来管理多个用户之间的操作，以及添加用户删除用户
type accountList struct {
	accountList []*account
}
type Users interface {
	login() *account
	signIn()
	deleteUser(*account)
	accountInquiry() int
	transfer(*account)
}

// NewAccount 初始化结构体
func NewAccount() *accountList {
	account1 := newAccount("10001", "张三", "123456", 10000.0, 5000)
	account2 := newAccount("10002", "李四", "123456", 10000.0, 5000)
	s := make([]*account, 0)
	s = append(s, account1)
	s = append(s, account2)
	return &accountList{s}
}

//查询当前银行是否有两个以上用户
func (l *accountList) accountInquiry() int {
	num := 0
	for i, _ := range l.accountList {
		num = i
	}
	return num
}

//用户转账
func (l *accountList) transfer(user *account) {
	var cardId string
	s := l.accountList
	//定义要转账的账户
	var transferUser *account
	if l.accountInquiry() >= 2 {
		for {
			fmt.Println("请输入要转账的账户:")
			fmt.Scan(&cardId)
			//根据id在切片中查找对应的数据
			if cardId == user.cardID {
				fmt.Println("您不能给自己转账，请重新输入")
			} else {
				for i := 0; i < len(s); i++ {
					if s[i].cardID == cardId {
						//找到对应的ID便存入要转账的账户
						transferUser = s[i]
						goto transfer
					}
					if i == len(s)-1 {
						//遍历完还没找到
						fmt.Println("卡号输入错误，请重新输入")
					}
				}
			}
		}
	transfer:
		var str string
		var money float64
		for {
			//go语言中，中文占用三个字节，直接截取时要从3开始截取
			//fmt.Printf("您当前要为:*%s转账\n", transferUser.userName[3:])
			//或者先转换为[]rune，截取后，在转回string
			nameRune := []rune(transferUser.userName)
			fmt.Printf("您当前要为:*%s转账\n", string(nameRune[1:]))
			fmt.Println("请您输入姓氏确认：")
			fmt.Scan(&str)
			if strings.HasPrefix(transferUser.userName, str) {
				for {
					fmt.Println("请输入要转账的金额")
					fmt.Scan(&money)
					if money > user.money {
						fmt.Println("转账金额超出您的余额，请重新操作")
					} else if money > user.quotaMoney {
						fmt.Println("超出转账金额限额，请重新操作")
					} else {
						user.money -= money
						transferUser.money += money
						goto ok
					}
				}

			} else {
				fmt.Println("输入不正确，请重新输入")
			}
		}
	ok:
		fmt.Println("转账成功")
	} else {
		fmt.Println("当前银行账户不足2个以上，无法完成转账")
	}

}

//用户登录验证
func (l *accountList) login() *account {
	s := l.accountList
	var a *account
	//定义变量保存卡号
	var cardId string
	//验证卡号
	for {
		fmt.Println("请输入卡号：")
		fmt.Scan(&cardId)
		for i := 0; i < len(s); i++ {
			if s[i].cardID == cardId {
				a = s[i]
				goto pass
			}
			if i == len(s)-1 {
				fmt.Println("卡号输入有误，请重新输入")
			}
		}
	}
pass:
	//验证密码
	//定义变量保存密码
	var passWord string
	//定义计时器
	num := 4
	for {
		fmt.Println("请输入密码：")
		fmt.Scan(&passWord)
		if passWord == a.passWord {
			fmt.Printf("贵宾%s,您好，欢迎进入系统，您的卡号:%s\n", a.userName, a.cardID)
			goto returnAccount
		} else {
			num--
			fmt.Printf("输入错误，还有%d次机会\n", num)
		}
		if num == 0 {
			fmt.Println("输入密码错误达到3次，返回ATM系统界面")
			return nil
		}
	}
returnAccount:
	return a
}

//用户注册
func (l *accountList) signIn() {
	rand.Seed(time.Now().UnixNano())
	num := "1234567890"
	var cardId string
	var name string
	var password string
	var okPassword string
	var quotaMoney float64
	fmt.Println("========欢迎进入银行用户办卡中心界面========")
	fmt.Println("请输入您的姓名：")
	fmt.Scan(&name)
	fmt.Println("请输入您的密码：")
	fmt.Scan(&password)
	for {
		fmt.Println("请您再次输入密码：")
		fmt.Scan(&okPassword)
		if okPassword == password {
			break
		} else {
			fmt.Println("输入不正确，请重新输入")
		}
	}
	fmt.Println("请设置当日取现额度")
	fmt.Scan(&quotaMoney)
	//创建随机卡号，并验证是否重复

	for {
		flag := true
		//重新循环一次清零cardId
		cardId = ""
		//生成8位数随机卡号
		for i := 0; i < 8; i++ {
			cardId += string(num[rand.Intn(10)])
			//cardId += num[rand.Intn(10) : rand.Intn(10)+1]
		}
		for i := 0; i < len(l.accountList); i++ {
			//判断卡号是否有重复或者开头为0
			if cardId == l.accountList[i].cardID || strings.HasPrefix(cardId, "0") {
				flag = false
			}
		}
		if flag {
			break
		}
	}
	account := newAccount(cardId, name, password, 0, quotaMoney)
	fmt.Printf("贵宾%s，您好，您的账户已开卡成功，您的卡号是:%s,\n", account.userName, account.cardID)
	//将用户添加到切片中
	l.accountList = append(l.accountList, account)
}

//删除用户
func (l *accountList) deleteUser(user *account) {
	var ok string
	fmt.Println("您真的确定要销毁您的账户吗")
	fmt.Println("y/n")
	fmt.Scan(&ok)
	if ok == "y" {
		if user.money > 0 {
			fmt.Println("您的账户还有钱没有取完，不允许销户")
			operate(user, l)
		} else {
			for i, v := range l.accountList {
				if v == user {
					//删除指定用户账户
					l.accountList = append(l.accountList[:i], l.accountList[i+1:]...)
				}
			}
			fmt.Println("账户销毁成功，返回首页")
			welcome(l)
		}
	} else {
		fmt.Println("账户继续保留，返回上一级")
		operate(user, l)
	}
}
