package bot

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

// HelpCommand is a command for returning help text for all registered plugins on a service.
func HelpCommand(bot *Bot, service Service, message Message) {
	help := []string{}

	for _, plugin := range bot.Services[service.Name()].Plugins {
		h := plugin.Help(bot, service)
		if h != nil && len(h) > 0 {
			help = append(help, h...)
		}
	}

	sort.Strings(help)

	if service.SupportsMultiline() {
		if err := service.SendMessage(message.Channel(), strings.Join(help, "\n")); err != nil {
			log.Println(err)
		}
	} else {
		for _, h := range help {
			if err := service.SendMessage(message.Channel(), h); err != nil {
				log.Println(err)
			}
		}
	}

}

// InviteHelp will return the help text for the invite command.
func InviteHelp(bot *Bot, service Service) (string, string) {
	switch service.Name() {
	case DiscordServiceName:
		return "<discordinvite>", "Joins the provided Discord server."
	case YouTubeServiceName:
		return "<livechatid>", "Joins the provided YouTube chat by id (this may be hard to find)."
	}
	return "<channel>", "Joins the provided channel."
}

// InviteCommand is a command for accepting an invite to a channel.
func InviteCommand(bot *Bot, service Service, message Message) {
	_, parts := parseCommand(service, message)
	if len(parts) == 1 {
		join := parts[0]
		if service.Name() == DiscordServiceName {
			join = strings.Replace(join, "https://discord.gg/", "", -1)
			join = strings.Replace(join, "http://discord.gg/", "", -1)
		}
		if err := service.Join(join); err != nil {
			fmt.Printf("Error joining %v %v", service.Name(), err)
		} else if service.Name() == DiscordServiceName {
			service.PrivateMessage(message.UserID(), "I have joined that server.")
		}
	}
}
