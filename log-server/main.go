package main

import (
	server "github.com/sshtel/app-logger/log-server/server"
	"fmt"
)

func main() {
	fmt.Println("Start log-server..")
	server := new(server.Server)
	server.Run()
}
