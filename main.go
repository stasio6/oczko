package main

import (
	"fmt"
	"os"

	"oczko.com/client"
	"oczko.com/server"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "server" {
		fmt.Println("Running in server mode")
		server.RunServer()
	} else {
		fmt.Println("Running in client mode")
		client.MainClient()
	}
}
