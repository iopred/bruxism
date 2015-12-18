package main

import (
  "fmt"
  "log"
  "strings"
  "time"
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
  return "!topstreamers - Returns a list of the current top streamers."
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
            service.SendMessage(messageChannel, fmt.Sprintf("%v. *Last updated %v seconds ago.*", p.lastMessage, int(n.Sub(p.lastUpdate)/time.Second)))
          }
          continue
        }

        p.lastUpdate = n

        m, err := p.TopStreamers()
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

func (p *TopStreamersPlugin) TopStreamers() (string, error) {
  videos, err := p.youTube.GetTopLivestreams(5)
  if err != nil {
    return "", err
  }

  channels := make([]string, 0)

  for _, video := range videos {
    channels = append(channels, fmt.Sprintf("%v (%v)", video.Snippet.ChannelTitle, video.LiveStreamingDetails.ConcurrentViewers))
  }

  return fmt.Sprintf("Current top streamers: %v", strings.Join(channels, ", ")), nil
}

func (p *TopStreamersPlugin) Save() []byte {
  return nil
}

func NewTopStreamersPlugin(yt *YouTube) *TopStreamersPlugin {
  return &TopStreamersPlugin{
    youTube: yt,
  }
}
