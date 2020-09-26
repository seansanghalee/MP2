// Package message provides primitives for constructing struct type Message
// and displaying it to a user-friendly format
package message

import "fmt"

type Message struct {
	To, From, Content string
}

// Construct constructs a type message struct
func Construct(to, from, content string) Message {
	message := Message{to, from, content}
	return message
}

// Display prints out the message in a readable format
func Display(message Message) {
	fmt.Println(message.From, ":", message.Content)
}
