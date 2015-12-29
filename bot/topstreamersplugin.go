package bot

import (
  "errors"
  "fmt"
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

func (p *TopStreamersPlugin) Load(bot *Bot, service Service, data []byte) error {
  return nil
}

func (p *TopStreamersPlugin) Save() ([]byte, error) {
  return nil, nil
}

func (p *TopStreamersPlugin) Help(bot *Bot, service Service) []string {
  return []string{"!topstreamers - List the current top streamers."}
}

func (p *TopStreamersPlugin) Message(bot *Bot, service Service, message Message) {
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
  return &TopStreamersPlugin{
    youTube: yt,
  }
}
