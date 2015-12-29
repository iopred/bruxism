package bot

import (
	"log"
	"sort"
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
