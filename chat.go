package main
import (
    "net"
    "fmt"
)

func main () {
    listener, err := net.Listen("tcp", "127.0.0.1:8080")
    if err != nil { fmt.Println(err) }
	    fmt.Println("listening")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println(err)
            return
        }    
    	go handler(conn)
    }
    
    
}

func handler(conn net.Conn) {
    for {
        buff := make([]byte, 1024)
        _, err := conn.Read(buff)
        if err != nil { fmt.Println(err)}
        fmt.Println("get string:", string(buff))
    }
}
