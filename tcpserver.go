package main

import "net"
import "fmt"

// server
func main() {
	ip := "127.0.0.1"
	port := 8888
	addr := fmt.Sprintf("%s:%d", ip, port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("linsten error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("accept error:", err)
	}

	//var buff [1024]byte // err, is slice, not array
	buff := make([]byte, 1024)
	cnt, err := conn.Read(buff)
	if err != nil {
		fmt.Println("conn read err:", err)
	}

	fmt.Println("get length:", cnt, ",data:", string(buff))

	sendCnt, err := conn.Write(buff)
	if err != nil {
		fmt.Println("send err:", err)
	}
	fmt.Println("send length:", sendCnt, ",data:", string(buff))

}
