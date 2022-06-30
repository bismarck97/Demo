package main

import (
	"fmt"
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

//处理和客户端的通讯
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()
	//循环的读客户端发送的信息
	buf := make([]byte, 1024*4)
	for {
		fmt.Println("等待读取...")
		_, err := conn.Read(buf[:4])
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("读到的buf = ", buf[:4])
	}
}
