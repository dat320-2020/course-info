package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func serverLoop(endpoint string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", endpoint)
	if err != nil {
		log.Fatal(err)
	}
	lis, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Waiting for clients...")
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Connection closed by %s\n", conn.RemoteAddr())
			} else {
				fmt.Println(err)
			}
			return
		}
		_, err = conn.Write(buf[0:n])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[0:n]))
	}
}
