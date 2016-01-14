package bruxism

import (
	"fmt"
	"log"
	"regexp"

	"github.com/iopred/discordgo"
)

// DiscordServiceName is the service name for the Discord service.
const DiscordServiceName string = "Discord"

// DiscordMessage is a Message wrapper around discordgo.Message.
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

// UserAvatar returns the avatar url for this message.
func (m *DiscordMessage) UserAvatar() string {
	return discordgo.USER_AVATAR(m.Author.ID, m.Author.Avatar)
}

// Message returns the message content for this message.
func (m *DiscordMessage) Message() string {
	d := discordgo.Message(*m)
	return d.ContentWithMentionsReplaced()
}

// Message returns the message content for this message.
func (m *DiscordMessage) RawMessage() string {
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
	args        []interface{}
	Session     *discordgo.Session
	messageChan chan Message
}

// NewDiscord creates a new discord service.
func NewDiscord(args ...interface{}) *Discord {
	return &Discord{
		args:        args,
		messageChan: make(chan Message, 200),
	}
}

var channelIDRegex = regexp.MustCompile("<#[0-9]*>")

func (d *Discord) onMessage(s *discordgo.Session, message *discordgo.Message) {
	if message.Content == "" {
		return
	}

	message.Content = channelIDRegex.ReplaceAllStringFunc(message.Content, func(str string) string {
		c, err := d.Session.State.Channel(str[2 : len(str)-1])
		if err != nil {
			return str
		}

		return "#" + c.Name
	})

	d.messageChan <- (*DiscordMessage)(message)
}

// Name returns the name of the service.
func (d *Discord) Name() string {
	return DiscordServiceName
}

// Open opens the service and returns a channel which all messages will be sent on.
func (d *Discord) Open() (<-chan Message, error) {
	var err error

	d.Session, err = discordgo.New(d.args...)
	if err != nil {
		return nil, err
	}
	d.Session.OnMessageCreate = d.onMessage

	return d.messageChan, nil
}

// IsMe returns whether or not a message was sent by the bot.
func (d *Discord) IsMe(message Message) bool {
	if d.Session.State.User == nil {
		return false
	}
	return message.UserID() == d.Session.State.User.ID
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
func (d *Discord) DeleteMessage(channel, messageID string) error {
	return d.Session.ChannelMessageDelete(channel, messageID)
}

// BanUser bans a user.
func (d *Discord) BanUser(channel, userID string, duration int) error {
	return d.Session.GuildBanCreate(channel, userID, 0)
}

// UnbanUser unbans a user.
func (d *Discord) UnbanUser(channel, userID string) error {
	return d.Session.GuildBanDelete(channel, userID)
}

// UserName returns the bots name.
func (d *Discord) UserName() string {
	if d.Session.State.User == nil {
		return ""
	}
	return d.Session.State.User.Username
}

// SetPlaying will set the current game being played by the bot.
func (d *Discord) SetPlaying(game string) error {
	return d.Session.UpdateStatus(0, game)
}

// Join accept an invite or return an error.
// If AlreadyJoinedError is return, @me has already accepted that invite.
func (d *Discord) Join(join string) error {
	if i, err := d.Session.Invite(join); err == nil {
		if _, err := d.Session.State.Guild(i.Guild.ID); err == nil {
			return ErrAlreadyJoined
		}
	}

	if _, err := d.Session.InviteAccept(join); err != nil {
		return err
	}
	return nil
}

// Typing sets that the bot is typing.
func (d *Discord) Typing(channel string) error {
	return d.Session.ChannelTyping(channel)
}

// PrivateMessage will send a private message to a user.
func (d *Discord) PrivateMessage(userID, message string) error {
	c, err := d.Session.UserChannelCreate(userID)
	if err != nil {
		return err
	}
	return d.SendMessage(c.ID, message)
}

// SupportsMultiline returns whether the service supports multiline messages.
func (d *Discord) SupportsMultiline() bool {
	return true
}

// CommandPrefix returns the command prefix for the service.
func (d *Discord) CommandPrefix() string {
	return fmt.Sprintf("@%s ", d.UserName())
}

// IsPrivate returns whether or not a message was private.
func (d *Discord) IsPrivate(message Message) bool {
	_, err := d.Session.State.PrivateChannel(message.Channel())
	return err == nil
}
