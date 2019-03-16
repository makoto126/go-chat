package core

import (
	"encoding/json"
	"log"
	"net"
	"strings"
)

type Client struct {
	conn  net.Conn
	name  string
	cm    *ClientManager
	ready chan struct{}
}

func NewClient() *Client {
	return &Client{
		ready: make(chan struct{}),
	}
}

func (c *Client) Login(conn net.Conn, cm *ClientManager) {
	buff := make([]byte, 32)
	n, err := conn.Read(buff)
	if err != nil {
		log.Panic(err)
	}
	c.name = strings.TrimSpace(string(buff[:n]))
	c.conn = conn
	c.cm = cm
	cm.register <- c
}

func (c *Client) Logout() {
	c.cm.unregister <- c
	c.conn.Close()
}

func (c *Client) NewMessage(s string) *Message {
	return &Message{
		From: c.name,
		Text: s,
	}
}

func (c *Client) loginMessage() *Message {
	return c.NewMessage("joined!\n")
}

func (c *Client) logoutMessage() *Message {
	return c.NewMessage("leaved!\n")
}

type Message struct {
	From string `json:"from,omitempty"`
	Text string `json:"text,omitempty"`
}

func (m *Message) ToString() []byte {
	return []byte(m.From + ": " + m.Text)
}

func (m *Message) ToJson() []byte {
	b, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
	}
	return b
}
