package main

import (
	"MP2/message"
	"bufio"
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

// showClients prints out username of clients currently connected to the server
func showClients(clients map[string]net.Conn) {
	rand.Seed(time.Now().Unix())
	status := []string{
		"is ready to party!", "is chatting away!", "is getting down and dirty!",
		"is talking my ear off!", "needs to be kicked out!", "has no off button!",
		"wont shut up!!!!", "is my least favorite person!", "... who let that guy in!",
		"...everyone be quiet he might just leave.", "is my second least favorite person",
	}

	fmt.Println("\nCurrent Participants:")
	for client, _ := range clients {
		fmt.Println("["+client+"]", status[rand.Intn(len(status)-1)])
	}

	return
}

// addClient() adds the new client to the clients map
func addClient(clients map[string]net.Conn, username string, conn net.Conn) {
	clients[username] = conn
	fmt.Println(username, "joined the chat!")
	showClients(clients)

	return
}

// getUsername() gets the username sent from the client
func getUsername(conn net.Conn) string {
	dec := gob.NewDecoder(conn)
	var username string
	dec.Decode(&username)

	return username
}

// serve() handles interclient messaging and client EXIT commands
func serve(clients map[string]net.Conn, conn net.Conn) {
	for {
		var msg message.Message
		// receives the message, decodes it
		dec := gob.NewDecoder(conn)
		dec.Decode(&msg)

		//if the message was exit command, delete that user from map
		if strings.ToUpper(msg.To) == "EXIT" {
			delete(clients, msg.From)
			conn.Close()
			message.Display(msg)
			return
		} else if msg.To != "" {
			newConn, found := clients[msg.To]
			if found == true {
				enc := gob.NewEncoder(newConn)
				enc.Encode(msg)
			} else {
				errorMsg := message.Message{msg.From, "SERVER", "No Such User" + "[" + msg.To + "]"}
				enc := gob.NewEncoder(clients[msg.From])
				enc.Encode(errorMsg)
			}

		}

	}

}

//waitForExit() waits for server EXIT command, informs all clients, then quits program
func waitForExit(clients map[string]net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	exit, _ := reader.ReadString('\n')
	if strings.ToUpper(exit) == "EXIT\n" {
		exitMsg := message.Message{"EXIT", "SERVER", "Server Exiting"}
		for _, element := range clients {
			enc := gob.NewEncoder(element)
			enc.Encode(exitMsg)
		}
		os.Exit(0)
	}
}

func main() {

	var clients = make(map[string]net.Conn) // a map to keep track of clients and its TCP connection
	var port string

	fmt.Print("Enter the port number: ")
	reader := bufio.NewReader(os.Stdin)
	port, _ = reader.ReadString('\n')
	port = ":" + port
	port = strings.TrimSpace(port)

	go waitForExit(clients)

	for {
		l, err := net.Listen("tcp", port)
		if err != nil {
			fmt.Println(err)
			return
		}

		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		l.Close()

		username := getUsername(c)
		addClient(clients, username, c)

		go serve(clients, c)
	}

}
