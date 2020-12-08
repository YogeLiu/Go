package main

import (
	"fmt"
	"net"
	"strings"
)

func main()  {
	// 开启监听
	listener, err := net.Listen("tcp", "0.0.0.0")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	// 阻塞等待用户连接
	conn, errAccept := listener.Accept()
	if err != nil {
		fmt.Println("errAccept=", errAccept)
		return
	}
	// 接收用户请求
	buf := make([]byte, 1024)
	n, errBuf := conn.Read(buf)
	if errBuf != nil {
		fmt.Println("errBuf=", errBuf)
		return
	}
	str := strings.ToUpper(string(buf[:n]))
	conn.Write([]byte(str))
	fmt.Println("buf = ", string(buf[:n]))
	// 关闭
	defer listener.Close()
}
