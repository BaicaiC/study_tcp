package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:10000")
	if err != nil {
		panic(err)
	}
	var input string

	buf := make([]byte, 100000)
	for {
		fmt.Scanf("%s\n", &input)
		if input == "n" {
			break
		}
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		log.Println("receive: ", n, buf)
	}
}
