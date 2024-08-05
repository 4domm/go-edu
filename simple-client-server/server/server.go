package server

import (
	"fmt"
	"log/slog"
	"net"
)

func StartServer(port string) {
	s, err := net.Listen("tcp", ":"+port)
	if err != nil {
		slog.Info("Listen err:", err)
	}
	defer s.Close()
	for {
		conn, err := s.Accept()
		if err != nil {
			slog.Info("Accept err:", err)
			return
		}
		go sendMessage(conn, "OK")
	}
}

func sendMessage(conn net.Conn, msg string) {
	_, err := fmt.Fprintf(conn, msg)
	if err != nil {
		slog.Info("Send err:", err)
		return
	}
	conn.Close()
}
