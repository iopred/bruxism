package bot

import (
	"encoding/json"
	"fmt"
	"log"
)

type PlayingPlugin struct {
	Playing string
}

func (p *PlayingPlugin) Name() string {
	return "Playing"
}

func (p *PlayingPlugin) Help(bot *Bot, service Service) []string {
	return []string{fmt.Sprintf("!playing [game] - Set which game %v is playing.", service.UserName())}
}

func (p *PlayingPlugin) Load(bot *Bot, service Service, data []byte) error {
	if data != nil {
		if err := json.Unmarshal(data, p); err != nil {
			log.Println("Error loading data", err)
		}
	}

	service.SetPlaying(p.Playing)

	return nil
}

func (p *PlayingPlugin) Message(bot *Bot, service Service, message Message) {
	if !service.IsMe(message) {
		if matchesCommand("playing", message) {
			p.Playing, _ = parseCommand(message)
			service.SetPlaying(p.Playing)
		}
	}
}

func (p *PlayingPlugin) Save() ([]byte, error) {
	if data, err := json.Marshal(p); err != nil {
		return nil, err
	} else {
		return data, nil
	}
	return nil, nil
}

func NewPlayingPlugin() *PlayingPlugin {
	return &PlayingPlugin{}
}
