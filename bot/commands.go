package bot

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func HelpCommand(bot *Bot, service Service, message Message) {
	help := make([]string, 0)

	for _, plugin := range bot.Services[service.Name()].Plugins {
		h := plugin.Help(bot, service)
		if h != nil && len(h) > 0 {
			help = append(help, h...)
		}
	}

	sort.Strings(help)

	messageChannel := message.Channel()

	for _, h := range help {
		if err := service.SendMessage(messageChannel, h); err != nil {
			log.Println(err)
		}
	}
}

func InviteHelp(bot *Bot, service Service) (string, string) {
	switch service.Name() {
	case DiscordServiceName:
		return "<discordinvite>", "Joins the provided Discord server."
	case YouTubeServiceName:
		return "<livechatid>", "Joins the provided YouTube chat by id (this may be hard to find)."
	}
	return "<channel>", "Joins the provided channel."
}

func InviteCommand(bot *Bot, service Service, message Message) {
	_, parts := parseCommand(message)
	if len(parts) == 1 {
		join := parts[0]
		if service.Name() == DiscordServiceName {
			join = strings.Replace(join, "https://discord.gg/", "", -1)
			join = strings.Replace(join, "http://discord.gg/", "", -1)
		}
		if err := service.Join(join); err != nil {
			fmt.Printf("Error joining %v %v", service.Name(), err)
		}
	}
}
