package main

import (
	"./defs"
	server "./server"
	"fmt"
)


func main() {
	defs.LoadConfigs()
	fmt.Println("Start log-server..")
	server := new(server.Server)
	server.Run()
}
