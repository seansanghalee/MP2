package message

import "fmt"

type Message struct {
	To, From, Content string
}

func Construct(to, from, content string) Message {
	message := Message{to, from, content}
	return message
}

func Display(message Message) {
	fmt.Println(message.From, ":", message.Content)
}
