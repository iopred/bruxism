package bot

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

type TopStreamersPlugin struct {
	SimplePlugin
	youTube     *YouTube
	lastUpdate  time.Time
	lastMessage string
}

func (p *TopStreamersPlugin) helpFunc(bot *Bot, service Service) []string {
	return commandHelp("topstreamers", "", "List the current top streamers.")
}

func (p *TopStreamersPlugin) messageFunc(bot *Bot, service Service, message Message) {
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

			m, err := p.TopStreamers(5)
			if err != nil {
				service.SendMessage(message.Channel(), "There was an error while requesting the top streamers, please try again later.")
				return
			}

			service.SendMessage(message.Channel(), m)
			p.lastMessage = m
		}
	}
}

func (p *TopStreamersPlugin) TopStreamers(count int) (string, error) {
	videoList, err := p.youTube.GetTopLivestreams(200)
	if err != nil {
		return "", err
	}

	if len(videoList) == 0 {
		return "", errors.New("No videos returned.")
	}

	channels := make([]string, 0)

	for i, video := range videoList {
		channels = append(channels, fmt.Sprintf("%v (%v)", video.Snippet.ChannelTitle, humanize.Comma(int64(video.LiveStreamingDetails.ConcurrentViewers))))
		if i >= count {
			break
		}
	}

	return fmt.Sprintf("Current top streamers: %v.", strings.Join(channels, ", ")), nil
}

func NewTopStreamersPlugin(yt *YouTube) *TopStreamersPlugin {
	ts := &TopStreamersPlugin{
		SimplePlugin: *NewSimplePlugin("TopStreamers"),
		youTube:      yt,
	}
	ts.message = ts.messageFunc
	ts.help = ts.helpFunc
	return ts
}
