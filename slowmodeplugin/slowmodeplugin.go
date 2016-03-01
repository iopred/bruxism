package slowmodeplugin

import (
	"encoding/json"
	"log"

	"github.com/iopred/bruxism"
)

type slowModePlugin struct {
	Enabled map[string]bool
}

// Name returns the name of the plugin.
func (p *slowModePlugin) Name() string {
	return "SlowMode"
}

// Load will load plugin state from a byte array.
func (p *slowModePlugin) Load(bot *bruxism.Bot, service bruxism.Service, data []byte) error {
	if data != nil {
		if err := json.Unmarshal(data, p); err != nil {
			log.Println("Error loading data", err)
		}
	}
	return nil
}

// Save will save plugin state to a byte array.
func (p *slowModePlugin) Save() ([]byte, error) {
	return json.Marshal(p)
}

// Help returns a list of help strings that are printed when the user requests them.
func (p *slowModePlugin) Help(bot *bruxism.Bot, service bruxism.Service, detailed bool) []string {
	if detailed {
		return nil
	}
	return bruxism.CommandHelp(service, "slowmode", "[<on|off>]", "Turn slow mode on or off, or return the current slow mode state.")
}

// Message handler.
func (p *slowModePlugin) Message(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	defer bruxism.MessageRecover()
	if !service.IsMe(message) {
		messageChannel := message.Channel()

		if bruxism.MatchesCommand(service, "slowmode", message) {
			enabled := p.Enabled[messageChannel]

			_, parts := bruxism.ParseCommand(service, message)

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
		} else if p.Enabled[messageChannel] && !service.IsModerator(message) {
			if err := service.BanUser(messageChannel, message.UserID(), 30); err != nil {
				log.Println(err)
			}
		}
	}
}

// NewSlowModePlugin will create a new slow mode plugin.
func NewSlowModePlugin() bruxism.Plugin {
	return &slowModePlugin{
		Enabled: make(map[string]bool),
	}
}
