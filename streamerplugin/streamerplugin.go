package streamerplugin

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/iopred/bruxism"
)

type streamerPluginRequest struct {
	lastUpdate  time.Time
	lastMessage string
}

type streamerPlugin struct {
	bruxism.SimplePlugin
	youTube  *bruxism.YouTube
	requests map[string]*streamerPluginRequest
}

func (p *streamerPlugin) helpFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, detailed bool) []string {
	if detailed {
		return nil
	}
	return bruxism.CommandHelp(service, "streamer", "<streamername|streamerid>", "Grabs details about a YouTube streamer.")
}

func (p *streamerPlugin) messageFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	if service.IsMe(message) {
		return
	}

	if !bruxism.MatchesCommand(service, "streamer", message) {
		return
	}

	query, parts := bruxism.ParseCommand(service, message)

	if len(parts) == 0 {
		return
	}

	r, _ := p.requests[query]
	if r == nil {
		r = &streamerPluginRequest{}
		p.requests[query] = r
	}

	n := time.Now()
	if !n.After(r.lastUpdate.Add(60 * time.Minute)) {
		if r.lastMessage != "" {
			service.SendMessage(message.Channel(), fmt.Sprintf("%s *Last updated %s.*", r.lastMessage, humanize.Time(r.lastUpdate)))
		}
		return
	}

	service.Typing(message.Channel())

	r.lastUpdate = n

	m, err := p.streamer(query, service.Name() == bruxism.DiscordServiceName)
	if err != nil {
		service.SendMessage(message.Channel(), "There was an error while requesting the streamer, please try again later.")
		return
	}

	service.SendMessage(message.Channel(), m)
	r.lastMessage = m
}

func (p *streamerPlugin) streamer(search string, bold bool) (string, error) {
	channelList, err := p.youTube.Service.Channels.List("id,snippet,statistics").ForUsername(search).Do()

	if err != nil {
		return "", nil
	}

	// If a straight username lookup failed, do a search.
	if len(channelList.Items) == 0 {
		list, err := p.youTube.Service.Search.List("id").Q(search).Type("channel").Do()

		if err != nil {
			return "", err
		}

		if len(list.Items) == 0 {
			return "No users found with that name.", nil
		}

		channelList, err = p.youTube.Service.Channels.List("id,snippet,statistics").Id(list.Items[0].Id.ChannelId).Do()

		if err != nil {
			return "", nil
		}

		if len(channelList.Items) == 0 {
			return "No users found with that name.", nil
		}
	}

	subscriberCount := ""
	if !channelList.Items[0].Statistics.HiddenSubscriberCount {
		subscriberCount = fmt.Sprintf("%s subscribers, ", humanize.Comma(int64(channelList.Items[0].Statistics.SubscriberCount)))
	}

	b := ""
	if bold {
		b = "**"
	}

	liveText := ""

	liveVideos, err := p.youTube.GetLiveVideos(channelList.Items[0].Id)
	if err == nil && len(liveVideos) > 0 {
		liveText = fmt.Sprintf(" %sCurrently live: https://gaming.youtube.com/watch?v=%s%s", b, liveVideos[0].Id, b)
	}

	return fmt.Sprintf("%s%s%s: %s%s videos, %s views.%s", b, channelList.Items[0].Snippet.Title, b, subscriberCount, humanize.Comma(int64(channelList.Items[0].Statistics.VideoCount)), humanize.Comma(int64(channelList.Items[0].Statistics.ViewCount)), liveText), nil
}

// New will create a new streamer plugin.
func New(yt *bruxism.YouTube) bruxism.Plugin {
	p := &streamerPlugin{
		SimplePlugin: *bruxism.NewSimplePlugin("Streamer"),
		youTube:      yt,
		requests:     make(map[string]*streamerPluginRequest),
	}
	p.MessageFunc = p.messageFunc
	p.HelpFunc = p.helpFunc
	return p
}
