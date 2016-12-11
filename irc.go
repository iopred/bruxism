package bruxism

import (
	"errors"
	"io"

	"github.com/fluffle/goirc/client"
)

// IRCServiceName is the service name for the IRC service.
const IRCServiceName string = "IRC"

// IRCMessage is a Message wrapper around client.Line.
type IRCMessage client.Line

// Channel returns the channel id for this message.
func (m *IRCMessage) Channel() string {
	i := client.Line(*m)
	return i.Target()
}

// UserName returns the user name for this message.
func (m *IRCMessage) UserName() string {
	return m.Nick
}

// UserID returns the user id for this message.
func (m *IRCMessage) UserID() string {
	return m.Nick
}

// UserAvatar returns the avatar url for this message.
func (m *IRCMessage) UserAvatar() string {
	return ""
}

// Message returns the message content for this message.
func (m *IRCMessage) Message() string {
	return m.RawMessage()
}

// RawMessage returns the raw message content for this message.
func (m *IRCMessage) RawMessage() string {
	i := client.Line(*m)
	return i.Text()
}

// MessageID returns the message ID for this message.
func (m *IRCMessage) MessageID() string {
	return ""
}

// Type returns the type of message.
func (m *IRCMessage) Type() MessageType {
	return MessageTypeCreate
}

// IRC is a Service provider for IRC.
type IRC struct {
	host        string
	nick        string
	password    string
	channels    []string
	Conn        *client.Conn
	messageChan chan Message
}

// NewIRC creates a new IRC service.
func NewIRC(host, nick, password string, channels []string) *IRC {
	return &IRC{
		host:        host,
		nick:        nick,
		password:    password,
		channels:    channels,
		messageChan: make(chan Message, 200),
	}
}

func (i *IRC) onMessage(conn *client.Conn, line *client.Line) {
	m := IRCMessage(*line)
	i.messageChan <- &m
}

func (i *IRC) onConnect(conn *client.Conn, line *client.Line) {
	for _, c := range i.channels {
		conn.Join(c)
	}
}

func (i *IRC) onDisconnect(conn *client.Conn, line *client.Line) {
	conn.ConnectTo(i.host)
}

// Name returns the name of the service.
func (i *IRC) Name() string {
	return IRCServiceName
}

// Open opens the service and returns a channel which all messages will be sent on.
func (i *IRC) Open() (<-chan Message, error) {
	i.Conn = client.SimpleClient(i.nick, i.nick, i.nick)
	i.Conn.Config().Version = i.nick
	i.Conn.Config().QuitMessage = ""

	i.Conn.HandleFunc("connected", i.onConnect)
	i.Conn.HandleFunc("disconnected", i.onDisconnect)
	i.Conn.HandleFunc(client.PRIVMSG, i.onMessage)

	go i.Conn.ConnectTo(i.host, i.password)

	return i.messageChan, nil
}

// IsMe returns whether or not a message was sent by the bot.
func (i *IRC) IsMe(message Message) bool {
	return message.UserName() == i.UserName()
}

// SendMessage sends a message.
func (i *IRC) SendMessage(channel, message string) error {
	i.Conn.Privmsg(channel, message)
	return nil
}

// SendAction sends an action.
func (i *IRC) SendAction(channel, message string) error {
	i.Conn.Action(channel, message)
	return nil
}

// DeleteMessage deletes a message.
func (i *IRC) DeleteMessage(channel, messageID string) error {
	return errors.New("Deleting messages not supported on IRC.")
}

// SendFile sends a file.
func (i *IRC) SendFile(channel, name string, r io.Reader) error {
	return errors.New("Send file not supported.")
}

// BanUser bans a user.
func (i *IRC) BanUser(channel, userID string, duration int) error {
	i.Conn.Kick(channel, userID)
	return nil
}

// UnbanUser unbans a user.
func (i *IRC) UnbanUser(channel, userID string) error {
	return errors.New("Unbanning users not supported on IRC.")
}

// UserName returns the bots name.
func (i *IRC) UserName() string {
	return i.Conn.Me().Nick
}

// UserID returns the bots user id.
func (i *IRC) UserID() string {
	return i.Conn.Me().Nick
}

// Join will join a channel.
func (i *IRC) Join(join string) error {
	i.Conn.Join(join)
	return nil
}

// Typing sets that the bot is typing.
func (i *IRC) Typing(channel string) error {
	return errors.New("Typing not supported on IRC.")
}

// PrivateMessage will send a private message to a user.
func (i *IRC) PrivateMessage(userID, message string) error {
	return i.SendMessage(userID, message)
}

// SupportsPrivateMessages returns whether the service supports private messages.
func (i *IRC) SupportsPrivateMessages() bool {
	return true
}

// SupportsMultiline returns whether the service supports multiline messages.
func (i *IRC) SupportsMultiline() bool {
	return false
}

// CommandPrefix returns the command prefix for the service.
func (i *IRC) CommandPrefix() string {
	return "!"
}

// IsBotOwner returns whether or not a message sender was the owner of the bot.
func (i *IRC) IsBotOwner(message Message) bool {
	return false
}

// IsPrivate returns whether or not a message was private.
func (i *IRC) IsPrivate(message Message) bool {
	return message.UserName() == message.Channel()
}

// IsChannelOwner returns whether or not the sender of a message is the owner.
func (i *IRC) IsChannelOwner(message Message) bool {
	return false
}

// IsModerator returns whether or not the sender of a message is a moderator.
func (i *IRC) IsModerator(message Message) bool {
	return false
}

// ChannelCount returns the number of channels the bot is in.
func (i *IRC) ChannelCount() int {
	return len(i.channels)
}

// SupportsMessageHistory returns if the service supports message history.
func (i *IRC) SupportsMessageHistory() bool {
	return false
}

// MessageHistory returns the message history for a channel.
func (i *IRC) MessageHistory(channel string) []Message {
	return nil
}
