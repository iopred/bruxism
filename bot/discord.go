package bot

import (
	"errors"
	"log"

	"github.com/iopred/discordgo"
)

const DiscordServiceName string = "Discord"

type DiscordMessage discordgo.Message

// Channel returns the channel id for this message.
func (m *DiscordMessage) Channel() string {
	return m.ChannelID
}

// UserName returns the user name for this message.
func (m *DiscordMessage) UserName() string {
	return m.Author.Username
}

// UserID returns the user id for this message.
func (m *DiscordMessage) UserID() string {
	return m.Author.ID
}

// Message returns the message content for this message.
func (m *DiscordMessage) Message() string {
	return m.Content
}

// MessageID returns the message ID for this message.
func (m *DiscordMessage) MessageID() string {
	return m.ID
}

// IsModerator returns whether or not the sender of this message is a moderator.
func (m *DiscordMessage) IsModerator() bool {
	return false
}

// Discord is a Service provider for Discord.
type Discord struct {
	email       string
	password    string
	Session     *discordgo.Session
	messageChan chan Message
	Me          *discordgo.User
}

// NewDiscord creates a new discord service.
func NewDiscord(email, password string) *Discord {
	return &Discord{
		email:       email,
		password:    password,
		messageChan: make(chan Message, 200),
	}
}

func (d *Discord) onMessage(s *discordgo.Session, message discordgo.Message) {
	dm := DiscordMessage(message)
	d.messageChan <- &dm
}

// Name returns the name of the service.
func (d *Discord) Name() string {
	return DiscordServiceName
}

// Open opens the service and returns a channel which all messages will be sent on.
func (d *Discord) Open() (<-chan Message, error) {
	var err error

	d.Session = &discordgo.Session{
		OnMessageCreate: d.onMessage,
	}

	d.Session.Token, err = d.Session.Login(d.email, d.password)
	if err != nil {
		return nil, err
	}

	err = d.Session.Open()
	if err != nil {
		return nil, err
	}

	err = d.Session.Handshake()
	if err != nil {
		return nil, err
	}

	u, err := d.Session.User("@me")
	if err != nil {
		return nil, err
	}

	d.Me = &u

	// Listen for events.
	go d.Session.Listen()

	return d.messageChan, nil
}

// IsMe returns whether or not a message was sent by the bot.
func (d *Discord) IsMe(message Message) bool {
	return message.UserID() == d.Me.ID
}

// SendMessage sends a message.
func (d *Discord) SendMessage(channel, message string) error {
	if _, err := d.Session.ChannelMessageSend(channel, message); err != nil {
		log.Println("Error sending discord message: ", err)
		return err
	}
	return nil
}

// DeleteMessage deletes a message.
func (d *Discord) DeleteMessage(messageID string) error {
	return errors.New("Deleting not supported on Discord.")
}

// BanUser bans a user.
func (d *Discord) BanUser(channel, user string, duration int) error {
	return errors.New("Banning not supported on Discord.")
}

// UserName returns the bots name.
func (d *Discord) UserName() string {
	return d.Me.Username
}

// SetPlaying will set the current game being played by the bot.
func (d *Discord) SetPlaying(game string) error {
	return d.Session.UpdateStatus(0, game)
}

// Join will join a channel.
func (d *Discord) Join(join string) error {
	if _, err := d.Session.InviteAccept(join); err != nil {
		return err
	}
	return nil
}
