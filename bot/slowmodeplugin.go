package bot

import (
	"encoding/json"
	"log"
)

type SlowModePlugin struct {
	Enabled map[string]bool
}

func (p *SlowModePlugin) Name() string {
	return "SlowMode"
}

func (p *SlowModePlugin) Load(bot *Bot, service Service, data []byte) error {
	if data != nil {
		if err := json.Unmarshal(data, p); err != nil {
			log.Println("Error loading data", err)
		}
	}
	return nil
}

func (p *SlowModePlugin) Save() ([]byte, error) {
	if data, err := json.Marshal(p); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func (p *SlowModePlugin) Help(bot *Bot, service Service) []string {
	return commandHelp("slowmode", "[<on|off>]", "Turn slow mode on or off, or return the current slow mode state.")
}

func (p *SlowModePlugin) Message(bot *Bot, service Service, message Message) {
	if !service.IsMe(message) {
		messageChannel := message.Channel()

		if matchesCommand("slowmode", message) {
			enabled := p.Enabled[messageChannel]

			_, parts := parseCommand(message)

			if len(parts) == 1 {
				switch parts[0] {
				case "on":
					enabled = true
				case "off":
					enabled = false
				}
			}

			changed := enabled != p.Enabled[messageChannel]
			if changed {
				p.Enabled[messageChannel] = enabled
			}

			if enabled {
				if changed {
					service.SendMessage(messageChannel, "Slow mode is now on (You will be temporarily banned for 30 seconds when you chat).")
				} else {
					service.SendMessage(messageChannel, "Slow mode is on (You will be temporarily banned for 30 seconds when you chat).")
				}
			} else {
				if changed {
					service.SendMessage(messageChannel, "Slow mode is now off.")
				} else {
					service.SendMessage(messageChannel, "Slow mode is off.")
				}
			}
		} else if p.Enabled[messageChannel] && !message.IsModerator() {
			if err := service.BanUser(messageChannel, message.UserId(), 30); err != nil {
				log.Println(err)
			}
		}
	}
}

func NewSlowModePlugin() *SlowModePlugin {
	return &SlowModePlugin{
		Enabled: make(map[string]bool),
	}
}
