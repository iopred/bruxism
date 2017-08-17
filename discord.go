package bruxism

import (
	"fmt"
	"io"
	"log"
	"regexp"

	"github.com/iopred/discordgo"
)

// The number of guilds supported by one shard.
const numGuildsPerShard = 2400

// DiscordServiceName is the service name for the Discord service.
const DiscordServiceName string = "Discord"

// DiscordMessage is a Message wrapper around discordgo.Message.
type DiscordMessage struct {
	Discord          *Discord
	DiscordgoMessage *discordgo.Message
	MessageType      MessageType
	Nick             *string
	Content          *string
}

// Channel returns the channel id for this message.
func (m *DiscordMessage) Channel() string {
	return m.DiscordgoMessage.ChannelID
}

// UserName returns the user name for this message.
func (m *DiscordMessage) UserName() string {
	me := m.DiscordgoMessage
	if me.Author == nil {
		return ""
	}

	if m.Nick == nil {
		n := m.Discord.NicknameForID(me.Author.ID, me.Author.Username, me.ChannelID)
		m.Nick = &n
	}
	return *m.Nick
}

// UserID returns the user id for this message.
func (m *DiscordMessage) UserID() string {
	if m.DiscordgoMessage.Author == nil {
		return ""
	}

	return m.DiscordgoMessage.Author.ID
}

// UserAvatar returns the avatar url for this message.
func (m *DiscordMessage) UserAvatar() string {
	if m.DiscordgoMessage.Author == nil {
		return ""
	}

	return discordgo.EndpointUserAvatar(m.DiscordgoMessage.Author.ID, m.DiscordgoMessage.Author.Avatar)
}

// Message returns the message content for this message.
func (m *DiscordMessage) Message() string {
	if m.Content == nil {
		c := m.DiscordgoMessage.ContentWithMentionsReplaced()
		c = m.Discord.replaceRoleNames(m.DiscordgoMessage, c)
		c = m.Discord.replaceChannelNames(m.DiscordgoMessage, c)

		m.Content = &c
	}
	return *m.Content
}

// RawMessage returns the raw message content for this message.
func (m *DiscordMessage) RawMessage() string {
	return m.DiscordgoMessage.Content
}

// MessageID returns the message ID for this message.
func (m *DiscordMessage) MessageID() string {
	return m.DiscordgoMessage.ID
}

// Type returns the type of message.
func (m *DiscordMessage) Type() MessageType {
	return m.MessageType
}

// Discord is a Service provider for Discord.
type Discord struct {
	args        []interface{}
	messageChan chan Message

	Shards int

	// The first session, used to send messages (and maintain backwards compatibility).
	Session             *discordgo.Session
	Sessions            []*discordgo.Session
	OwnerUserID         string
	ApplicationClientID string
}

// NewDiscord creates a new discord service.
func NewDiscord(args ...interface{}) *Discord {
	return &Discord{
		args:        args,
		messageChan: make(chan Message, 200),
	}
}

var channelIDRegex = regexp.MustCompile("<#[0-9]*>")

func (d *Discord) replaceChannelNames(message *discordgo.Message, content string) string {
	return channelIDRegex.ReplaceAllStringFunc(content, func(str string) string {
		c, err := d.Channel(str[2 : len(str)-1])
		if err != nil {
			return str
		}

		return "#" + c.Name
	})
}

var roleIDRegex = regexp.MustCompile("<@&[0-9]*>")

func (d *Discord) replaceRoleNames(message *discordgo.Message, content string) string {
	return roleIDRegex.ReplaceAllStringFunc(content, func(str string) string {
		roleID := str[3 : len(str)-1]

		c, err := d.Channel(message.ChannelID)
		if err != nil {
			return str
		}

		g, err := d.Guild(c.GuildID)
		if err != nil {
			return str
		}

		for _, r := range g.Roles {
			if r.ID == roleID {
				return "@" + r.Name
			}
		}

		return str
	})
}

