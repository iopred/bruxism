package main

import (
  "errors"
  "log"
  "strconv"

  "github.com/Xackery/discord"
)

const DiscordServiceName string = "Discord"

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

func (m *DiscordMessage) IsModerator() bool {
  return false
}

type Discord struct {
  email       string
  password    string
  Client      *discord.Client
  messageChan chan Message
}

func NewDiscord(email, password string) *Discord {
  return &Discord{
    email:       email,
    password:    password,
    Client:      &discord.Client{},
    messageChan: make(chan Message, 200),
  }
}

func (d *Discord) onMessage(event discord.Event, message discord.Message) {
  dm := DiscordMessage(message)
  d.messageChan <- &dm
}

func (d *Discord) Name() string {
  return DiscordServiceName
}

func (d *Discord) Open() (<-chan Message, error) {

  if err := d.Client.Login(d.email, d.password); err != nil {
    return nil, err
  }

  d.Client.OnMessageCreate = d.onMessage

  go func() {
    d.Client.Listen()
  }()

  return d.messageChan, nil
}

func (d *Discord) IsMe(message Message) bool {
  return message.UserId() == strconv.Itoa(d.Client.User.ID)
}

func (d *Discord) SendMessage(channel, message string) error {
  id, err := strconv.Atoi(channel)
  if err != nil {
    log.Println("Error converting channel id: ", err)
    return err
  }

  _, err = d.Client.ChannelMessageSend(id, message)
  if err != nil {
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
