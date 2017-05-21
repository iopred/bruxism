package inviteplugin

import (
	"fmt"
	"log"
	"strings"

	"github.com/iopred/bruxism"
)

func discordInviteID(id string) string {
	id = strings.Replace(id, "://discordapp.com/invite/", "://discord.gg/", -1)
	id = strings.Replace(id, "https://discord.gg/", "", -1)
	id = strings.Replace(id, "http://discord.gg/", "", -1)
	return id
}

// InviteHelp will return the help text for the invite command.
func InviteHelp(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) (string, string) {
	switch service.Name() {
	case bruxism.DiscordServiceName:
		discord := service.(*bruxism.Discord)

		if discord.ApplicationClientID != "" {
			return "", fmt.Sprintf("Returns a URL to add %s to your server.", service.UserName())
		}
		return "<discordinvite>", "Joins the provided Discord server."
	case bruxism.YouTubeServiceName:
		return "<videoid>", "Joins the provided YouTube live stream."
	}
	return "<channel>", "Joins the provided channel."
}

// InviteCommand is a command for accepting an invite to a channel.
func InviteCommand(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, command string, parts []string) {
	if service.Name() == bruxism.DiscordServiceName {
		discord := service.(*bruxism.Discord)

		if discord.ApplicationClientID != "" {
			service.SendMessage(message.Channel(), fmt.Sprintf("Please visit <https://discordapp.com/oauth2/authorize?client_id=%s&scope=bot> to add %s to your server.", discord.ApplicationClientID, service.UserName()))
			return
		}
	}

	if len(parts) == 1 {
		join := parts[0]
		if service.Name() == bruxism.DiscordServiceName {
			join = discordInviteID(join)
		}
		if err := service.Join(join); err != nil {
			if service.Name() == bruxism.DiscordServiceName && err == bruxism.ErrAlreadyJoined {
				service.PrivateMessage(message.UserID(), "I have already joined that server.")
				return
			}
			log.Println("Error joining %s %v", service.Name(), err)
		} else if service.Name() == bruxism.DiscordServiceName {
			service.PrivateMessage(message.UserID(), "I have joined that server.")
		}
	}
}
