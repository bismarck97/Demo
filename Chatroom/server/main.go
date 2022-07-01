package main

import (
	"Demo/Chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func main() {
	//提示信息
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listen.Close()
	//一旦监听成功，就等到客户端来链接服务器
	for {
		fmt.Println("等待客户端来链接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err:=", err)
			continue
		}
		//一旦链接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}

//编写一个ServerProcessMes函数
//功能:根据客户端发送消息种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录的逻辑

	}
}

//处理和客户端的通讯
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()
	//循环的读客户端发送的信息
	for {
		//将读取数据包，直接封装成一个函数readPkg(),返回Message,err
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端退出...")
				return
			} else {
				fmt.Println("readPkg err:", err)
				return
			}
		}
		fmt.Println("mes:", mes)
	}
}
func readPkg(conn net.Conn) (mes message.Message, err error) {
	//创建一个切片往里写内容
	buf := make([]byte, 1024*4)
	fmt.Println("等待读取...")
	//conn.Read在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了 conn 则，就不会阻塞
	_, err = conn.Read(buf[:4]) //第一次先读取长度
	if err != nil {
		return
	}
	//根据读到的buf[:4]转成uint32类型
	pkgLen := binary.BigEndian.Uint32(buf[0:4])
	//根据pkgLen 读取消息内容
	n, err := conn.Read(buf[:pkgLen]) //第二次根据长度读取数据
	if n != int(pkgLen) || err != nil {
		return
	}
	fmt.Println("buf[:pkgLen]", buf[:pkgLen])
	//把pkgLen 反序列化成 ->message.Message   &mes!!!
	err = json.Unmarshal(buf[:pkgLen], &mes) //要操作指针才有用
	if err != nil {
		return
	}
	return
}
