package main

import (
  "errors"
  "flag"
  "strconv"

  "github.com/Xackery/discord"
)

var discordEmail string
var discordPassword string

func init() {
  flag.StringVar(&discordEmail, "discordemail", "", "Discord account email.")
  flag.StringVar(&discordPassword, "discordpassword", "", "Discord account password.")
  flag.Parse()
}

type DiscordMessage discord.Message

func (m *DiscordMessage) Channel() string {
  return strconv.Itoa(m.ChannelID)
}

func (m *DiscordMessage) UserName() string {
  return m.Author.Username
}

func (m *DiscordMessage) UserId() string {
  return strconv.Itoa(m.Author.ID)
}

func (m *DiscordMessage) Message() string {
  return m.Content
}

func (m *DiscordMessage) MessageId() string {
  return strconv.Itoa(m.ID)
}

type Discord struct {
  Client      *discord.Client
  MessageChan chan Message
}

func NewDiscord() *Discord {
  return &Discord{
    Client:      &discord.Client{},
    MessageChan: make(chan Message, 200),
  }
}

func (d *Discord) onMessage(event discord.Event, message discord.Message) {
  // Ignore messages from ourselves.
  if message.Author.ID != d.Client.User.ID {
    dm := DiscordMessage(message)
    d.MessageChan <- &dm
  }
}

func (d *Discord) Name() string {
  return "Discord"
}

func (d *Discord) MessageChannel() <-chan Message {
  return d.MessageChan
}

func (d *Discord) Open() error {

  if err := d.Client.Login(discordEmail, discordPassword); err != nil {
    return err
  }

  d.Client.OnMessageCreate = d.onMessage

  go func() {
    d.Client.Listen()
  }()

  return nil
}

func (d *Discord) SendMessage(channel, message string) error {
  id, err := strconv.Atoi(channel)
  if err != nil {
    return err
  }

  _, err = d.Client.ChannelMessageSend(id, message)
  if err != nil {
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
