package main

import (
	"./defs"
	global "./global"
	server "./server"
	"fmt"
)

func main() {
	defs.LoadEnvs()

	global.InitMongoService()

	fmt.Println("Start log-server..")
	server := new(server.Server)
	server.Run()
}
