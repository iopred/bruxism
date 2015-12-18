package main

type Message interface {
  Channel() string
  UserName() string
  UserId() string
  Message() string
  MessageId() string
  IsModerator() bool
}

type SimpleMessage struct {
  channel string
  message string
}

func NewSimpleMessage(channel, message string) *SimpleMessage {
  return &SimpleMessage{
    channel: channel,
    message: message,
  }
}

func (s *SimpleMessage) Channel() string {
  return s.channel
}
func (s *SimpleMessage) UserName() string {
  return ""
}
func (s *SimpleMessage) UserId() string {
  return ""
}
func (s *SimpleMessage) Message() string {
  return s.message
}
func (s *SimpleMessage) MessageId() string {
  return ""
}
func (s *SimpleMessage) IsModerator() bool {
  return false
}

type Service interface {
  Name() string
  Open() (<-chan Message, error)
  IsMe(message Message) bool
  SendMessage(channel, message string) error
  DeleteMessage(messageId string) error
  BanUser(channel, user string, duration int) error
}

type Plugin interface {
  Name() string
  Help() string
  Register(*Bot, Service, []byte) error
  Save() []byte
}
