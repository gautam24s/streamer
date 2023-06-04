package main

import (
	"github.com/unfixbug/streamer/config"
	"github.com/unfixbug/streamer/server"
)

func main() {
	config.Init("development.yaml")
	server.Init()
}
