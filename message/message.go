package message

type Message struct {
	To, From, Content string
}

func Construct(to, from, content string) Message {
	message := Message{to, from, content}
	return message
}