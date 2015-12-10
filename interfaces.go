package main

type Message interface {
  Channel() string
  User() string
  Message() string
}

type Service interface {
  Name() string
  MessageChannel() chan Message
  Open() error
  Send(channel, message string) error
}
