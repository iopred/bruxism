package bruxism

import (
	"errors"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
)

// The number of guilds supported by one shard.
const numGuildsPerShard = 2400

// DiscordServiceName is the service name for the Discord service.
const DiscordServiceName string = "Discord"

type ParsedCommand struct {
	Matches bool
	Command string
	Rest    string
	Parts   []string
}

// DiscordMessage is a Message wrapper around discordgo.Message.
type DiscordMessage struct {
	sync.RWMutex
	Discord          *Discord
	DiscordgoMessage *discordgo.Message
	MessageType      MessageType
	Nick             *string
	Content          *string
	parsedCommands   map[string]*ParsedCommand
}

// Channel returns the channel id for this message.
func (m *DiscordMessage) Channel() string {
	return m.DiscordgoMessage.ChannelID
}

// UserName returns the user name for this message.
func (m *DiscordMessage) UserName() string {
	if m.Nick == nil {
		n := m.Discord.Nickname(m)
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

		if c == "" {
			c = m.AttachmentURL()
		}

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

func (m *DiscordMessage) AttachmentURL() string {
	if len(m.DiscordgoMessage.Attachments) == 0 {
		return ""
	}
	a := m.DiscordgoMessage.Attachments[0]
	if a.ProxyURL != "" {
		if a.Width != 0 && a.Height != 0 {
			height := 100
			width := height * a.Width / a.Height
			if width > 200 {
				width = 200
				height = width * a.Height / a.Width
			}
			return fmt.Sprintf("%s?width=%d&height=%d", a.ProxyURL, width, height)
		}
		return a.ProxyURL
	} else if a.URL != "" {
		return a.URL
	}
	return ""
}

func (m *DiscordMessage) AttachmentURLLarge() string {
	if len(m.DiscordgoMessage.Attachments) == 0 {
		return ""
	}
	a := m.DiscordgoMessage.Attachments[0]
	if a.ProxyURL != "" {
		return a.ProxyURL
	} else if a.URL != "" {
		return a.URL
	}
	return ""
}

func (m *DiscordMessage) getParsedCommand(prefix string) (parsedCommand *ParsedCommand) {
	prefix = strings.ToLower(prefix)

	m.Lock()
	defer m.Unlock()

	if m.parsedCommands == nil {
		m.parsedCommands = map[string]*ParsedCommand{}
	} else {
		parsedCommand = m.parsedCommands[prefix]
		if parsedCommand != nil {
			return parsedCommand
		}
	}

	parsedCommand = &ParsedCommand{}
	m.parsedCommands[prefix] = parsedCommand

	trimmedPrefix := strings.ToLower(strings.TrimSpace(prefix))
	messageLower := strings.ToLower(strings.TrimSpace(m.Message()))
	if !strings.HasPrefix(messageLower, trimmedPrefix) {
		return
	}
	parsedCommand.Matches = true

	parts := strings.Fields(m.Message())
	if len(parts) == 0 {
		return
	}

	if len(trimmedPrefix) != len(prefix) {
		if len(parts) > 1 {
			parts = parts[1:]
		}
	} else {
		parts[0] = parts[0][len(prefix):]
	}

	parsedCommand.Command = parts[0]
	if len(parts) > 1 {
		parsedCommand.Parts = parts[1:]
		parsedCommand.Rest = strings.Join(parsedCommand.Parts, " ")
	}
	return
}

func (m *DiscordMessage) MatchesCommand(prefix, command string) (bool, bool) {
	parsedCommand := m.getParsedCommand(prefix)
	return parsedCommand.Matches && (command == "" || parsedCommand.Command == strings.ToLower(command)), true
}

func (m *DiscordMessage) ParseCommand(prefix string) (string, []string, bool) {
	parsedCommand := m.getParsedCommand(prefix)
	return parsedCommand.Rest, parsedCommand.Parts, true
}

// Discord is a Service provider for Discord.
type Discord struct {
	token       string
	messageChan chan Message

	// The first session, used to send messages (and maintain backwards compatibility).
	Session             *discordgo.Session
	Sessions            []*discordgo.Session
	OwnerUserID         string
	ApplicationClientID string
}

// NewDiscord creates a new discord service.
// If the token is for a bot, it must be prefixed with "Bot "
// 		e.g. "Bot ..."
// Or if it is an OAuth2 token, it must be prefixed with "Bearer "
//		e.g. "Bearer ..."
func NewDiscord(token string) *Discord {
	return &Discord{
		token:       token,
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
	if message.Author != nil && message.Author.Bot {
		return
	}

	dm := &DiscordMessage{
		Discord:          d,
		DiscordgoMessage: message.Message,
		MessageType:      MessageTypeCreate,
	}
	if dm.Message() != "" {
		d.messageChan <- dm
	}
}

func (d *Discord) onMessageUpdate(s *discordgo.Session, message *discordgo.MessageUpdate) {
	if message.Author != nil && message.Author.Bot {
		return
	}

	dm := &DiscordMessage{
		Discord:          d,
		DiscordgoMessage: message.Message,
		MessageType:      MessageTypeUpdate,
	}
	if dm.Message() != "" {
		d.messageChan <- dm
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
	gateway, err := discordgo.New(d.token)
	if err != nil {
		return nil, err
	}

	s, err := gateway.GatewayBot()
	if err != nil {
		return nil, err
	}

	d.Sessions = make([]*discordgo.Session, s.Shards)

	log.Printf("%s opening with %d shards\n", d.Name(), s.Shards)
	wg := sync.WaitGroup{}
	for i := 0; i < s.Shards; i++ {
		log.Printf("%s opening shard %d\n", d.Name(), i+1)
		session, err := discordgo.New(d.token)
		if err != nil {
			return nil, err
		}
		if d.Session == nil {
			d.Session = session
		}
		d.Sessions[i] = session
		session.ShardCount = s.Shards
		session.ShardID = i
		session.State.TrackPresences = false
		wg.Add(1)
		go func(session *discordgo.Session) {
			defer wg.Done()
			err := session.Open()
			if err != nil {
				log.Printf("error opening shard %s", err)
			}
		}(d.Sessions[i])
	}
	wg.Wait()

	for _, session := range d.Sessions {
		session.AddHandler(d.onMessageCreate)
		session.AddHandler(d.onMessageUpdate)
		session.AddHandler(d.onMessageDelete)
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
	return err == nil && c.Type == discordgo.ChannelTypeDM
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
	isChannelOwner := d.IsChannelOwner(message)

	p, err := d.MessagePermissions(message)
	if err != nil {
		return isChannelOwner
	}

	return isChannelOwner || p&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator || p&discordgo.PermissionManageChannels == discordgo.PermissionManageChannels || p&discordgo.PermissionManageServer == discordgo.PermissionManageServer
}

// ChannelCount returns the number of channels the bot is in.
func (d *Discord) ChannelCount() int {
	return len(d.Guilds())
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

func (d *Discord) Member(guildID string, userID string) (member *discordgo.Member, err error) {
	for _, s := range d.Sessions {
		member, err = s.State.Member(guildID, userID)
		if err == nil {
			return member, nil
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

func (d *Discord) MessagePermissions(message Message) (apermissions int64, err error) {
	m, ok := message.(*DiscordMessage)
	if !ok {
		return 0, errors.New("invalid message")
	}
	for _, s := range d.Sessions {
		apermissions, err = s.State.MessagePermissions(m.DiscordgoMessage)
		if err == nil {
			return apermissions, nil
		}
	}
	return d.UserChannelPermissions(message.UserID(), message.Channel())
}

func (d *Discord) UserChannelPermissions(userID, channelID string) (apermissions int64, err error) {
	for _, s := range d.Sessions {
		apermissions, err = s.State.UserChannelPermissions(userID, channelID)
		if err == nil {
			return apermissions, nil
		}
	}
	return 0, errors.New("could not find user permissions")
}

func (d *Discord) MessageColor(message Message) int {
	m, ok := message.(*DiscordMessage)
	if !ok {
		return 0
	}
	for _, s := range d.Sessions {
		color := s.State.MessageColor(m.DiscordgoMessage)
		if color != 0 {
			return color
		}
	}
	return 0
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
	m, ok := message.(*DiscordMessage)
	if !ok {
		return ""
	}
	if m.DiscordgoMessage.Member != nil && m.DiscordgoMessage.Member.Nick != "" {
		return m.DiscordgoMessage.Member.Nick
	}
	if m.DiscordgoMessage.Author == nil {
		return ""
	}
	return d.NicknameForID(message.UserID(), m.DiscordgoMessage.Author.Username, message.Channel())
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
					return m.User.Username
				}
			}
		}
	}
	return userName
}

func (d *Discord) UpdateStatus(idle int, game string) {
	for _, s := range d.Sessions {
		s.UpdateGameStatus(idle, game)
	}
}
func (d *Discord) UpdateStreamingStatus(idle int, game string, url string) {
	for _, s := range d.Sessions {
		s.UpdateStreamingStatus(idle, game, url)
	}
}
