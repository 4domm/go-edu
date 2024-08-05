package main

import (
	"log/slog"
	"os"
	"simple-client-server/client"
	"simple-client-server/config"
	"simple-client-server/server"
)

func main() {
	if len(os.Args) < 2 {
		slog.Info("Usage: path/config.yaml")
		return
	}
	conf, err := config.MustLoadConfig(os.Args[1])
	if err != nil {
		slog.Info("cant load config", err)
		return
	}
	go server.StartServer(conf.Port)
	client.StartClient(conf.Port)
}
