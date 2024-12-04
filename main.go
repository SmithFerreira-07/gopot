package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	address := "2222" //port
	fmt.Printf("Starting honeypot on %s...\n", address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	log.Printf("Connection received from: %s", remoteAddr)

	banner := "SSH-2.0-OpenSSH_8.2p1 Ubuntu-4ubuntu0.5\n"
	conn.Write([]byte(banner))

	reader := bufio.NewReader(conn)
	for {
		conn.Write([]byte("Username: "))
		username, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading username: %v", err)
			return
		}

		conn.Write([]byte("Password: "))
		password, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading password: %v", err)
			return
		}

		log.Printf("Credentials from %s - Username: %s, Password: %s",
			remoteAddr,
			strings.TrimSpace(username),
			strings.TrimSpace(password),
		)

		conn.Write([]byte("Access denied\n"))
	}
}
