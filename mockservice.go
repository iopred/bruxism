package bruxism

import (
	"errors"
	"io"
)

// Service is a service interface, wraps a single service such as YouTube or Discord.
type MockService struct {
	name, username, userid string
	messages               chan Message
}

func NewMockService() *MockService {
	return &MockService{
		messages: make(chan Message, 512),
	}
}

func (m *MockService) SetName(name string) *MockService {
	m.name = name
	return m
}

func (m *MockService) SetUserName(username string) *MockService {
	m.username = username
	return m
}

func (m *MockService) SetUserID(userid string) *MockService {
	m.userid = userid
	return m
}

// Satisfy Service interface below

func (s *MockService) Name() string {
	return s.name
}

func (s *MockService) UserName() string {
	return s.username
}

func (s *MockService) UserID() string {
	return s.userid
}

func (s *MockService) Open() (<-chan Message, error) {
	return s.messages, nil
}

func (s *MockService) NewMessage(message Message) {
	s.messages <- message
}

func (s *MockService) SendFile(channel, name string, r io.Reader) error {
	return errors.New("Not implemented")
}

func (s *MockService) IsMe(message Message) bool {
	if message.UserID() == s.UserID() {
		return true
	}
	return false
}

func (s *MockService) SendMessage(channel, message string) error {
	return nil
}

func (s *MockService) SendAction(channel, message string) error {
	return nil
}

func (s *MockService) DeleteMessage(channel, messageID string) error {
	return nil
}

func (s *MockService) BanUser(channel, userID string, duration int) error {
	return nil
}

func (s *MockService) UnbanUser(channel, userID string) error {
	return nil
}

func (s *MockService) Join(join string) error {
	return nil
}

func (s *MockService) Typing(channel string) error {
	return nil
}

func (s *MockService) PrivateMessage(userID, messageID string) error {
	return nil
}

func (s *MockService) IsBotOwner(message Message) bool {
	return s.IsChannelOwner(message)
}

func (s *MockService) IsPrivate(message Message) bool {
	return false
}

func (s *MockService) IsChannelOwner(message Message) bool {
	if message.Channel() == message.UserID() {
		return true
	}
	return false
}

func (s *MockService) IsModerator(message Message) bool {
	return false
}

func (s *MockService) SupportsPrivateMessages() bool {
	return false
}

func (s *MockService) SupportsMultiline() bool {
	return false
}

func (s *MockService) CommandPrefix() string {
	return "!"
}

func (s *MockService) ChannelCount() int {
	return 1
}

var _ Service = &MockService{}
