package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

type StreamerPluginRequest struct {
	lastUpdate  time.Time
	lastMessage string
}

type StreamerPlugin struct {
	youTube  *YouTube
	requests map[string]*StreamerPluginRequest
}

func (p *StreamerPlugin) Name() string {
	return "Streamer"
}

func (p *StreamerPlugin) Help() string {
	return "!streamer [streamername|streamerid] - Grabs details about a streamer."
}

func (p *StreamerPlugin) Register(bot *Bot, service Service, data []byte) error {
	messageChannel := bot.NewMessageChannel(service)
	go func() {
		for {
			message := <-messageChannel

			if service.IsMe(message) {
				continue
			}

			messageChannel := message.Channel()
			messageMessage := message.Message()

			if strings.HasPrefix(messageMessage, "!streamer") {
				splits := strings.Split(messageMessage, "!streamer ")
				if len(splits) != 2 {
					continue
				}

				query := string(splits[1])

				r, _ := p.requests[query]
				if r == nil {
					r = &StreamerPluginRequest{}
					p.requests[query] = r
				}

				n := time.Now()
				if !n.After(r.lastUpdate.Add(60 * time.Minute)) {
					if r.lastMessage != "" {
						service.SendMessage(messageChannel, fmt.Sprintf("%v *Last updated %v.*", r.lastMessage, humanize.Time(r.lastUpdate)))
					}
					continue
				}

				r.lastUpdate = n

				m, err := p.Streamer(query)
				if err != nil {
					service.SendMessage(messageChannel, "There was an error while requesting the streamer, please try again later.")
					continue
				}

				service.SendMessage(messageChannel, m)
				r.lastMessage = m
			}
		}
	}()
	return nil
}

func (p *StreamerPlugin) Streamer(search string) (string, error) {
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

func (p *StreamerPlugin) Save() []byte {
	return nil
}

func NewStreamerPlugin(yt *YouTube) *StreamerPlugin {
	return &StreamerPlugin{
		youTube:  yt,
		requests: make(map[string]*StreamerPluginRequest),
	}
}
