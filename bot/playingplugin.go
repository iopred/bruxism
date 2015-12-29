package bot

import (
	"encoding/json"
	"fmt"
	"log"
)

type PlayingPlugin struct {
	Playing string
}

// The name of the plugin.
func (p *PlayingPlugin) Name() string {
	return "Playing"
}

// Loads plugin state from a byte array.
func (p *PlayingPlugin) Load(bot *Bot, service Service, data []byte) error {
	if data != nil {
		if err := json.Unmarshal(data, p); err != nil {
			log.Println("Error loading data", err)
		}
	}

	service.SetPlaying(p.Playing)

	return nil
}

// Saves plugin state to a byte array.
func (p *PlayingPlugin) Save() ([]byte, error) {
	return json.Marshal(p)
}

// Returns a list of help strings that are printed when the user requests them.
func (p *PlayingPlugin) Help(bot *Bot, service Service) []string {
	return commandHelp("playing", "<game>", fmt.Sprintf("Set which game %s is playing.", service.UserName()))
}

// Message handler.
func (p *PlayingPlugin) Message(bot *Bot, service Service, message Message) {
	if !service.IsMe(message) {
		if matchesCommand("playing", message) {
			p.Playing, _ = parseCommand(message)
			service.SetPlaying(p.Playing)
		}
	}
}

// Creates a new playing plugin.
func NewPlayingPlugin() *PlayingPlugin {
	return &PlayingPlugin{}
}
