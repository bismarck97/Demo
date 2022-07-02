package process

import (
	"Demo/Chatroom/common/message"
	"Demo/Chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

// ServerProcessLogin 专门处理登录请求
func ServerProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//核心代码...
	//1.先从mes中取出mes.Data,并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2.再声明一个 LoginResMes，并完成赋值
	var loginResMes message.LoginResMes

	//如果用户的id=100，密码=123456，认为合法，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法的
		loginResMes.Code = 200

	} else {
		//不合法的
		loginResMes.Code = 500 //500状态码，表示该用户不存在
		loginResMes.Error = "该用户不存在,请注册再使用"
	}
	//3.将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)
	//5.对resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//6.发送data 我们将其封装到writePkg函数
	err = utils.WritePkg(conn, data)
	return
}
