package main

import (
	"net"
	"fmt"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":1208")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server started ...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		go connHandler(conn)
	}
}

func connHandler(conn net.Conn) {
	if conn == nil {
		return
	}
	buf := make([]byte, 4096)
	for {
		cnt, err := conn.Read(buf)
		fmt.Printf("Read reads data %d\n", cnt)
		if err != nil || cnt == 0 {
			conn.Close()
			break
		}
		inStr := strings.TrimSpace(string(buf[0:cnt]))
		split := strings.Split(inStr, " ")
		fmt.Println(conn.RemoteAddr())
		switch split[0] {
		case "ping":
			conn.Write([]byte("pong\n"))
		case "echo":
			echoStr := strings.Join(split[1:], " ") + "\n"
			conn.Write([]byte(echoStr))
		case "quit":
			conn.Close()
		default:
			fmt.Printf("Unsupported command: %s\n", split[0])
		}
	}
	fmt.Printf("Connection from %v closed. \n", conn.RemoteAddr())
}
