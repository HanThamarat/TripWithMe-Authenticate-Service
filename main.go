package main

import (
	"github.com/HanThamarat/TripWithMe-Authenticate-Service/packages/conf"
	"github.com/HanThamarat/TripWithMe-Authenticate-Service/server"
)

func main() {
	config := conf.GetConfig();

	server.NewFiberServer(config).Start();
}