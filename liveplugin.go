package bruxism

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

const livePluginChannelId = ""

type liveChannel struct {
	live bool
}

type livePlugin struct {
	youTube *YouTube
	// Map from UserID -> ChannelID
	ChannelIDs map[string]string
	// Map from ChannelID -> Live
	Live map[string]bool
}

// Name returns the name of the plugin.
func (p *livePlugin) Name() string {
	return "Live"
}

// Load will load plugin state from a byte array.
func (p *livePlugin) Load(bot *Bot, service Service, data []byte) error {
	if data != nil {
		if err := json.Unmarshal(data, p); err != nil {
			log.Println("Error loading data", err)
		}
	}

	go p.Run(bot, service)
	return nil
}

func (p *livePlugin) pollChannel(bot *Bot, service Service, id string) {
}

func (p *livePlugin) poll(bot *Bot, service Service) {
	for _, channelID := range p.ChannelIDs {
		go p.pollChannel(bot, service, channelId)

		p.youTube.Service.Videos.List("snippet")
	}
}

// Run will poll YouTube for channels going live and send messages.
func (p *livePlugin) Run(bot *Bot, service Service) {
	for {
		p.poll(bot, service)
		<-time.After(5 * time.Minute)
	}

}

// Save will save plugin state to a byte array.
func (p *livePlugin) Save() ([]byte, error) {
	return json.Marshal(p)
}

// Help returns a list of help strings that are printed when the user requests them.
func (p *livePlugin) Help(bot *Bot, service Service, detailed bool) []string {
	return nil
}

// Message handler.
func (p *livePlugin) Message(bot *Bot, service Service, message Message) {
	defer messageRecover()
	if !service.IsMe(message) && service.IsPrivate(message) {
		if matchesCommand(service, "setyoutubechannel", message) {
			query, _ := parseCommand(service, message)
			if len(query) == 24 && strings.Index(query, ",") == -1 {
				p.ChannelIDs[message.UserID()] = query
				service.SendMessage(message.Channel(), "Channel ID set.")
			} else {
				service.SendMessage(message.Channel(), "Sorry, please provide a YouTube Channel ID. eg: UC392dac34_32fafe2deadbeef")
			}
		}
	}
}

// NewlivePlugin will create a new slow mode plugin.
func NewlivePlugin(yt *YouTube) Plugin {
	return &livePlugin{
		youTube:    yt,
		ChannelIDs: map[string]string{},
		Live:       map[string]bool{},
	}
}
