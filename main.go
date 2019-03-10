package main

import (
	"fmt"
	"os"
)

func main() {
	if 1 == len(os.Args) {
		TCPClient()
	}

	switch os.Args[1] {
	case "tcp":
		TCPServer()
	case "ws":
		WSServer()
	default:
		fmt.Printf(`Usage:
%s        tcp client mode
%s tcp    tcp server mode
%s ws     websocket server mode
`, os.Args[0], os.Args[0], os.Args[0])
	}
}
