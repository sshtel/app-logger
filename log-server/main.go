package main

import (
	server "./server"
	"fmt"
)

func main() {
	fmt.Println("Start log-server..")
	server := new(server.Server)
	server.Run()
}
