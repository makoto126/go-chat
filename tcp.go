package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/makoto126/go-chat/core"
)

func TCPServer() {
	fmt.Println("Starting tcp server...")
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Panic(err)
	}
	cm := core.NewClientManager()
	go cm.Start()

	for {
		conn, _ := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		client := core.NewClient()
		client.Login(conn, cm)
		go cm.Serve(client)
	}
}

func TCPClient() {
	fmt.Println("Starting tcp client...")
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Firstly, please input your name: ")
	go func() {
		io.Copy(os.Stdout, conn)
	}()
	io.Copy(conn, os.Stdin)
}
