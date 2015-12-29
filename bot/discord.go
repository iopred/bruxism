package bot

import (
  "errors"
  "log"

  "github.com/bwmarrin/discordgo"
)

const DiscordServiceName string = "Discord"

type DiscordMessage discordgo.Message

func (m *DiscordMessage) Channel() string {
  return m.ChannelID
}

func (m *DiscordMessage) UserName() string {
  return m.Author.Username
}

func (m *DiscordMessage) UserId() string {
  return m.Author.ID
}

func (m *DiscordMessage) Message() string {
  return m.Content
}

func (m *DiscordMessage) MessageId() string {
  return m.ID
}

func (m *DiscordMessage) IsModerator() bool {
  return false
}

type Discord struct {
  email       string
  password    string
  Session     *discordgo.Session
  messageChan chan Message
  Me          *discordgo.User
}

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

func (d *Discord) Name() string {
  return DiscordServiceName
}

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

func (d *Discord) IsMe(message Message) bool {
  return message.UserId() == d.Me.ID
}

func (d *Discord) SendMessage(channel, message string) error {
  if _, err := d.Session.ChannelMessageSend(channel, message); err != nil {
    log.Println("Error sending discord message: ", err)
    return err
  }

  return nil
}

func (d *Discord) DeleteMessage(messageId string) error {
  return errors.New("Deleting not supported on Discord.")
}

func (d *Discord) BanUser(channel, user string, duration int) error {
  return errors.New("Banning not supported on Discord.")
}

func (d *Discord) UserName() string {
  return d.Me.Username
}

func (d *Discord) SetPlaying(game string) error {
  return d.Session.UpdateStatus(0, game)
}
