package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func clientLoop(endpoint string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", endpoint)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	var buf [512]byte
	fmt.Print("Enter:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Bytes()
		if string(s) == "exit" {
			break
		}
		_, err := conn.Write(s)
		if err != nil {
			fmt.Println(err)
			return
		}
		n, err := conn.Read(buf[0:])
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Connection closed by the sever: %s\n", conn.RemoteAddr())
			} else {
				fmt.Println(err)
			}
			return
		}
		fmt.Println("Server responded with:", string(buf[0:n]))
		fmt.Print("Enter:")
	}
}
