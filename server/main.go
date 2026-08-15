package main

import "fmt"

func main() {
	store := NewStore()

	server := NewServer(store)

	err := server.Start()
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
