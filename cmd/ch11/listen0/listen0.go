package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = conn.Write([]byte("Nice to meet you!"))
		if err != nil {
			log.Fatal(err)
		}
		conn.Close()
	}

}
