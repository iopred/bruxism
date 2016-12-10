package bruxism

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/atotto/clipboard"
	"google.golang.org/api/youtube/v3"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// YouTubeServiceName is the service name for the YouTube service.
const YouTubeServiceName string = "YouTube"

// Enums for LiveMessageSnippet types.
const (
	LiveChatMessageSnippetTypeText       string = "textMessageEvent"
	LiveChatMessageSnippetTypeFanFunding string = "fanFundingEvent"
)

// Enums for LiveChatBanSnippet types.
const (
	LiveChatBanSnippetTypeTemporary string = "temporary"
	LiveChatBanSnippetTypePermanent        = "permanent"
)

// LiveChatMessage is a Message wrapper around youtube.LiveChatMessage.
type LiveChatMessage youtube.LiveChatMessage

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

// UserAvatar returns the avatar url for this message.
func (m *LiveChatMessage) UserAvatar() string {
	return m.AuthorDetails.ProfileImageUrl
}

// Message returns the message content for this message.
func (m *LiveChatMessage) Message() string {
	switch m.Snippet.Type {
	case LiveChatMessageSnippetTypeText:
		return html.UnescapeString(m.Snippet.TextMessageDetails.MessageText)
	}
	return html.UnescapeString(m.Snippet.DisplayMessage)
}

// RawMessage returns the message content for this message.
func (m *LiveChatMessage) RawMessage() string {
	switch m.Snippet.Type {
	case LiveChatMessageSnippetTypeText:
		return m.Snippet.TextMessageDetails.MessageText
	}
	return m.Snippet.DisplayMessage
}

// MessageID returns the message ID for this message.
func (m *LiveChatMessage) MessageID() string {
	return m.Id
}

// Type returns the type of message.
func (m *LiveChatMessage) Type() MessageType {
	return MessageTypeCreate
}

type fanFunding struct {
	sync.Mutex
	Messages map[string]*youtube.LiveChatMessage
}

// YouTube is a Service provider for YouTube.
type YouTube struct {
	url            bool
	auth           string
	configFilename string
	tokenFilename  string
	config         *oauth2.Config
	token          *oauth2.Token
	Client         *http.Client
	Service        *youtube.Service
	messageChan    chan Message
	InsertChan     chan interface{}
	DeleteChan     chan interface{}
	fanFunding     fanFunding
	me             *youtube.Channel
	channelCount   int
	joined         map[string]bool
}

// NewYouTube creates a new YouTube service.
func NewYouTube(url bool, auth, configFilename, tokenFilename string) *YouTube {
	return &YouTube{
		url:            url,
		auth:           auth,
		configFilename: configFilename,
		tokenFilename:  tokenFilename,
		messageChan:    make(chan Message, 200),
		InsertChan:     make(chan interface{}, 200),
		DeleteChan:     make(chan interface{}, 200),
		fanFunding:     fanFunding{Messages: make(map[string]*youtube.LiveChatMessage)},
		joined:         make(map[string]bool),
	}
}

