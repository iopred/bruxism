package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type PlayingPlugin struct {
	service Service
	Playing string
}

func (p *PlayingPlugin) Name() string {
	return "Playing"
}

func (p *PlayingPlugin) Help() string {
	return fmt.Sprintf("!playing [game] - Set which game %v is playing.", p.service.UserName())
}

func (p *PlayingPlugin) Register(bot *Bot, service Service, data []byte) error {
	if data != nil {
		if err := json.Unmarshal(data, p); err != nil {
			log.Println("Error loading data", err)
		}
	}

	p.service = service
	service.SetPlaying(p.Playing)

	messageChannel := bot.NewMessageChannel(service)
	go func() {
		for {
			message := <-messageChannel

			if service.IsMe(message) {
				continue
			}

			messageMessage := message.Message()

			if strings.HasPrefix(messageMessage, "!playing") {
				splits := strings.Split(messageMessage, "!playing ")
				if len(splits) == 1 {
					p.Playing = ""
				} else {
					p.Playing = splits[1]
				}
				service.SetPlaying(p.Playing)
			}
		}
	}()
	return nil
}

func (p *PlayingPlugin) Save() []byte {
	if data, err := json.Marshal(p); err == nil {
		return data
	}
	return nil
}

func NewPlayingPlugin() *PlayingPlugin {
	return &PlayingPlugin{}
}
