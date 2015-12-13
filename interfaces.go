package main

type Message interface {
  Channel() string
  UserName() string
  UserId() string
  Message() string
  MessageId() string
}

type Service interface {
  Name() string
  MessageChannel() <-chan Message
  Open() error
  SendMessage(channel, message string) error
  DeleteMessage(messageId string) error
  BanUser(channel, user string, duration int) error
}
