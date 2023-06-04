package server

import "github.com/unfixbug/streamer/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.hostport"))
}