func (d *Discord) onMessageCreate(s *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Content == "" || (message.Author != nil && message.Author.Bot) {
		return
	}

	d.messageChan <- &DiscordMessage{
		Discord:          d,
		DiscordgoMessage: message.Message,
		MessageType:      MessageTypeCreate,
	}
}

func (d *Discord) onMessageUpdate(s *discordgo.Session, message *discordgo.MessageUpdate) {
	if message.Content == "" || (message.Author != nil && message.Author.Bot) {
		return
	}

	d.messageChan <- &DiscordMessage{
		Discord:          d,
		DiscordgoMessage: message.Message,
		MessageType:      MessageTypeUpdate,
	}
}

func (d *Discord) onMessageDelete(s *discordgo.Session, message *discordgo.MessageDelete) {
	d.messageChan <- &DiscordMessage{
		Discord:          d,
		DiscordgoMessage: message.Message,
		MessageType:      MessageTypeDelete,
	}
}

// Name returns the name of the service.
func (d *Discord) Name() string {
	return DiscordServiceName
}

// Open opens the service and returns a channel which all messages will be sent on.
func (d *Discord) Open() (<-chan Message, error) {
	shards := d.Shards
	if shards < 1 {
		shards = 1
	}

	d.Sessions = make([]*discordgo.Session, shards)

	for i := 0; i < shards; i++ {
		session, err := discordgo.New(d.args...)
		if err != nil {
			return nil, err
		}
		session.ShardCount = shards
		session.ShardID = i
		session.AddHandler(d.onMessageCreate)
		session.AddHandler(d.onMessageUpdate)
		session.AddHandler(d.onMessageDelete)
		session.State.TrackPresences = false

		d.Sessions[i] = session
	}

	d.Session = d.Sessions[0]

	for i := 0; i < len(d.Sessions); i++ {
		d.Sessions[i].Open()
	}

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
	if channel == "" {
		log.Println("Empty channel could not send message", message)
		return nil
	}

	if _, err := d.Session.ChannelMessageSend(channel, message); err != nil {
		log.Println("Error sending discord message: ", err)
		return err
	}

	return nil
}

// SendAction sends an action.
func (d *Discord) SendAction(channel, message string) error {
	if channel == "" {
		log.Println("Empty channel could not send message", message)
		return nil
	}

	p, err := d.UserChannelPermissions(d.UserID(), channel)
	if err != nil {
		return d.SendMessage(channel, message)
	}

	if p&discordgo.PermissionEmbedLinks == discordgo.PermissionEmbedLinks {
		if _, err := d.Session.ChannelMessageSendEmbed(channel, &discordgo.MessageEmbed{
			Color:       d.UserColor(d.UserID(), channel),
			Description: message,
		}); err != nil {
			return err
		}
		return nil
	}

	return d.SendMessage(channel, message)
}

// DeleteMessage deletes a message.
func (d *Discord) DeleteMessage(channel, messageID string) error {
	return d.Session.ChannelMessageDelete(channel, messageID)
}

