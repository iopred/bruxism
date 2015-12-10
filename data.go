package main

type Message interface {
	Channel() string
	User() string
	Message() string
}

type Service interface {
	Name() string
	Open() (chan Message, error)
	Send(channel, message string) error
}
