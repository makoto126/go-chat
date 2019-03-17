package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/makoto126/go-chat/core"
	"golang.org/x/net/websocket"
)

func WSServer() {
	fmt.Println("Starting ws server...")
	cm := core.NewClientManager()
	go cm.Start()

	var chat websocket.Handler
	chat = func(ws *websocket.Conn) {
		client := core.NewClient()
		client.Login(ws, cm)
		cm.Serve(client)
	}

	http.Handle("/ws", chat)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Panic(err)
	}
}
