package main

import (
	"encoding/json"
	"log"
	"strings"
)

type SlowModePlugin struct {
	Enabled map[string]bool
}

func (e *SlowModePlugin) Name() string {
	return "SlowMode"
}

func (e *SlowModePlugin) Register(bot *Bot, service Service, data []byte) error {
	if data != nil {
		if err := json.Unmarshal(data, e); err != nil {
			log.Println("Error loading data", err)
		}
	}

	messageChannel := bot.NewMessageChannel(service)
	go func() {
		for {
			message := <-messageChannel
			messageMessage := strings.Replace(message.Message(), " ", "", -1)
			messageChannel := message.Channel()

			if messageMessage == "!slowmodeon" && message.IsModerator() {
				e.Enabled[messageChannel] = true
				service.SendMessage(messageChannel, "Slow mode is now on (You will be temporarily banned for 30 seconds when you chat).")
			} else if messageMessage == "!slowmodeoff" && message.IsModerator() {
				e.Enabled[messageChannel] = false
				service.SendMessage(messageChannel, "Slow mode is now off.")
			} else if messageMessage == "!slowmode" && message.IsModerator() {
				if e.Enabled[messageChannel] {
					service.SendMessage(messageChannel, "Slow mode is on.")
				} else {
					service.SendMessage(messageChannel, "Slow mode is off.")
				}
			} else if e.Enabled[messageChannel] && !message.IsModerator() {
				if err := service.BanUser(messageChannel, message.UserId(), 30); err != nil {
					log.Println(err)
				}
			}
		}
	}()
	return nil
}

func (e *SlowModePlugin) Save() []byte {
	if data, err := json.Marshal(e); err == nil {
		return data
	}
	return nil
}

func NewSlowModePlugin() *SlowModePlugin {
	return &SlowModePlugin{
		Enabled: make(map[string]bool),
	}
}
