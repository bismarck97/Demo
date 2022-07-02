package main

import (
	"Demo/Chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

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
