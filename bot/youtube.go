package bot

import (
	"encoding/json"
	"errors"
	"html"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/atotto/clipboard"
	ytc "github.com/iopred/ytlivechatapi"
	"google.golang.org/api/youtube/v3"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// YouTubeServiceName is the service name for the YouTube service.
const YouTubeServiceName string = "YouTube"

// LiveChatMessage is a Message wrapper around ytc.LiveChatMessage.
type LiveChatMessage ytc.LiveChatMessage

// Channel returns the channel id for this message.
func (m *LiveChatMessage) Channel() string {
	return m.Snippet.LiveChatId
}

// UserName returns the user name for this message.
func (m *LiveChatMessage) UserName() string {
	return m.AuthorDetails.DisplayName
}

// UserID returns the user id for this message.
func (m *LiveChatMessage) UserID() string {
	return m.AuthorDetails.ChannelId
}

// Message returns the message content for this message.
func (m *LiveChatMessage) Message() string {
	switch m.Snippet.Type {
	case ytc.LiveChatMessageSnippetTypeText:
		return html.UnescapeString(m.Snippet.TextMessageDetails.MessageText)
	}
	return html.UnescapeString(m.Snippet.DisplayMessage)
}

// MessageID returns the message ID for this message.
func (m *LiveChatMessage) MessageID() string {
	return m.Id
}

// IsModerator returns whether or not the sender of this message is a moderator.
func (m *LiveChatMessage) IsModerator() bool {
	return m.AuthorDetails.IsChatOwner || m.AuthorDetails.IsChatModerator
}

type fanFunding struct {
	sync.Mutex
	Messages map[string]*ytc.LiveChatMessage
}

// YouTube is a Service provider for YouTube.
type YouTube struct {
	url            bool
	auth           string
	configFilename string
	tokenFilename  string
	liveChatIds    string
	config         *oauth2.Config
	token          *oauth2.Token
	Client         *ytc.Client
	Service        *youtube.Service
	messageChan    chan Message
	InsertChan     chan interface{}
	DeleteChan     chan interface{}
	fanFunding     fanFunding
	me             *youtube.Channel
}

// NewYouTube creates a new YouTube service.
func NewYouTube(url bool, auth, configFilename, tokenFilename, liveChatIds string) *YouTube {
	return &YouTube{
		url:            url,
		auth:           auth,
		configFilename: configFilename,
		tokenFilename:  tokenFilename,
		liveChatIds:    liveChatIds,
		messageChan:    make(chan Message, 200),
		InsertChan:     make(chan interface{}, 200),
		DeleteChan:     make(chan interface{}, 200),
		fanFunding:     fanFunding{Messages: make(map[string]*ytc.LiveChatMessage)},
	}
}

func (yt *YouTube) pollBroadcasts(broadcasts *ytc.LiveBroadcastListResponse, err error) {
	if err != nil {
		log.Println(err)
		return
	}

	for _, broadcast := range broadcasts.Items {
		// If the broadcast has ended, it can't have a valid chat.
		if broadcast.Status != nil && broadcast.Status.LifeCycleStatus == "complete" {
			continue
		}

		go yt.pollMessages(broadcast)
	}
}

func (yt *YouTube) pollMessages(broadcast *ytc.LiveBroadcast) {
	pageToken := ""
	for {
		liveChatMessageListResponse, err := yt.Client.ListLiveChatMessages(broadcast.Snippet.LiveChatId, pageToken)

		if err != nil {
			log.Println(err)
		} else if liveChatMessageListResponse.Error != nil {
			log.Println(liveChatMessageListResponse.Error.NewError("polling messages"))
		} else {
			// Ignore the first results, we only want new chats.
			if pageToken != "" {
				for _, message := range liveChatMessageListResponse.Items {
					liveChatMessage := LiveChatMessage(*message)
					yt.messageChan <- &liveChatMessage

					switch message.Snippet.Type {
					case ytc.LiveChatMessageSnippetTypeFanFunding:
						yt.addFanFundingMessage(message)
					}
				}
			}
			pageToken = liveChatMessageListResponse.NextPageToken
		}

		if liveChatMessageListResponse.PollingIntervalMillis != 0 {
			time.Sleep(time.Duration(liveChatMessageListResponse.PollingIntervalMillis) * time.Millisecond)
		} else {
			time.Sleep(10 * time.Second)
		}
	}
}

func (yt *YouTube) writeMessagesToFile(messages []*ytc.LiveChatMessage, filename string) {
	output := ""
	for _, message := range messages {
		output += html.UnescapeString(message.Snippet.DisplayMessage) + "\n"
	}
	err := ioutil.WriteFile(filename, []byte(output), 0777)
	if err != nil {
		log.Println(err)
	}
}

func (yt *YouTube) addFanFundingMessage(message *ytc.LiveChatMessage) {
	yt.fanFunding.Lock()
	defer yt.fanFunding.Unlock()

	if yt.fanFunding.Messages[message.Id] == nil {
		yt.fanFunding.Messages[message.Id] = message
		yt.writeMessagesToFile([]*ytc.LiveChatMessage{message}, "youtubelatest.txt")
	}

	largest := message
	for _, check := range yt.fanFunding.Messages {
		if check.Snippet.FanFundingEventDetails.AmountMicros > largest.Snippet.FanFundingEventDetails.AmountMicros {
			largest = check
		}
	}

	yt.writeMessagesToFile([]*ytc.LiveChatMessage{largest}, "youtubelargest.txt")
}

func (yt *YouTube) generateOauthURLAndExit() {
	// Redirect user to Google's consent page to ask for permission
	// for the scopes specified above.
	url := yt.config.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	clipboard.WriteAll(url)
	log.Fatalln("Visit the following URL to generate an auth code, then rerun with -auth=<code> (It has also been copied to your clipboard):\n%v", url)
}

func (yt *YouTube) createConfig() error {
	b, err := ioutil.ReadFile(yt.configFilename)
	if err != nil {
		return err
	}

	yt.config, err = google.ConfigFromJSON(b, "https://www.googleapis.com/auth/youtube")
	return err
}

func (yt *YouTube) createToken() error {
	if yt.auth != "" {
		// Let's exchange our code
		tok, err := yt.config.Exchange(oauth2.NoContext, yt.auth)
		if err != nil {
			return err
		}
		yt.token = tok

		b, err := json.Marshal(yt.token)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(yt.tokenFilename, b, 0777)
		if err != nil {
			return err
		}
	} else {
		b, err := ioutil.ReadFile(yt.tokenFilename)
		if err == nil {
			yt.token = &oauth2.Token{}
			// A token was found, unmarshall it and use it.
			err := json.Unmarshal(b, yt.token)
			if err != nil {
				return err
			}
		} else {
			// There was an error with the token, maybe it doesn't exist.
			// If we haven't been given an auth code, we must generate one.
			yt.generateOauthURLAndExit()
		}
	}
	return nil
}

// Name returns the name of the service.
func (yt *YouTube) Name() string {
	return YouTubeServiceName
}

func (yt *YouTube) handleRequests() {
	// This is a map of channel id's to channels, it is used to send messages to a goroutine that is rate limiting each chatroom.
	channelInsertChans := make(map[string]chan *ytc.LiveChatMessage)

	// Chat messages need to be separated by one second, they must be handled by a separate goroutine.
	insertLiveChatMessageLimited := func(liveChatMessage *ytc.LiveChatMessage) {
		channelInsertChan := channelInsertChans[liveChatMessage.Snippet.LiveChatId]
		if channelInsertChan == nil {
			channelInsertChan = make(chan *ytc.LiveChatMessage, 200)
			channelInsertChans[liveChatMessage.Snippet.LiveChatId] = channelInsertChan

			// Start a goroutine to rate limit sends.
			go func() {
				for {
					if err := yt.Client.InsertLiveChatMessage(<-channelInsertChan); err != nil {
						log.Println(err)
					}
					time.Sleep(1 * time.Second)
				}
			}()
		}
		channelInsertChan <- liveChatMessage
	}

	for {
		var err error
		select {
		case request := <-yt.InsertChan:
			switch request := request.(type) {
			case *ytc.LiveChatMessage:
				insertLiveChatMessageLimited(request)
			case *ytc.LiveChatBan:
				err = yt.Client.InsertLiveChatBan(request)
			case *ytc.LiveChatModerator:
				err = yt.Client.InsertLiveChatModerator(request)
			}
		case request := <-yt.DeleteChan:
			switch request := request.(type) {
			case *ytc.LiveChatMessage:
				err = yt.Client.DeleteLiveChatMessage(request)
			case *ytc.LiveChatBan:
				err = yt.Client.DeleteLiveChatBan(request)
			case *ytc.LiveChatModerator:
				err = yt.Client.DeleteLiveChatModerator(request)
			}
		}
		if err != nil {
			log.Println(err)
		}

		// Sleep for a millisecond, this will guarantee a maximum QPS of 1000.
		time.Sleep(1 * time.Millisecond)
	}
}

// Open opens the service and returns a channel which all messages will be sent on.
func (yt *YouTube) Open() (<-chan Message, error) {
	if err := yt.createConfig(); err != nil {
		return nil, err
	}

	// An oauth URL was requested, error early.
	if yt.url {
		yt.generateOauthURLAndExit()
	}

	if err := yt.createToken(); err != nil {
		return nil, err
	}

	client := yt.config.Client(oauth2.NoContext, yt.token)
	yt.Client = ytc.NewClient(client)

	if service, err := youtube.New(client); err == nil {
		yt.Service = service
	} else {
		return nil, err
	}

	me, err := yt.GetMe()
	if err != nil {
		return nil, err
	}
	yt.me = me

	yt.pollBroadcasts(yt.Client.ListLiveBroadcasts("default=true"))
	yt.pollBroadcasts(yt.Client.ListLiveBroadcasts("mine=true"))
	if yt.liveChatIds != "" {
		liveChatIDsArray := strings.Split(yt.liveChatIds, ",")

		for _, liveChatID := range liveChatIDsArray {
			yt.Join(liveChatID)
		}
	}

	// Start a goroutine to handle all requests.
	go yt.handleRequests()

	return yt.messageChan, nil
}

// IsMe returns whether or not a message was sent by the bot.
func (yt *YouTube) IsMe(message Message) bool {
	return message.UserID() == yt.me.Id
}

// SendMessage sends a message.
func (yt *YouTube) SendMessage(channel, message string) error {
	yt.InsertChan <- ytc.NewLiveChatMessage(channel, message)
	return nil
}

// DeleteMessage deletes a message.
func (yt *YouTube) DeleteMessage(messageID string) error {
	yt.DeleteChan <- &ytc.LiveChatMessage{Id: messageID}
	return nil
}

// BanUser bans a user.
func (yt *YouTube) BanUser(channel, user string, duration int) error {
	yt.InsertChan <- ytc.NewLiveChatBan(channel, user, duration)
	return nil
}

// UserName returns the bots name.
func (yt *YouTube) UserName() string {
	return yt.me.Snippet.Title
}

// SetPlaying will set the current game being played by the bot.
func (yt *YouTube) SetPlaying(game string) error {
	return errors.New("Set playing not supported on YouTube.")
}

// Join will join a channel.
func (yt *YouTube) Join(join string) error {
	liveChatMessageListResponse, err := yt.Client.ListLiveChatMessages(join, "")
	if err != nil {
		return err
	}

	if liveChatMessageListResponse.Error != nil {
		return liveChatMessageListResponse.Error.NewError("joining channel")
	}

	liveBroadcastListResponse := &ytc.LiveBroadcastListResponse{
		Items: []*ytc.LiveBroadcast{
			{
				Snippet: &ytc.LiveBroadcastSnippet{
					LiveChatId: join,
				},
			},
		},
	}
	yt.pollBroadcasts(liveBroadcastListResponse, nil)

	return nil
}

type videoList []*youtube.Video

func (v videoList) Len() int {
	return len(v)
}

func (v videoList) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v videoList) Less(i, j int) bool {
	return v[i].LiveStreamingDetails.ConcurrentViewers > v[j].LiveStreamingDetails.ConcurrentViewers
}

