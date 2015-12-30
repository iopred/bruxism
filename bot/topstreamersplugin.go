package bot

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

type topStreamersPlugin struct {
	SimplePlugin
	youTube     *YouTube
	lastUpdate  time.Time
	lastMessage string
}

func (p *topStreamersPlugin) helpFunc(bot *Bot, service Service) []string {
	return commandHelp("topstreamers", "", "List the current top streamers on YouTube.")
}

func (p *topStreamersPlugin) messageFunc(bot *Bot, service Service, message Message) {
	if !service.IsMe(message) {
		if matchesCommand("topstreamers", message) {
			n := time.Now()
			if !n.After(p.lastUpdate.Add(1 * time.Minute)) {
				if p.lastMessage != "" {
					service.SendMessage(message.Channel(), fmt.Sprintf("%v *Last updated %v.*", p.lastMessage, humanize.Time(p.lastUpdate)))
				}
				return
			}

			p.lastUpdate = n

			m, err := p.topStreamers(5)
			if err != nil {
				service.SendMessage(message.Channel(), "There was an error while requesting the top streamers, please try again later.")
				return
			}

			service.SendMessage(message.Channel(), m)
			p.lastMessage = m
		}
	}
}

func (p *topStreamersPlugin) topStreamers(count int) (string, error) {
	videoList, err := p.youTube.GetTopLivestreams(200)
	if err != nil {
		return "", err
	}

	if len(videoList) == 0 {
		return "", errors.New("No videos returned.")
	}

	channels := []string{}

	for i, video := range videoList {
		channels = append(channels, fmt.Sprintf("%v (%v)", video.Snippet.ChannelTitle, humanize.Comma(int64(video.LiveStreamingDetails.ConcurrentViewers))))
		if i >= count {
			break
		}
	}

	return fmt.Sprintf("Current top streamers: %v.", strings.Join(channels, ", ")), nil
}

// NewTopStreamersPlugin will create a new top streamers plugin.
func NewTopStreamersPlugin(yt *YouTube) Plugin {
	p := &topStreamersPlugin{
		SimplePlugin: *NewSimplePlugin("TopStreamers"),
		youTube:      yt,
	}
	p.message = p.messageFunc
	p.help = p.helpFunc
	return p
}
