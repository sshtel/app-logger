package main

import (
	"fmt"

	"github.com/sshtel/app-logger/log-server/defs"
	"github.com/sshtel/app-logger/log-server/global"
	"github.com/sshtel/app-logger/log-server/server"
)

func main() {
	defs.LoadEnvs()

	global.InitMongoService()

	fmt.Println("Start log-server..")
	server := new(server.Server)
	server.Run()
}
