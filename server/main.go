package main

import (
	"fmt"
	"log"
	"net"
)

func server(conn net.Conn) {
	var input string
	s := make([]byte, 20000000)
	for i := 0; i < 20000000; i++ {
		s[i] = byte(i)
	}

	for {
		n, err := conn.Write(s)
		if err != nil {
			return
		}
		fmt.Scanf("%s\n", &input)
		if input == "n" {
			break
		}
		log.Println("send: ", n, "test segment d")
	}
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:10000")
	if err != nil {
		panic(err)
	}

	log.Println("server run ...")
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		log.Println("client: ", conn.RemoteAddr().String())
		go server(conn)
	}

}
