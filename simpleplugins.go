package bruxism

import (
	"fmt"
	"strings"
)

func discordInviteID(id string) string {
	id = strings.Replace(id, "://discordapp.com/invite/", "://discord.gg/", -1)
	id = strings.Replace(id, "https://discord.gg/", "", -1)
	id = strings.Replace(id, "http://discord.gg/", "", -1)
	return id
}

func directMessageInviteMessageFunc(bot *Bot, service Service, message Message) {
	if service.Name() == DiscordServiceName && !service.IsMe(message) && service.IsPrivate(message) {
		messageMessage := message.Message()
		id := discordInviteID(messageMessage)
		if id != messageMessage && strings.HasPrefix(messageMessage, "http") {
			if err := service.Join(id); err != nil {
				if err == ErrAlreadyJoined {
					service.PrivateMessage(message.UserID(), "I have already joined that server.")
					return
				}
				fmt.Printf("Error joining %v %v", service.Name(), err)
				return
			}
			service.PrivateMessage(message.UserID(), "I have joined that server.")
		}
	}
}

// NewDirectMessageInvitePlugin creates a new direct message invite plugin.
func NewDirectMessageInvitePlugin() Plugin {
	p := NewSimplePlugin("DirectMessageInvite")
	p.message = directMessageInviteMessageFunc
	return p
}
