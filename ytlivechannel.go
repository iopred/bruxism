package bruxism

import (
	"errors"
	"strings"
	"sync"
	"time"

	"google.golang.org/api/youtube/v3"
)

// YTLiveChannel is a monitor that will send new live videos to a provided channel.
type YTLiveChannel struct {
	sync.RWMutex
	service *youtube.Service
	// map channelID -> chan
	liveVideoChans map[string][]chan *youtube.Video
	channelNames   map[string]string
}

func NewYTLiveChannel(service *youtube.Service) *YTLiveChannel {
	return &YTLiveChannel{service: service}
}

func (y *YTLiveChannel) Monitor(channel string, liveVideoChan chan *youtube.Video) error {
	y.Lock()
	defer y.Unlock()

	if y.channelNames[channel] == "" {
		clr, err := y.service.Channels.List("snippet").Id(channel).Do()
		if err != nil {
			return errors.New("Error loading channel.")
		}
		if len(clr.Items) != 1 {
			return errors.New("No channel found.")
		}
		if y.channelNames == nil {
			y.channelNames = map[string]string{}
		}
		y.channelNames[channel] = clr.Items[0].Snippet.Title
	}

	if y.liveVideoChans == nil {
		y.liveVideoChans = map[string][]chan *youtube.Video{}
	}
	created := len(y.liveVideoChans[channel]) == 0
	y.liveVideoChans[channel] = append(y.liveVideoChans[channel], liveVideoChan)
	if created {
		go y.poll(channel)
	}
	return nil
}

func (y *YTLiveChannel) ChannelName(channel string) string {
	y.RLock()
	defer y.RUnlock()

	return y.channelNames[channel]
}

func (y *YTLiveChannel) poll(channel string) {
	var lastAnnounce time.Time
	seen := map[string]bool{}
	now := time.Now()
	first := true
	for {
		videos, _ := y.getLiveVideos(channel)
		for k, v := range videos {
			if !seen[k] {
				seen[k] = true
				// Don't announce the videos that are already live.
				if first {
					continue
				}
				// Don't allow more than 1 announcement per hour.
				if !now.After(lastAnnounce.Add(1 * time.Hour)) {
					continue
				}
				lastAnnounce = now
				y.RLock()
				for _, c := range y.liveVideoChans[channel] {
					c <- v
				}
				y.RUnlock()
			}
		}
		first = false
		<-time.After(5 * time.Minute)
	}

}

func (y *YTLiveChannel) getLiveVideos(channel string) (map[string]*youtube.Video, error) {
	if y.service == nil {
		return nil, errors.New("Service not available.")
	}

	search, err := y.service.Search.List("id").ChannelId(channel).EventType("live").Type("video").Do()
	if err != nil {
		return nil, err
	}

	ids := []string{}
	for _, searchResult := range search.Items {
		ids = append(ids, searchResult.Id.VideoId)
	}

	m := map[string]*youtube.Video{}

	i := 0
	for i < len(ids) {
		next := i + 50
		if next >= len(ids) {
			next = len(ids)
		}
		videoList, err := y.service.Videos.List("id,snippet,liveStreamingDetails").MaxResults(50).Id(strings.Join(ids[i:next], ",")).Do()
		if err != nil {
			return nil, err
		}

		for _, v := range videoList.Items {
			if v.LiveStreamingDetails.ActualEndTime == "" {
				m[v.Id] = v
			}
		}

		i = next
	}

	return m, nil
}
