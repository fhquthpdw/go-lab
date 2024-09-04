package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	event, err := InitializeEvent("Hello world!")
	if err != nil {
		fmt.Println(err)
		return
	}
	event.Start()
}

type Message struct {
	Content string
	Code    int
}

func NewMessage(phrase string) Message {
	return Message(phrase)
}

type Greeter struct {
	Message Message
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter
}

func NewEvent(g Greeter) (Event, error) {
	if time.Now().Unix()%2 == 0 {
		return Event{}, errors.New("new event error")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

//func InitializeEvent2() Event {
//	message := NewMessage()
//	greeter := NewGreeter(message)
//	event := NewEvent(greeter)
//
//	return event
//}
