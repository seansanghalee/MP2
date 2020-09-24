package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func login() (string, string, string) {
	var port, ip, username string

	fmt.Print("Enter the host address: ")
	fmt.Scanln(&ip)
	fmt.Print("Enter the port number: ")
	fmt.Scanln(&port)
	fmt.Print("Enter your username: ")
	fmt.Scanln(&username)
	return ip, port, username
}

func sendUsername(username string, conn net.Conn) {
	enc := gob.NewEncoder(conn)
	enc.Encode(username)
	return
}

//func sendMessage() {
//
//}

func main() {

	ip, port, username := login()

	address := ip + ":" + port
	c, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	sendUsername(username, c)

	//for {
	//	sendMessage()
	//}

}
