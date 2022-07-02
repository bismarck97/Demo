package main

import (
	"Demo/Chatroom/common/message"
	"fmt"
	"net"
)

// ServerProcessMes 功能:根据客户端发送消息种类不同，决定调用哪个函数来处理
func ServerProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		//err =  ServerProcessLogin(conn, mes)

	case message.RegisterMesType:
	//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}