// GetMe gets the channel for the bot.
func (yt *YouTube) GetMe() (*youtube.Channel, error) {
	channelList, err := yt.Service.Channels.List("id,snippet").Mine(true).Do()

	if err != nil {
		return nil, err
	}

	if len(channelList.Items) != 1 {
		return nil, errors.New("Invalid response while requesting Me")
	}

	return channelList.Items[0], nil
}

// GetTopLivestreamIDs gets the video ids for the current top livestreams.
func (yt *YouTube) GetTopLivestreamIDs(count int) ([]string, error) {
	ids := []string{}

	pageToken := ""

	for {
		list := yt.Service.PlaylistItems.List("id,contentDetails").MaxResults(50).PlaylistId("PLiCvVJzBupKmEehQ3hnNbbfBjLUyvGlqx")
		if pageToken != "" {
			list.PageToken(pageToken)
		}
		playlistItemListResponse, err := list.Do()

		if err != nil {
			return nil, err
		}

		for _, playlistItem := range playlistItemListResponse.Items {
			ids = append(ids, playlistItem.ContentDetails.VideoId)
		}

		if len(ids) >= count || playlistItemListResponse.NextPageToken == "" {
			break
		}

		pageToken = playlistItemListResponse.NextPageToken
	}

	return ids, nil
}

// GetVideosByIDList gets all the videos for a list of video IDs.
func (yt *YouTube) GetVideosByIDList(ids []string) ([]*youtube.Video, error) {
	videos := []*youtube.Video{}

	i := 0
	for i < len(ids) {
		next := i + 50
		if next >= len(ids) {
			next = len(ids)
		}
		list := yt.Service.Videos.List("id,snippet,liveStreamingDetails").MaxResults(50).Id(strings.Join(ids[i:next], ","))

		videoListResponse, err := list.Do()

		if err != nil {
			return nil, err
		}

		videos = append(videos, videoListResponse.Items...)

		if len(videos) != len(ids) {
			break
		}
	}

	return videos, nil
}

// GetTopLivestreams gets a list of videos for the current top livestreams.
func (yt *YouTube) GetTopLivestreams(count int) ([]*youtube.Video, error) {
	ids, err := yt.GetTopLivestreamIDs(count)
	if err != nil {
		return nil, nil
	}

	videos, err := yt.GetVideosByIDList(ids)
	if err != nil {
		return nil, nil
	}

	videoList := videoList(videos)
	sort.Sort(videoList)

	return videoList, nil
}
