package core

import (
	"log"
	"net"

	"golang.org/x/net/websocket"
)

type ClientManager struct {
	clients    map[*Client]chan *Message
	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		clients:    make(map[*Client]chan *Message),
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (cm *ClientManager) Start() {
	for {
		select {
		case client := <-cm.register:
			cm.clients[client] = make(chan *Message)
			client.ready <- struct{}{}
			go func() {
				cm.broadcast <- client.loginMessage()
			}()
		case client := <-cm.unregister:
			close(cm.clients[client])
			delete(cm.clients, client)
			go func() {
				cm.broadcast <- client.logoutMessage()
			}()
		case message := <-cm.broadcast:
			for _, ch := range cm.clients {
				ch <- message
			}
		}
	}
}

func (cm *ClientManager) Serve(client *Client) {
	exit := make(chan struct{})
	defer close(exit)

	<-client.ready
	go cm.receiveFrom(client, exit)
	go cm.sendTo(client, exit)
	close(client.ready)

	<-exit
	client.Logout()
	<-exit
}

func (cm *ClientManager) receiveFrom(client *Client, exit chan struct{}) {
	for {
		buff := make([]byte, 4096)
		n, err := client.conn.Read(buff)
		if err != nil {
			break
		}
		if n > 0 {
			cm.broadcast <- client.NewMessage(string(buff[:n]))
		}
	}
	exit <- struct{}{}
}

func (cm *ClientManager) sendTo(client *Client, exit chan struct{}) {
	for {
		Message, ok := <-cm.clients[client]
		if !ok {
			break
		}

		var err error
		switch conn := client.conn.(type) {
		case *net.TCPConn:
			_, err = conn.Write(Message.ToString())
		case *websocket.Conn:
			_, err = conn.Write(Message.ToJson())
		}
		if err != nil {
			log.Println(err)
			break
		}
	}
	exit <- struct{}{}
}
