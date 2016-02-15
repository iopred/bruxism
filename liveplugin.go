package bruxism

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

const livePluginChannelId = ""

type livePlugin struct {
	youTube *YouTube
	Initial bool
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

type LiveChatStatusListResponse struct {
	Kind     string `json:"kind"`
	Etag     string `json:"etag"`
	PageInfo struct {
		TotalResults   int `json:"totalResults,omitempty"`
		ResultsPerPage int `json:"resultsPerPage,omitempty"`
	} `json:"pageInfo,omitempty"`
	Items []struct {
		Kind    string `json:"kind"`
		Etag    string `json:"etag"`
		Id      string `json:"id"`
		Snippet struct {
			ChannelID string `json:"channelId"`
		} `json:"snippet"`
	} `json:"items"`
}

func (p *livePlugin) poll(bot *Bot, service Service) {
	channelIDs := make([]string, 0, len(p.ChannelIDs))
	for channelID := range p.ChannelIDs {
		channelIDs := append(channelIDs, channelID)
	}
	resp, err := p.youTube.Client.Get("https://www.googleapis.com/youtube/v3/liveChat/status?part=id,snippet&channelId=" + strings.Join(channelIDs, ","))
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	liveChatStatusListResponse := &LiveChatStatusListResponse{}
	err = json.Unmarshal(body, liveChatStatusListResponse)
	if err != nil {
		return
	}

	NewLive := map[string]bool{}

	for _, i := range liveChatStatusListResponse.Items {
		NewLive[i.Snippet.ChannelID] = true
	}

	if !p.Initial {
		for i := range NewLive {
			if !p.Live[i] {
				service.SendMessage(livePluginChannelId, fmt.Sprintf("<@%s> has gone live."))
			}
		}
	}

	p.Live = NewLive

	p.Initial = true
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
