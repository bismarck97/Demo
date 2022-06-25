package service

import (
	"Demo/ATM-MVC/model"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	service = NewAccountService()
}

// Service 声明接口
var service *AccountService

// AccountService 该accountService，完成对Account的操作，包括增删改查
type AccountService struct {
	//声明一个字段，表示当前切片含有多少个客户
	account []*model.Account
	//用户卡号
	cardId string
}

// NewAccountService 编写一个方法，可以返回*AccountService
func NewAccountService() *AccountService {
	account1 := model.NewAccount("10000001", "张三", "123456", 10000.0, 5000)
	account2 := model.NewAccount("10000002", "李四", "123456", 10000.0, 5000)
	s := make([]*model.Account, 0)
	s = append(s, account1)
	s = append(s, account2)
	return &AccountService{s, ""}
}

// AccountServiceEr 定义用户service层接口
type AccountServiceEr interface {
	RandCardId() (cardId string)
	SignIn(userName string, passWord string, quotaMoney float64) (name, id string, ok bool)
	selectAccount(id string) (user *model.Account, flag bool)
	SelectCardId(id string) (flag bool)
	SelectPassWord(id string, password string) (string, bool)
	SelectUser(id string) (info string)
	SetMoney(id string, money float64) (float64, bool)
	QueryMoney(id string) (money, quotaMoney float64)
	AccountInquiry() int
	WithdrawingMoney(id string, money float64) (float64, bool)
	SelectUserName(id string) string
	TransferMoney(money float64, id1 string, id2 string) (float64, bool)
	QueryPassWord(password string, id string) bool
	ChangePassword(password2 string, id string) bool
	DeleteAccount(id string) bool
}

// RandCardId 生成随机卡号
func (s *AccountService) RandCardId() string {
	num := "1234567890"
	for {
		flag := true
		//重新循环一次清零cardId
		s.cardId = ""
		//生成8位数随机卡号
		for i := 0; i < 8; i++ {
			s.cardId += string(num[rand.Intn(10)])
			//cardId += num[rand.Intn(10) : rand.Intn(10)+1]
		}
		for i := 0; i < len(s.account); i++ {
			//判断卡号是否有重复或者开头为0
			if s.cardId == s.account[i].CardID || strings.HasPrefix(s.cardId, "0") {
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return s.cardId

}

// SignIn 用户注册
func (s *AccountService) SignIn(userName string, passWord string, quotaMoney float64) (name, id string, ok bool) {
	//生成随机ID
	cardId := service.RandCardId()
	//生成用户
	account := model.NewAccount(cardId, userName, passWord, 0, quotaMoney)
	//保存到切片中
	s.account = append(s.account, account)
	//查询是否成功
	u, ok := service.selectAccount(cardId)
	return u.UserName, u.CardID, ok
}

//根据id返回对应的用户
func (s *AccountService) selectAccount(id string) (user *model.Account, flag bool) {
	flag = true
	for i := 0; i < len(s.account); i++ {
		if s.account[i].CardID == id {
			user = s.account[i]
			break
		}
		//i如果和account切片的长度相同，则证明循环完毕都没找到对应的用户，返回false
		if i == len(s.account)-1 {
			flag = false
		}
	}
	return
}

// SelectCardId 根据ID查找是否有用户
func (s *AccountService) SelectCardId(id string) (flag bool) {
	_, f := service.selectAccount(id)
	return f
}

// SelectPassWord 判断是否输入正确
func (s *AccountService) SelectPassWord(id string, password string) (string, bool) {
	u, _ := service.selectAccount(id)
	return u.UserName, u.PassWord == password
}

// SelectUser 根据Id查找用户信息
func (s *AccountService) SelectUser(id string) (info string) {
	u, _ := service.selectAccount(id)
	return u.GetInfo()
}

// SetMoney 用户存钱
func (s *AccountService) SetMoney(id string, money float64) (float64, bool) {
	u, _ := service.selectAccount(id)
	u.Money += money
	return u.Money, true
}

// QueryMoney 查询余额和取现额度
func (s *AccountService) QueryMoney(id string) (money, quotaMoney float64) {
	u, _ := service.selectAccount(id)
	return u.Money, u.QuotaMoney

}

// AccountInquiry 查询当前银行有几个用户
func (s *AccountService) AccountInquiry() int {
	//num := 0
	//for i, _ := range s.account {
	//	num = i
	//}
	//fmt.Println(num)
	return len(service.account)
}

// WithdrawingMoney 用户取款
func (s *AccountService) WithdrawingMoney(id string, money float64) (float64, bool) {
	u, _ := service.selectAccount(id)
	u.Money -= money
	return u.Money, true
}

// SelectUserName 查询用户姓名
func (s *AccountService) SelectUserName(id string) string {
	u, _ := service.selectAccount(id)
	return u.UserName
}

// TransferMoney 用户转账操作
func (s *AccountService) TransferMoney(money float64, id1 string, id2 string) (float64, bool) {
	m, ok := service.WithdrawingMoney(id1, money)
	service.SetMoney(id2, money)
	return m, ok
}

// QueryPassWord 查询密码
func (s *AccountService) QueryPassWord(password string, id string) bool {
	u, _ := service.selectAccount(id)
	return u.PassWord == password
}

// ChangePassword 修改密码
func (s *AccountService) ChangePassword(password2 string, id string) bool {
	u, f := service.selectAccount(id)
	u.PassWord = password2
	return f
}

// DeleteAccount 删除指定用户
func (s *AccountService) DeleteAccount(id string) bool {
	v, b := service.selectAccount(id)

	for i, account := range s.account {
		if account == v {
			//删除指定用户账户
			s.account = append(s.account[:i], s.account[i+1:]...)
		}
	}
	return b
}
