package main

import (
	_ "collect/api/config"
	"collect/api/server"
)

func main() {
	server.Init()
}
