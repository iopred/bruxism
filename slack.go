package bruxism

import (
	"errors"
	"fmt"
	"io"

	"github.com/nlopes/slack"
)

// SlackServiceName is the service name for the Slack service.
const SlackServiceName string = "Slack"

// SlackMessage is a Message wrapper around slack.MessageEvent.
type SlackMessage struct {
	SlackMessage *slack.Msg
	MessageType  MessageType
}

// Channel returns the channel id for this message.
func (m *SlackMessage) Channel() string {
	return m.SlackMessage.Channel
}

// UserName returns the user name for this message.
func (m *SlackMessage) UserName() string {
	return m.SlackMessage.Username
}

// UserID returns the user id for this message.
func (m *SlackMessage) UserID() string {
	return m.SlackMessage.User
}

// UserAvatar returns the avatar url for this message.
func (m *SlackMessage) UserAvatar() string {
	return ""
}

// Message returns the message content for this message.
func (m *SlackMessage) Message() string {
	return m.SlackMessage.Text
}

// RawMessage returns the raw message content for this message.
func (m *SlackMessage) RawMessage() string {
	return m.SlackMessage.Text
}

// MessageID returns the message ID for this message.
func (m *SlackMessage) MessageID() string {
	return m.SlackMessage.Timestamp
}

// Type returns the type of message.
func (m *SlackMessage) Type() MessageType {
	return m.MessageType
}

// Slack is a Service provider for Slack.
type Slack struct {
	token       string
	messageChan chan Message

	Client *slack.Client
	RTM    *slack.RTM
	Me     *slack.AuthTestResponse

	OwnerUserID string
}

// NewSlack creates a new Slack service.
func NewSlack(token string) *Slack {
	return &Slack{
		token:       token,
		messageChan: make(chan Message, 200),
	}
}

func (s *Slack) handle() {
	for {
		select {
		case msg := <-s.RTM.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				switch ev.SubType {
				case "message_changed":
					s.messageChan <- &SlackMessage{ev.SubMessage, MessageTypeUpdate}
				case "message_deleted":
					ev.Msg.Timestamp = ev.Msg.DeletedTimestamp
					s.messageChan <- &SlackMessage{&ev.Msg, MessageTypeDelete}
				case "":
					s.messageChan <- &SlackMessage{&ev.Msg, MessageTypeCreate}
				}
			}
		}
	}
}

// Name returns the name of the service.
func (s *Slack) Name() string {
	return SlackServiceName
}

// Open opens the service and returns a channel which all messages will be sent on.
func (s *Slack) Open() (<-chan Message, error) {
	s.Client = slack.New(s.token)

	var err error
	s.Me, err = s.Client.AuthTest()
	if err != nil {
		return nil, err
	}

	s.RTM = s.Client.NewRTM()
	go s.RTM.ManageConnection()
	go s.handle()

	return s.messageChan, nil
}

// IsMe returns whether or not a message was sent by the bot.
func (s *Slack) IsMe(message Message) bool {
	return message.UserID() == s.Me.UserID
}

// SendMessage sends a message.
func (s *Slack) SendMessage(channel, message string) error {
	s.RTM.SendMessage(s.RTM.NewOutgoingMessage(message, channel))
	return nil
}

// SendAction sends an action.
func (s *Slack) SendAction(channel, message string) error {
	return s.SendMessage(channel, message)
}

// DeleteMessage deletes a message.
func (s *Slack) DeleteMessage(channel, messageID string) error {
	return errors.New("Slack does not support deleting messages.")
}

// SendFile sends a file.
func (s *Slack) SendFile(channel, name string, r io.Reader) error {
	return errors.New("Slack does not support sending files.")
}

// BanUser bans a user.
func (s *Slack) BanUser(channel, userID string, duration int) error {
	return errors.New("Slack does not support banning.")
}

// UnbanUser unbans a user.
func (s *Slack) UnbanUser(channel, userID string) error {
	return errors.New("Slack does not support unbanning.")
}

// UserName returns the bots name.
func (s *Slack) UserName() string {
	return s.Me.User
}

// UserID returns the bots user id.
func (s *Slack) UserID() string {
	return s.Me.UserID
}

// Join accept an invite or return an error.
func (s *Slack) Join(join string) error {
	_, err := s.Client.JoinChannel(join)
	return err
}

// Typing sets that the bot is typing.
func (s *Slack) Typing(channel string) error {
	s.RTM.SendMessage(s.RTM.NewTypingMessage(channel))
	return nil
}

// PrivateMessage will send a private message to a user.
func (s *Slack) PrivateMessage(userID, message string) error {
	s.RTM.SendMessage(s.RTM.NewOutgoingMessage(message, userID))
	return nil
}

// SupportsPrivateMessages returns whether the service supports private messages.
func (s *Slack) SupportsPrivateMessages() bool {
	return false
}

// SupportsMultiline returns whether the service supports multiline messages.
func (s *Slack) SupportsMultiline() bool {
	return true
}

// CommandPrefix returns the command prefix for the service.
func (s *Slack) CommandPrefix() string {
	return fmt.Sprintf("<@%s>: ", s.Me.UserID)
}

// IsBotOwner returns whether or not a message sender was the owner of the bot.
func (s *Slack) IsBotOwner(message Message) bool {
	return message.UserID() == s.OwnerUserID
}

// IsPrivate returns whether or not a message was private.
func (s *Slack) IsPrivate(message Message) bool {
	return false
}

// IsChannelOwner returns whether or not the sender of a message is the owner.
func (s *Slack) IsChannelOwner(message Message) bool {
	return false
}

// IsModerator returns whether or not the sender of a message is a moderator.
func (s *Slack) IsModerator(message Message) bool {
	return false
}

// ChannelCount returns the number of channels the bot is in.
func (s *Slack) ChannelCount() int {
	c, err := s.Client.GetChannels(true)
	if err != nil {
		return 0
	}
	return len(c)
}

// SupportsMessageHistory returns if the service supports message history.
func (s *Slack) SupportsMessageHistory() bool {
	return false
}

// MessageHistory returns the message history for a channel.
func (s *Slack) MessageHistory(channel string) []Message {
	return nil
}
