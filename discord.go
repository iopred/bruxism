package main

import (
	"strconv"

	"github.com/Xackery/discord"
)

type DiscordMessage discord.Message

func (m *DiscordMessage) Channel() string {
	return strconv.Itoa(m.ChannelID)
}

func (m *DiscordMessage) User() string {
	return m.Author.Username
}

func (m *DiscordMessage) Message() string {
	return m.Content
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
	dm := DiscordMessage(message)
	if message.Author.ID != d.Client.User.ID {
		d.MessageChan <- &dm
	}
}

func (d *Discord) Name() string {
	return "Discord"
}

func (d *Discord) MessageChannel() chan Message {
	return d.MessageChan
}

func (d *Discord) Open() error {

	if err := d.Client.Login("iopred+bot@gmail.com", "dgmrthe1"); err != nil {
		return err
	}

	d.Client.OnMessageCreate = d.onMessage

	return d.Client.Listen()
}

func (d *Discord) Send(channel, message string) error {
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
