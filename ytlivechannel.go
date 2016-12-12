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

	channelNames map[string]string
}

func NewYTLiveChannel(service *youtube.Service) *YTLiveChannel {
	return &YTLiveChannel{
		service:        service,
		channelNames:   map[string]string{},
		liveVideoChans: map[string][]chan *youtube.Video{},
	}
}

// Monitor monitors a channel for new live videos and sends them down liveVideoChan.
// If the channel is live when this is called, it will not send the video down the channel.
func (y *YTLiveChannel) Monitor(channel string, liveVideoChan chan *youtube.Video) error {
	y.Lock()
	defer y.Unlock()

	videoChans := y.liveVideoChans[channel]
	for _, v := range videoChans {
		if v == liveVideoChan {
			return errors.New("already monitoring that channel")
		}
	}

	created := len(y.liveVideoChans[channel]) == 0
	y.liveVideoChans[channel] = append(y.liveVideoChans[channel], liveVideoChan)

	if created {
		go y.poll(channel)
	}
	return nil
}

// UnmonitorAll unmonitors a channel for live videos.
func (y *YTLiveChannel) Unmonitor(channel string, liveVideoChan chan *youtube.Video) error {
	y.Lock()
	defer y.Unlock()

	videoChans := y.liveVideoChans[channel]
	for i, v := range videoChans {
		if v == liveVideoChan {
			y.liveVideoChans[channel] = append(videoChans[:i], videoChans[i+1:]...)
			return nil
		}
	}

	return errors.New("channel not being monitored")
}

func (y *YTLiveChannel) ChannelName(channel string) string {
	y.RLock()
	defer y.RUnlock()

	return y.channelNames[channel]
}

func (y *YTLiveChannel) poll(channel string) {
	var lastAnnounce time.Time
	seen := map[string]bool{}
	first := true
	for {
		videos, _ := y.getLiveVideos(channel)
		y.Lock()
		for k, v := range videos {
			if !seen[k] {
				seen[k] = true

				y.channelNames[channel] = v.Snippet.ChannelTitle

				// Don't announce the videos that are already live.
				if first {
					continue
				}
				now := time.Now()
				// Don't allow more than 1 announcement per hour.
				if !now.After(lastAnnounce.Add(1 * time.Hour)) {
					continue
				}
				lastAnnounce = now

				for _, c := range y.liveVideoChans[channel] {
					c <- v
				}
			}
		}
		if len(y.liveVideoChans[channel]) == 0 {
			y.Unlock()
			return
		}
		y.Unlock()
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
