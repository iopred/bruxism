package main

import (
	"log"
	"sort"
	"strings"
)

type HelpPlugin struct{}

func (p *HelpPlugin) Name() string {
	return "Help"
}

func (p *HelpPlugin) Help() string {
	return ""
}

func (p *HelpPlugin) Register(bot *Bot, service Service, data []byte) error {
	messageChannel := bot.NewMessageChannel(service)
	go func() {
		for {
			message := <-messageChannel

			if service.IsMe(message) {
				continue
			}

			messageChannel := message.Channel()
			messageMessage := message.Message()

			if strings.HasPrefix(messageMessage, "!help") || strings.HasPrefix(messageMessage, "!command") {
				help := make([]string, 0)

				for _, plugin := range bot.Services[service.Name()].Plugins {
					h := plugin.Help()
					if h != "" {
						help = append(help, h)
					}
				}

				sort.Strings(help)

				for _, h := range help {
					if err := service.SendMessage(messageChannel, h); err != nil {
						log.Println(err)
					}
				}
			}
		}
	}()
	return nil
}

func (p *HelpPlugin) Save() []byte {
	return nil
}

func NewHelpPlugin() *HelpPlugin {
	return &HelpPlugin{}
}
