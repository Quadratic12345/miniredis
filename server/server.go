package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Server struct {
	store *Store
}

func NewServer(store *Store) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		return err
	}

	defer ln.Close()

	fmt.Println("Mini Redis listening on :9999")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client connected:", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		command := scanner.Text()

		response := s.handleCommand(command)

		fmt.Fprintln(conn, response)
	}

	fmt.Println("Client disconnected:", conn.RemoteAddr())
}

func (s *Server) handleCommand(command string) string {
	parts := strings.Fields(command)

	if len(parts) == 0 {
		return "ERR empty command"
	}

	switch strings.ToUpper(parts[0]) {

	case "SET":
		if len(parts) < 3 {
			return "ERR usage: SET key value"
		}

		key := parts[1]
		value := strings.Join(parts[2:], " ")

		s.store.Set(key, value)

		return "OK"

	case "GET":
		if len(parts) != 2 {
			return "ERR usage: GET key"
		}

		key := parts[1]

		value, exists := s.store.Get(key)

		if !exists {
			return "(nil)"
		}

		return value

	case "DEL":
		if len(parts) != 2 {
			return "ERR usage: DEL key"
		}

		key := parts[1]

		deleted := s.store.Delete(key)

		if deleted {
			return "1"
		}

		return "0"

	default:
		return "ERR unknown command"
	}
}
