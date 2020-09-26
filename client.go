package main

import (
	"MP2/message"
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
)

func login() (string, string, string) {
	var ip, port, username string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the host address: ")
	ip, _ = reader.ReadString('\n')
	ip = strings.TrimSpace(ip)
	fmt.Print("Enter the port number: ")
	port, _ = reader.ReadString('\n')
	port = strings.TrimSpace(port)
	fmt.Print("Enter your username: ")
	username, _ = reader.ReadString('\n')
	username = strings.TrimSpace(username)

	return ip, port, username
}

func sendUsername(username string, conn net.Conn) {
	enc := gob.NewEncoder(conn)
	enc.Encode(username)

	return
}

func receiveMessage(conn net.Conn) {
	for {
		var msg message.Message
		dec := gob.NewDecoder(conn)
		dec.Decode(&msg)

		// is the message an exit signal from the server
		if msg.To == "EXIT" {
			message.Display(msg)
			os.Exit(0)
		} else {
			if msg.Content != "" {
				message.Display(msg)
			}
		}
	}
}

func sendMessage(username string, conn net.Conn) {
	for {
		var to, content string
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("To: ")
		to, _ = reader.ReadString('\n')
		to = strings.TrimSpace(to)
		if to == "EXIT" {
			msg := message.Construct(to, username, username+" is exiting the server...")
			enc := gob.NewEncoder(conn)
			enc.Encode(msg)
			fmt.Println("Exiting...")
			os.Exit(0)
		}
		fmt.Print("Enter your message: ")
		content, _ = reader.ReadString('\n')
		content = strings.TrimSpace(content)
		msg := message.Construct(to, username, content)
		enc := gob.NewEncoder(conn)
		enc.Encode(msg)
	}
}

func main() {

	ip, port, username := login()

	address := ip + ":" + port
	c, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	sendUsername(username, c)

	go receiveMessage(c)
	sendMessage(username, c)

}
