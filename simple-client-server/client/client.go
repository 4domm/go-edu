package client

import (
	"fmt"
	"log/slog"
	"net"
	"strconv"
)

func StartClient(port int) {
	conn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(port))
	slog.Info("connected to server")
	if err != nil {
		slog.Info("cant connect to server", err)
	}
	ansFromServer := make([]byte, 1024)
	size, err := conn.Read(ansFromServer)
	if err != nil {
		slog.Info("cant read from server", err)
	}
	slog.Info("read message")
	if string(ansFromServer[:size]) == "OK" {
		fmt.Println(string(ansFromServer[:size]))
		conn.Close()
	}
}