func (yt *YouTube) pollBroadcasts(broadcasts *youtube.LiveBroadcastListResponse, err error) {
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

func (yt *YouTube) pollMessages(broadcast *youtube.LiveBroadcast) {
	if yt.joined[broadcast.Snippet.LiveChatId] {
		return
	}
	yt.joined[broadcast.Snippet.LiveChatId] = true

	yt.channelCount++
	pageToken := ""
	for {
		list := yt.Service.LiveChatMessages.List(broadcast.Snippet.LiveChatId, "id,snippet,authorDetails").MaxResults(200)
		if pageToken != "" {
			list.PageToken(pageToken)
		}

		liveChatMessageListResponse, err := list.Do()

		if err != nil {
			log.Println(err)
		} else {
			// Ignore the first results, we only want new chats.
			if pageToken != "" {
				for _, message := range liveChatMessageListResponse.Items {
					liveChatMessage := LiveChatMessage(*message)
					yt.messageChan <- &liveChatMessage

					switch message.Snippet.Type {
					case LiveChatMessageSnippetTypeFanFunding:
						yt.addFanFundingMessage(message)
					}
				}
			}
			pageToken = liveChatMessageListResponse.NextPageToken
		}

		if liveChatMessageListResponse != nil && liveChatMessageListResponse.PollingIntervalMillis != 0 {
			time.Sleep(time.Duration(liveChatMessageListResponse.PollingIntervalMillis) * time.Millisecond)
		} else {
			time.Sleep(10 * time.Second)
		}
	}
}

func (yt *YouTube) writeMessagesToFile(messages []*youtube.LiveChatMessage, filename string) {
	output := ""
	for _, message := range messages {
		output += html.UnescapeString(message.Snippet.DisplayMessage) + "\n"
	}
	err := ioutil.WriteFile(filename, []byte(output), 0777)
	if err != nil {
		log.Println(err)
	}
}

func (yt *YouTube) addFanFundingMessage(message *youtube.LiveChatMessage) {
	yt.fanFunding.Lock()
	defer yt.fanFunding.Unlock()

	if yt.fanFunding.Messages[message.Id] == nil {
		yt.fanFunding.Messages[message.Id] = message
		yt.writeMessagesToFile([]*youtube.LiveChatMessage{message}, "youtubelatest.txt")
	}

	largest := message
	for _, check := range yt.fanFunding.Messages {
		if check.Snippet.FanFundingEventDetails.AmountMicros > largest.Snippet.FanFundingEventDetails.AmountMicros {
			largest = check
		}
	}

	yt.writeMessagesToFile([]*youtube.LiveChatMessage{largest}, "youtubelargest.txt")
}

func (yt *YouTube) generateOauthURLAndExit() {
	// Redirect user to Google's consent page to ask for permission
	// for the scopes specified above.
	url := yt.config.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	clipboard.WriteAll(url)
	log.Fatalln("Visit the following URL to generate an auth code, then rerun with -auth=<code> (It has also been copied to your clipboard):\n%s", url)
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
	channelInsertChans := make(map[string]chan *youtube.LiveChatMessage)

	// Chat messages need to be separated by one second, they must be handled by a separate goroutine.
	insertLiveChatMessageLimited := func(liveChatMessage *youtube.LiveChatMessage) {
		channelInsertChan := channelInsertChans[liveChatMessage.Snippet.LiveChatId]
		if channelInsertChan == nil {
			channelInsertChan = make(chan *youtube.LiveChatMessage, 200)
			channelInsertChans[liveChatMessage.Snippet.LiveChatId] = channelInsertChan

			// Start a goroutine to rate limit sends.
			go func() {
				for {
					if _, err := yt.Service.LiveChatMessages.Insert("snippet", <-channelInsertChan).Do(); err != nil {
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
			case *youtube.LiveChatMessage:
				insertLiveChatMessageLimited(request)
			case *youtube.LiveChatBan:
				yt.Service.LiveChatBans.Insert("snippet", request).Do()
			case *youtube.LiveChatModerator:
				yt.Service.LiveChatModerators.Insert("snippet", request).Do()
			}
		case request := <-yt.DeleteChan:
			switch request := request.(type) {
			case *youtube.LiveChatMessage:
				yt.Service.LiveChatMessages.Delete(request.Id).Do()
			case *youtube.LiveChatBan:
				yt.Service.LiveChatBans.Delete(request.Id).Do()
			case *youtube.LiveChatModerator:
				yt.Service.LiveChatModerators.Delete(request.Id).Do()
			}
		}
		if err != nil {
			log.Println(err)
		}

		// Sleep for a millisecond, this will guarantee a maximum QPS of 1000.
		time.Sleep(1 * time.Millisecond)
	}
}

func (yt *YouTube) Init() error {
	if err := yt.createConfig(); err != nil {
		return err
	}

	// An oauth URL was requested, error early.
	if yt.url {
		yt.generateOauthURLAndExit()
	}

	if err := yt.createToken(); err != nil {
		return err
	}

	yt.Client = yt.config.Client(oauth2.NoContext, yt.token)

	if service, err := youtube.New(yt.Client); err == nil {
		yt.Service = service
	} else {
		return err
	}

	return nil
}

// Open opens the service and returns a channel which all messages will be sent on.
func (yt *YouTube) Open() (<-chan Message, error) {
	me, err := yt.GetMe()
	if err != nil {
		return nil, err
	}
	yt.me = me

	yt.pollBroadcasts(yt.Service.LiveBroadcasts.List("id,snippet,status,contentDetails").Mine(true).BroadcastType("all").Do())

	// Start a goroutine to handle all requests.
	go yt.handleRequests()

	return yt.messageChan, nil
}

// IsMe returns whether or not a message was sent by the bot.
func (yt *YouTube) IsMe(message Message) bool {
	return message.UserID() == yt.me.Id
}

var messageReplacer = strings.NewReplacer("<", "(", ">", ")")

// SendMessage sends a message.
func (yt *YouTube) SendMessage(channel, message string) error {
	yt.InsertChan <- &youtube.LiveChatMessage{
		Snippet: &youtube.LiveChatMessageSnippet{
			LiveChatId: channel,
			Type:       LiveChatMessageSnippetTypeText,
			TextMessageDetails: &youtube.LiveChatTextMessageDetails{
				MessageText: messageReplacer.Replace(message),
			},
		},
	}
	return nil
}

// DeleteMessage deletes a message.
func (yt *YouTube) DeleteMessage(channel, messageID string) error {
	yt.DeleteChan <- &youtube.LiveChatMessage{Id: messageID}
	return nil
}

// SendFile sends a file.
func (yt *YouTube) SendFile(channel, name string, r io.Reader) error {
	return errors.New("Send file not supported.")
}

// BanUser bans a user.
func (yt *YouTube) BanUser(channel, userID string, duration int) error {
	liveChatBan := &youtube.LiveChatBan{
		Snippet: &youtube.LiveChatBanSnippet{
			LiveChatId: channel,
			BannedUserDetails: &youtube.ChannelProfileDetails{
				ChannelId: userID,
			},
		},
	}

	if duration == -1 {
		liveChatBan.Snippet.Type = LiveChatBanSnippetTypePermanent
	} else {
		liveChatBan.Snippet.Type = LiveChatBanSnippetTypeTemporary
		liveChatBan.Snippet.BanDurationSeconds = uint64(duration)
	}

	yt.InsertChan <- liveChatBan
	return nil
}

// UnbanUser unbans a user.
func (yt *YouTube) UnbanUser(channel, userID string) error {
	return errors.New("Unbanning user not supported on YouTube.")
}

// UserName returns the bots name.
func (yt *YouTube) UserName() string {
	return yt.me.Snippet.Title
}

// UserID returns the bots user id.
func (yt *YouTube) UserID() string {
	return yt.me.Id
}

// PrivateMessage will send a private message to a user.
func (yt *YouTube) PrivateMessage(userID, message string) error {
	return errors.New("Private messages not supported on YouTube.")
}

// Join will join a channel.
func (yt *YouTube) Join(join string) error {
	if yt.joined[join] {
		return nil
	}
	yt.joined[join] = true

	videos, err := yt.GetVideosByIDList([]string{join})

	if err != nil {
		return errors.New("No live video found.")
	}

	ok := false

	items := []*youtube.LiveBroadcast{}
	for _, v := range videos {
		if v.LiveStreamingDetails != nil && v.LiveStreamingDetails.ActiveLiveChatId != "" {
			items = append(items, &youtube.LiveBroadcast{
				Snippet: &youtube.LiveBroadcastSnippet{
					LiveChatId: v.LiveStreamingDetails.ActiveLiveChatId,
				},
			})
			ok = true
		}
	}

	if !ok {
		return errors.New("No live video found.")
	}

	liveBroadcastListResponse := &youtube.LiveBroadcastListResponse{
		Items: items,
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

// Typing sets that the bot is typing.
func (yt *YouTube) Typing(channel string) error {
	return errors.New("Tying is not supported on YouTube.")
}

// SupportsPrivateMessages returns whether the service supports private messages.
func (yt *YouTube) SupportsPrivateMessages() bool {
	return false
}

// SupportsMultiline returns whether the service supports multiline messages.
func (yt *YouTube) SupportsMultiline() bool {
	return false
}

// CommandPrefix returns the command prefix for the service.
func (yt *YouTube) CommandPrefix() string {
	return fmt.Sprintf("@%s ", yt.UserName())
}

// IsBotOwner returns whether or not a message sender was the owner of the bot.
func (yt *YouTube) IsBotOwner(message Message) bool {
	return false
}

// IsPrivate returns whether or not a message was private.
func (yt *YouTube) IsPrivate(message Message) bool {
	return false
}

// IsModerator returns whether or not the sender of a message is a moderator.
func (yt *YouTube) IsModerator(message Message) bool {
	m, ok := message.(*LiveChatMessage)
	if !ok {
		return false
	}
	return m.AuthorDetails.IsChatOwner || m.AuthorDetails.IsChatModerator
}

// ChannelCount returns the number of channels the bot is in.
func (yt *YouTube) ChannelCount() int {
	return yt.channelCount
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
		return nil, err
	}

	videos, err := yt.GetVideosByIDList(ids)
	if err != nil {
		return nil, err
	}

	videoList := videoList(videos)
	sort.Sort(videoList)

	return videoList, nil
}

// SupportsMessageHistory returns if the service supports message history.
func (yt *YouTube) SupportsMessageHistory() bool {
	return false
}

// MessageHistory returns the message history for a channel.
func (yt *YouTube) MessageHistory(channel string) []Message {
	return nil
}

// GetLiveVideos gets a list of live videos for a channel
func (yt *YouTube) GetLiveVideos(channelID string) ([]*youtube.Video, error) {
	if yt.Service == nil {
		return nil, errors.New("Service not available.")
	}

	s, err := yt.Service.Search.List("id").ChannelId(channelID).EventType("live").Type("video").Do()
	if err != nil {
		return nil, err
	}

	ids := []string{}
	for _, sr := range s.Items {
		ids = append(ids, sr.Id.VideoId)
	}

	videos, err := yt.GetVideosByIDList(ids)
	if err != nil {
		return nil, err
	}

	liveVideos := []*youtube.Video{}

	for _, v := range videos {
		if v.LiveStreamingDetails.ActualEndTime == "" {
			liveVideos = append(liveVideos, v)
		}
	}

	return liveVideos, nil
}
