package playingplugin

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/iopred/bruxism"
)

type playingPlugin struct {
	bruxism.SimplePlugin
	Playing string
}

// Name returns the name of the plugin.
func (p *playingPlugin) Name() string {
	return "Playing"
}

// Load will load plugin state from a byte array.
func (p *playingPlugin) Load(bot *bruxism.Bot, service bruxism.Service, data []byte) error {
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
func (p *playingPlugin) helpFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, detailed bool) []string {
	if detailed {
		return nil
	}

	discord := service.(*bruxism.Discord)
	if discord.OwnerUserID != "" && !service.IsBotOwner(message) {
		return nil
	}

	return bruxism.CommandHelp(service, "playing", "<game>", fmt.Sprintf("Set which game %s is playing.", service.UserName()))
}

// Message handler.
func (p *playingPlugin) messageFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	if !service.IsMe(message) {
		if bruxism.MatchesCommand(service, "playing", message) {
			discord := service.(*bruxism.Discord)
			if discord.OwnerUserID != "" && !service.IsBotOwner(message) {
				return
			}

			p.Playing, _ = bruxism.ParseCommand(service, message)
			service.SetPlaying(p.Playing)
		}
	}
}

// New will create a new top streamers plugin.
func New() bruxism.Plugin {
	p := &playingPlugin{}
	p.MessageFunc = p.messageFunc
	p.HelpFunc = p.helpFunc
	return p
}
