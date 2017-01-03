package playingplugin

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/iopred/bruxism"
)

type playingPlugin struct {
	bruxism.SimplePlugin
	Game string
	URL  string
}

// Name returns the name of the plugin.
func (p *playingPlugin) Name() string {
	return "Playing"
}

// Load will load plugin state from a byte array.
func (p *playingPlugin) Load(bot *bruxism.Bot, service bruxism.Service, data []byte) error {
	if service.Name() != bruxism.DiscordServiceName {
		panic("Playing Plugin only supports Discord.")
	}

	if data != nil {
		if err := json.Unmarshal(data, p); err != nil {
			log.Println("Error loading data", err)
		}
	}

	if p.Game != "" {
		service.(*bruxism.Discord).Session.UpdateStreamingStatus(0, p.Game, p.URL)
	} else {
		service.(*bruxism.Discord).Session.UpdateStatus(0, p.Game)
	}

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

	if !service.IsBotOwner(message) {
		return nil
	}

	return bruxism.CommandHelp(service, "playing", "<game>, <url>", fmt.Sprintf("Set which game %s is playing.", service.UserName()))
}

// Message handler.
func (p *playingPlugin) messageFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	if service.IsMe(message) {
		return
	}

	if !bruxism.MatchesCommand(service, "playing", message) {
		return
	}

	if !service.IsBotOwner(message) {
		return
	}

	query, _ := bruxism.ParseCommand(service, message)

	split := strings.Split(query, ",")

	p.Game = strings.Trim(split[0], " ")
	if len(split) > 1 {
		p.URL = strings.Trim(split[1], " ")
		service.(*bruxism.Discord).Session.UpdateStreamingStatus(0, p.Game, p.URL)
	} else {
		service.(*bruxism.Discord).Session.UpdateStatus(0, p.Game)
	}
}

// New will create a new top streamers plugin.
func New() bruxism.Plugin {
	p := &playingPlugin{}
	p.MessageFunc = p.messageFunc
	p.HelpFunc = p.helpFunc
	return p
}
