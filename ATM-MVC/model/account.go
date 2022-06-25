package model

import "fmt"

type Account struct {
	CardID     string  //用户卡号
	UserName   string  //用户名
	PassWord   string  //用户密码
	Money      float64 //账号余额
	QuotaMoney float64 //每次取现额度
}

// NewAccount 构造函数
func NewAccount(cardID, userName, PassWord string, money, quotaMoney float64) *Account {
	return &Account{
		CardID:     cardID,
		UserName:   userName,
		PassWord:   PassWord,
		Money:      money,
		QuotaMoney: quotaMoney,
	}
}

//GetInfo 返回用户的信息
func (a *Account) GetInfo() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t", a.CardID, a.UserName, a.Money, a.QuotaMoney)
}
