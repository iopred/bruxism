package bot

import (
	"encoding/json"
	"fmt"
	"log"
)

type playingPlugin struct {
	Playing string
}

// Name returns the name of the plugin.
func (p *playingPlugin) Name() string {
	return "Playing"
}

// Load will load plugin state from a byte array.
func (p *playingPlugin) Load(bot *Bot, service Service, data []byte) error {
	if data != nil {
		if err := json.Unmarshal(data, p); err != nil {
			log.Println("Error loading data", err)
		}
	}

	service.SetPlaying(p.Playing)

	return nil
}

// Save will save plugin state to a byte array.
func (p *playingPlugin) Save() ([]byte, error) {
	return json.Marshal(p)
}

// Help returns a list of help strings that are printed when the user requests them.
func (p *playingPlugin) Help(bot *Bot, service Service) []string {
	return commandHelp("playing", "<game>", fmt.Sprintf("Set which game %s is playing.", service.UserName()))
}

// Message handler.
func (p *playingPlugin) Message(bot *Bot, service Service, message Message) {
	if !service.IsMe(message) {
		if matchesCommand("playing", message) {
			p.Playing, _ = parseCommand(message)
			service.SetPlaying(p.Playing)
		}
	}
}

// NewPlayingPlugin will create a new playing plugin.
func NewPlayingPlugin() Plugin {
	return &playingPlugin{}
}
