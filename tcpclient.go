package main

import "net"
import "fmt"

// client
func main() {
	ip := "127.0.0.1"
	port := 8888
	addr := fmt.Sprintf("%s:%d", ip, port)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("dial error:", err)
	}

	//var buff [1024]byte // err, is slice, not array
	buff := make([]byte, 1024)
	buff[0] = 'a'
	cnt, err := conn.Write(buff)
	if err != nil {
		fmt.Println("conn write err:", err)
	}

	fmt.Println("send length:", cnt, ",data:", string(buff))

	recvBuff := make([]byte, 1024)
	recvCnt, err := conn.Read(recvBuff)
	if err != nil {
		fmt.Println("conn recv err:", err)
	}
	fmt.Println("recv length:", recvCnt, ",data:", string(recvBuff))
}
