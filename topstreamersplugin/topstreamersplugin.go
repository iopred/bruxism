package topstreamersplugin

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/iopred/bruxism"
)

type topStreamersPlugin struct {
	bruxism.SimplePlugin
	youTube     *bruxism.YouTube
	lastUpdate  time.Time
	lastMessage string
}

func (p *topStreamersPlugin) helpFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, detailed bool) []string {
	if detailed {
		return nil
	}
	return bruxism.CommandHelp(service, "topstreamers", "", "List the current top streamers on YouTube Gaming.")
}

func (p *topStreamersPlugin) messageFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	if service.IsMe(message) {
		return
	}

	if !bruxism.MatchesCommand(service, "topstreamers", message) {
		return
	}

	n := time.Now()
	if !n.After(p.lastUpdate.Add(1 * time.Minute)) {
		if p.lastMessage != "" {
			service.SendMessage(message.Channel(), fmt.Sprintf("%s *Last updated %s.*", p.lastMessage, humanize.Time(p.lastUpdate)))
		}
		return
	}

	service.Typing(message.Channel())

	p.lastUpdate = n

	m, err := p.topStreamers(5)
	if err != nil {
		service.SendMessage(message.Channel(), "There was an error while requesting the top streamers, please try again later.")
		return
	}

	service.SendMessage(message.Channel(), m)
	p.lastMessage = m
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
		channels = append(channels, fmt.Sprintf("%s (%s)", video.Snippet.ChannelTitle, humanize.Comma(int64(video.LiveStreamingDetails.ConcurrentViewers))))
		if i >= count {
			break
		}
	}

	return fmt.Sprintf("Current YouTube Gaming top streamers: %s.", strings.Join(channels, ", ")), nil
}

// New will create a new top streamers plugin.
func New(yt *bruxism.YouTube) bruxism.Plugin {
	p := &topStreamersPlugin{
		SimplePlugin: *bruxism.NewSimplePlugin("TopStreamers"),
		youTube:      yt,
	}
	p.MessageFunc = p.messageFunc
	p.HelpFunc = p.helpFunc
	return p
}
