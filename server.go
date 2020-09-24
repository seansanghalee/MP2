package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

func addClient(clients map[string]net.Conn, username string, conn net.Conn) {
	clients[username] = conn
	return
}

func getUsername(conn net.Conn) string{
	dec := gob.NewDecoder(conn)
	var username string
	dec.Decode(&username)

	return username
}

//func serve(c net.Conn) {
//
//}

func main() {

	var clients = make(map[string]net.Conn) // a map to keep track of clients and its TCP connection
	var port string

	fmt.Print("Enter the port number: ")
	reader := bufio.NewReader(os.Stdin)
	port = reader

	for {
		l, err := net.Listen("tcp", port)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer l.Close()

		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		username := getUsername(c)
		addClient(clients, username, c)

		fmt.Println(clients)

		//go serve(c)
	}

}