// SendFile sends a file.
func (d *Discord) SendFile(channel, name string, r io.Reader) error {
	if _, err := d.Session.ChannelFileSend(channel, name, r); err != nil {
		log.Println("Error sending discord message: ", err)
		return err
	}
	return nil
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

// UserID returns the bots user id.
func (d *Discord) UserID() string {
	if d.Session.State.User == nil {
		return ""
	}
	return d.Session.State.User.ID
}

// Join accept an invite or return an error.
// If AlreadyJoinedError is return, @me has already accepted that invite.
func (d *Discord) Join(join string) error {
	if i, err := d.Session.Invite(join); err == nil {
		if _, err := d.Guild(i.Guild.ID); err == nil {
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

// SupportsPrivateMessages returns whether the service supports private messages.
func (d *Discord) SupportsPrivateMessages() bool {
	return true
}

// SupportsMultiline returns whether the service supports multiline messages.
func (d *Discord) SupportsMultiline() bool {
	return true
}

// CommandPrefix returns the command prefix for the service.
func (d *Discord) CommandPrefix() string {
	return fmt.Sprintf("@%s ", d.UserName())
}

// IsBotOwner returns whether or not a message sender was the owner of the bot.
func (d *Discord) IsBotOwner(message Message) bool {
	return message.UserID() == d.OwnerUserID
}

// IsPrivate returns whether or not a message was private.
func (d *Discord) IsPrivate(message Message) bool {
	c, err := d.Channel(message.Channel())
	return err == nil && c.IsPrivate
}

// IsChannelOwner returns whether or not the sender of a message is a moderator.
func (d *Discord) IsChannelOwner(message Message) bool {
	c, err := d.Channel(message.Channel())
	if err != nil {
		return false
	}
	g, err := d.Guild(c.GuildID)
	if err != nil {
		return false
	}
	return g.OwnerID == message.UserID() || d.IsBotOwner(message)
}

// IsModerator returns whether or not the sender of a message is a moderator.
func (d *Discord) IsModerator(message Message) bool {
	p, err := d.UserChannelPermissions(message.UserID(), message.Channel())
	if err == nil {
		if p&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator || p&discordgo.PermissionManageChannels == discordgo.PermissionManageChannels || p&discordgo.PermissionManageServer == discordgo.PermissionManageServer {
			return true
		}
	}

	return d.IsChannelOwner(message)
}

// ChannelCount returns the number of channels the bot is in.
func (d *Discord) ChannelCount() int {
	return len(d.Guilds())
}

// SupportsMessageHistory returns if the service supports message history.
func (d *Discord) SupportsMessageHistory() bool {
	return true
}

// MessageHistory returns the message history for a channel.
func (d *Discord) MessageHistory(channel string) []Message {
	c, err := d.Channel(channel)
	if err != nil {
		return nil
	}

	messages := make([]Message, len(c.Messages))
	for i := 0; i < len(c.Messages); i++ {
		messages[i] = &DiscordMessage{
			Discord:          d,
			DiscordgoMessage: c.Messages[i],
			MessageType:      MessageTypeCreate,
		}
	}

	return messages
}

func (d *Discord) Channel(channelID string) (channel *discordgo.Channel, err error) {
	for _, s := range d.Sessions {
		channel, err = s.State.Channel(channelID)
		if err == nil {
			return channel, nil
		}
	}
	return
}

func (d *Discord) Guild(guildID string) (guild *discordgo.Guild, err error) {
	for _, s := range d.Sessions {
		guild, err = s.State.Guild(guildID)
		if err == nil {
			return guild, nil
		}
	}
	return
}

func (d *Discord) Guilds() []*discordgo.Guild {
	guilds := []*discordgo.Guild{}
	for _, s := range d.Sessions {
		guilds = append(guilds, s.State.Guilds...)
	}
	return guilds
}

func (d *Discord) UserChannelPermissions(userID, channelID string) (apermissions int, err error) {
	for _, s := range d.Sessions {
		apermissions, err = s.State.UserChannelPermissions(userID, channelID)
		if err == nil {
			return apermissions, nil
		}
	}
	return
}

func (d *Discord) UserColor(userID, channelID string) int {
	for _, s := range d.Sessions {
		color := s.State.UserColor(userID, channelID)
		if color != 0 {
			return color
		}
	}
	return 0
}

func (d *Discord) Nickname(message Message) string {
	return d.NicknameForID(message.UserID(), message.UserName(), message.Channel())
}

func (d *Discord) NicknameForID(userID, userName, channelID string) string {
	c, err := d.Channel(channelID)
	if err == nil {
		g, err := d.Guild(c.GuildID)
		if err == nil {
			for _, m := range g.Members {
				if m.User.ID == userID {
					if m.Nick != "" {
						return m.Nick
					}
					break
				}
			}
		}
	}
	return userName
}
