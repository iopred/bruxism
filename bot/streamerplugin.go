package bot

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
)

type streamerPluginRequest struct {
	lastUpdate  time.Time
	lastMessage string
}

type streamerPlugin struct {
	SimplePlugin
	youTube  *YouTube
	requests map[string]*streamerPluginRequest
}

func (p *streamerPlugin) helpFunc(bot *Bot, service Service) []string {
	return commandHelp("streamer", "<streamername|streamerid>", "Grabs details about a streamer.")
}

func (p *streamerPlugin) messageFunc(bot *Bot, service Service, message Message) {
	if !service.IsMe(message) {
		if matchesCommand("streamer", message) {
			query, parts := parseCommand(message)

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
					service.SendMessage(message.Channel(), fmt.Sprintf("%v *Last updated %v.*", r.lastMessage, humanize.Time(r.lastUpdate)))
				}
				return
			}

			r.lastUpdate = n

			m, err := p.streamer(query)
			if err != nil {
				service.SendMessage(message.Channel(), "There was an error while requesting the streamer, please try again later.")
				return
			}

			service.SendMessage(message.Channel(), m)
			r.lastMessage = m
		}
	}
}

func (p *streamerPlugin) streamer(search string) (string, error) {
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
		subscriberCount = fmt.Sprintf("%v subscribers, ", humanize.Comma(int64(channelList.Items[0].Statistics.SubscriberCount)))
	}
	return fmt.Sprintf("%v: %v%v videos, %v views.", channelList.Items[0].Snippet.Title, subscriberCount, humanize.Comma(int64(channelList.Items[0].Statistics.VideoCount)), humanize.Comma(int64(channelList.Items[0].Statistics.ViewCount))), nil
}

// Create will create a new streamer plugin.
func NewStreamerPlugin(yt *YouTube) Plugin {
	s := &streamerPlugin{
		SimplePlugin: *NewSimplePlugin("Streamer"),
		youTube:      yt,
		requests:     make(map[string]*streamerPluginRequest),
	}
	s.message = s.messageFunc
	s.help = s.helpFunc
	return s
}
