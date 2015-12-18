package main

import (
  "fmt"
  "log"
  "strings"
  "time"

  "github.com/dustin/go-humanize"
)

type TopStreamersPlugin struct {
  youTube     *YouTube
  lastUpdate  time.Time
  lastMessage string
}

func (p *TopStreamersPlugin) Name() string {
  return "TopStreamers"
}

func (p *TopStreamersPlugin) Help() string {
  return "!topstreamers - List the current top streamers."
}

func (p *TopStreamersPlugin) Register(bot *Bot, service Service, data []byte) error {
  messageChannel := bot.NewMessageChannel(service)
  go func() {
    for {
      message := <-messageChannel

      if service.IsMe(message) {
        continue
      }

      messageChannel := message.Channel()

      if strings.HasPrefix(message.Message(), "!topstreamers") {
        n := time.Now()
        if !n.After(p.lastUpdate.Add(1 * time.Minute)) {
          if p.lastMessage != "" {
            service.SendMessage(messageChannel, fmt.Sprintf("%v *Last updated %v.*", p.lastMessage, humanize.Time(p.lastUpdate)))
          }
          continue
        }

        p.lastUpdate = n

        m, err := p.TopStreamers(5)
        p.lastMessage = m

        if err != nil {
          log.Println(err)
          continue
        }

        service.SendMessage(messageChannel, p.lastMessage)
      }
    }
  }()
  return nil
}

func (p *TopStreamersPlugin) TopStreamers(count int) (string, error) {
  videoList, err := p.youTube.GetTopLivestreams(200)
  if err != nil {
    return "", err
  }

  channels := make([]string, 0)

  for i, video := range videoList {
    channels = append(channels, fmt.Sprintf("%v (%v)", video.Snippet.ChannelTitle, humanize.Comma(int64(video.LiveStreamingDetails.ConcurrentViewersInteger))))
    if i >= count {
      break
    }
  }

  return fmt.Sprintf("Current top streamers: %v.", strings.Join(channels, ", ")), nil
}

func (p *TopStreamersPlugin) Save() []byte {
  return nil
}

func NewTopStreamersPlugin(yt *YouTube) *TopStreamersPlugin {
  return &TopStreamersPlugin{
    youTube: yt,
  }
}
