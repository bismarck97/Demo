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

//编辑一个函数serverProcessLogin函数，专门处理登录请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
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
	err = WritePkg(conn, data)
	return
}

//编写一个ServerProcessMes函数
//功能:根据客户端发送消息种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
	//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

//处理和客户端的通讯
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()
	//循环的读客户端发送的信息
	for {
		//将读取数据包，直接封装成一个函数readPkg(),返回Message,err
		mes, err := ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端退出...")
				return
			} else {
				fmt.Println("readPkg err:", err)
				return
			}
		}
		//fmt.Println("mes:", mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			fmt.Println("serverProcessMes err:", err)
			return
		}
	}
}

// WritePkg 发送数据包
func WritePkg(conn net.Conn, data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	//1.发送长度
	n, err := conn.Write(buf[:])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes[:]) fail", err)
		return
	}
	//2.发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}

// ReadPkg 获取数据包的内容，返回Message对象
func ReadPkg(conn net.Conn) (mes message.Message, err error) {
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

	//把pkgLen 反序列化成 ->message.Message   &mes!!!
	err = json.Unmarshal(buf[:pkgLen], &mes) //要操作指针才有用
	if err != nil {
		return
	}
	return
}
