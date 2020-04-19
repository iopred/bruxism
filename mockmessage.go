package bruxism

type MockMessage struct {
	id, channel, username, userid, avatar, message, rawMessage string
}

func NewMockMessage() *MockMessage {
	return &MockMessage{}
}

// Functional options to set various properties
func (m *MockMessage) SetChannel(channel string) *MockMessage {
	m.channel = channel
	return m
}

func (m *MockMessage) SetUserName(username string) *MockMessage {
	m.username = username
	return m
}

func (m *MockMessage) SetUserID(userid string) *MockMessage {
	m.userid = userid
	return m
}

func (m *MockMessage) SetAvatar(avatar string) *MockMessage {
	m.avatar = avatar
	return m
}

func (m *MockMessage) SetMessage(message string) *MockMessage {
	m.message = message
	return m
}

func (m *MockMessage) SetRawMessage(message string) *MockMessage {
	m.rawMessage = message
	return m
}

func (m *MockMessage) SetMessageID(messageID string) *MockMessage {
	m.id = messageID
	return m
}

// Satisfy Message interface below
func (m *MockMessage) Channel() string {
	return m.channel
}

func (m *MockMessage) UserName() string {
	return m.username
}

func (m *MockMessage) UserID() string {
	return m.userid
}

func (m *MockMessage) UserAvatar() string {
	return m.avatar
}

func (m *MockMessage) Message() string {
	return m.message
}

func (m *MockMessage) RawMessage() string {
	return m.rawMessage
}

func (m *MockMessage) MessageID() string {
	return m.id
}

func (m *MockMessage) Type() MessageType {
	return MessageTypeCreate
}

func (m *MockMessage) MatchesCommand(prefix, command string) (bool, bool) {
	return false, false
}

func (m *MockMessage) ParseCommand(prefix string) (string, []string, bool) {
	return "", []string{}, false
}

var _ Message = &MockMessage{}