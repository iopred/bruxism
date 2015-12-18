package main

import (
  "encoding/json"
  "fmt"
  "html"
  "io/ioutil"
  "log"
  "sort"
  "strconv"
  "strings"
  "sync"
  "time"

  "github.com/atotto/clipboard"
  ytc "github.com/iopred/ytlivechatapi"

  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
)

const YouTubeServiceName string = "YouTube"

type LiveChatMessage ytc.LiveChatMessage

func (m *LiveChatMessage) Channel() string {
  return m.Snippet.LiveChatId
}

func (m *LiveChatMessage) UserName() string {
  return m.AuthorDetails.DisplayName
}

func (m *LiveChatMessage) UserId() string {
  return m.AuthorDetails.ChannelId
}

func (m *LiveChatMessage) Message() string {
  switch m.Snippet.Type {
  case ytc.LiveChatMessageSnippetTypeText:
    return html.UnescapeString(m.Snippet.TextMessageDetails.MessageText)
  }
  return html.UnescapeString(m.Snippet.DisplayMessage)
}

func (m *LiveChatMessage) MessageId() string {
  return m.Id
}

func (m *LiveChatMessage) IsModerator() bool {
  return m.AuthorDetails.IsChatOwner || m.AuthorDetails.IsChatModerator
}

type FanFunding struct {
  sync.Mutex
  Messages map[string]*ytc.LiveChatMessage
}

type YouTube struct {
  url            bool
  auth           string
  configFilename string
  tokenFilename  string
  liveChatIds    string
  config         *oauth2.Config
  token          *oauth2.Token
  Client         *ytc.Client
  messageChan    chan Message
  InsertChan     chan interface{}
  DeleteChan     chan interface{}
  fanFunding     FanFunding
  me             string
}

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
    fanFunding:     FanFunding{Messages: make(map[string]*ytc.LiveChatMessage)},
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

func (yt *YouTube) generateOauthUrlAndExit() {
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
    } else {
      err := ioutil.WriteFile(yt.tokenFilename, b, 0777)
      if err != nil {
        return err
      }
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
      yt.generateOauthUrlAndExit()
    }
  }
  return nil
}

func (yt *YouTube) Name() string {
  return YouTubeServiceName
}

func (yt *YouTube) Open() (<-chan Message, error) {
  if err := yt.createConfig(); err != nil {
    return nil, err
  }

  // An oauth URL was requested, error early.
  if yt.url {
    yt.generateOauthUrlAndExit()
  }

  if err := yt.createToken(); err != nil {
    return nil, err
  }

  yt.Client = ytc.NewClient(yt.config.Client(oauth2.NoContext, yt.token))

  me, err := yt.Client.GetMe()
  if err != nil {
    return nil, err
  }
  yt.me = me

  yt.pollBroadcasts(yt.Client.ListLiveBroadcasts("default=true"))
  yt.pollBroadcasts(yt.Client.ListLiveBroadcasts("mine=true"))
  if yt.liveChatIds != "" {
    liveChatIdsArray := strings.Split(yt.liveChatIds, ",")

    additionalBroadcasts := &ytc.LiveBroadcastListResponse{
      Items: make([]*ytc.LiveBroadcast, 0),
    }

    for _, liveChatId := range liveChatIdsArray {
      additionalBroadcasts.Items = append(additionalBroadcasts.Items, &ytc.LiveBroadcast{
        Snippet: &ytc.LiveBroadcastSnippet{
          LiveChatId: liveChatId,
        },
      })
    }

    yt.pollBroadcasts(additionalBroadcasts, nil)
  }

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

  // Start a goroutine to handle all requests.
  go func() {
    for {
      select {
      case request := <-yt.InsertChan:
        switch request := request.(type) {
        case *ytc.LiveChatMessage:
          insertLiveChatMessageLimited(request)
        case *ytc.LiveChatBan:
          if err := yt.Client.InsertLiveChatBan(request); err != nil {
            log.Println(err)
          }
        case *ytc.LiveChatModerator:
          if err := yt.Client.InsertLiveChatModerator(request); err != nil {
            log.Println(err)
          }
        }
      case request := <-yt.DeleteChan:
        switch request := request.(type) {
        case *ytc.LiveChatMessage:
          if err := yt.Client.DeleteLiveChatMessage(request); err != nil {
            log.Println(err)
          }
        case *ytc.LiveChatBan:
          if err := yt.Client.DeleteLiveChatBan(request); err != nil {
            log.Println(err)
          }
        case *ytc.LiveChatModerator:
          if err := yt.Client.DeleteLiveChatModerator(request); err != nil {
            log.Println(err)
          }
        }
      }

      // Sleep for a millisecond, this will guarantee a maximum QPS of 1000.
      time.Sleep(1 * time.Millisecond)
    }
  }()

  return yt.messageChan, nil
}

func (yt *YouTube) IsMe(message Message) bool {
  return message.UserId() == yt.me
}

func (yt *YouTube) SendMessage(channel, message string) error {
  yt.InsertChan <- ytc.NewLiveChatMessage(channel, message)
  return nil
}

func (yt *YouTube) DeleteMessage(messageId string) error {
  yt.DeleteChan <- &ytc.LiveChatMessage{Id: messageId}
  return nil
}

func (yt *YouTube) BanUser(channel, user string, duration int) error {
  yt.InsertChan <- ytc.NewLiveChatBan(channel, user, duration)
  return nil
}

type PlaylistItemSnippetThumbnails struct {
  Url    string `json:"url,omitempty"`
  Width  int    `json:"width,omitempty"`
  Height int    `json:"height,omitempty"`
}

type PlaylistItemSnippetResourceId struct {
  Kind    string `json="kind,omitempty"`
  VideoId string `json="videoId,omitempty"`
}

type PlaylistItemSnippet struct {
  PublishedAt  string                                    `json="publishedAt,omitempty"`
  ChannelId    string                                    `json="channelId,omitempty"`
  Title        string                                    `json="title,omitempty"`
  Description  string                                    `json="description,omitempty"`
  Thumbnails   map[string]*PlaylistItemSnippetThumbnails `json:"thumbnails,omitempty,omitempty"`
  ChannelTitle string                                    `json="channelTitle,omitempty"`
  PlaylistId   string                                    `json="playlistId,omitempty"`
  Position     int                                       `json="position,omitempty"`
  ResourceId   *PlaylistItemSnippetResourceId            `json="resourceId,omitempty"`
}

type PlaylistItemContentDetails struct {
  VideoId string `json="videoId,omitempty"`
  StartAt string `json="startAt,omitempty"`
  EndAt   string `json="endAt,omitempty"`
  Note    string `json="note,omitempty"`
}

type PlaylistItemStatus struct {
  PrivacyStatus string `json="privacyStatus,omitempty"`
}

const PlaylistItemKind string = "youtube#playlistItem"

type PlaylistItem struct {
  Error          *ytc.Error                  `json="error,omitempty"`
  Kind           string                      `json="kind,omitempty"`
  Etag           string                      `json="etag,omitempty"`
  Id             string                      `json="id,omitempty"`
  Snippet        *PlaylistItemSnippet        `json="snippet,omitempty"`
  ContentDetails *PlaylistItemContentDetails `json="contentDetails,omitempty"`
  Status         *PlaylistItemStatus         `json="status,omitempty"`
}

const PlaylistItemListResponseKing string = "youtube#playlistItemListResponse"

type PlaylistItemListResponse struct {
  Error         *ytc.Error      `json="error,omitempty"`
  Kind          string          `json="kind,omitempty"`
  Etag          string          `json="etag,omitempty"`
  NextPageToken string          `json="nextPageToken,omitempty"`
  PrevPageToken string          `json="prevPageToken,omitempty"`
  PageInfo      *ytc.PageInfo   `json="pageInfo,omitempty"`
  Items         []*PlaylistItem `json="items,omitempty"`
}

type VideoSnippet struct {
  PublishedAt  string `json="publishedAt,omitempty"`
  ChannelId    string `json="channelId,omitempty"`
  Title        string `json="title,omitempty"`
  Description  string `json="description,omitempty"`
  ChannelTitle string `json="channelTitle,omitempty"`
}

type VideoLiveStreamingDetails struct {
  ActualStartTime          string `json="actualStartTime,omitempty"`
  ActualEndTime            string `json="actualEndTime,omitempty"`
  ScheduledStartTime       string `json="scheduledStartTime,omitempty"`
  ScheduledEndTime         string `json="scheduledEndTime,omitempty"`
  ConcurrentViewers        string `json="concurrentViewers,omitempty"`
  ConcurrentViewersInteger int
}

const VideoKind string = "youtube#video"

type Video struct {
  Error                *ytc.Error                 `json="error,omitempty"`
  Kind                 string                     `json="kind,omitempty"`
  Etag                 string                     `json="etag,omitempty"`
  Id                   string                     `json="id,omitempty"`
  Snippet              *VideoSnippet              `json="snippet,omitempty"`
  LiveStreamingDetails *VideoLiveStreamingDetails `json="liveStreamingDetails,omitempty"`
}

const VideoListResponseKind string = "youtube#videoListResponse"

type VideoListResponse struct {
  Kind          string        `json="kind,omitempty"`
  Etag          string        `json="etag,omitempty"`
  NextPageToken string        `json="nextPageToken,omitempty"`
  PrevPageToken string        `json="prevPageToken,omitempty"`
  PageInfo      *ytc.PageInfo `json="pageInfo,omitempty"`
  Items         []*Video      `json="items,omitempty"`
}

type VideoList []*Video

func (v VideoList) Len() int {
  return len(v)
}

func (v VideoList) Swap(i, j int) {
  v[i], v[j] = v[j], v[i]
}

func (v VideoList) Less(i, j int) bool {
  return v[i].LiveStreamingDetails.ConcurrentViewersInteger > v[j].LiveStreamingDetails.ConcurrentViewersInteger
}

func (yt *YouTube) GetTopLivestreamIds(count int) ([]string, error) {
  ids := make([]string, 0)

  pageTokenString := ""

  for {
    resp, err := yt.Client.Get(fmt.Sprintf("https://www.googleapis.com/youtube/v3/playlistItems?maxResults=50&part=id,contentDetails&playlistId=PLiCvVJzBupKmEehQ3hnNbbfBjLUyvGlqx%v", pageTokenString))
    if err != nil {
      return nil, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return nil, err
    }

    playlistItemListResponse := &PlaylistItemListResponse{}
    err = json.Unmarshal(body, playlistItemListResponse)
    if err != nil {
      return nil, err
    }

    for _, playlistItem := range playlistItemListResponse.Items {
      ids = append(ids, playlistItem.ContentDetails.VideoId)
    }

    if len(ids) >= count || playlistItemListResponse.NextPageToken == "" {
      break
    }

    pageTokenString = "&pageToken=" + playlistItemListResponse.NextPageToken
  }

  return ids, nil
}

func (yt *YouTube) GetVideosByIdList(ids []string) ([]*Video, error) {
  videos := make([]*Video, 0)

  i := 0
  for i < len(ids) {
    next := i + 50
    if next >= len(ids) {
      next = len(ids)
    }
    resp, err := yt.Client.Get(fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?maxResults=50&part=id,snippet,liveStreamingDetails&id=%v", strings.Join(ids[i:next], ",")))

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return nil, err
    }

    videoListResponse := &VideoListResponse{}
    err = json.Unmarshal(body, videoListResponse)
    if err != nil {
      return nil, err
    }

    for _, video := range videoListResponse.Items {
      if i, err := strconv.Atoi(video.LiveStreamingDetails.ConcurrentViewers); err == nil {
        video.LiveStreamingDetails.ConcurrentViewersInteger = i
      }
    }

    videos = append(videos, videoListResponse.Items...)

    if len(videos) != len(ids) || videoListResponse.NextPageToken == "" {
      break
    }
  }

  return videos, nil
}

func (yt *YouTube) GetTopLivestreams(count int) ([]*Video, error) {
  ids, err := yt.GetTopLivestreamIds(count)
  if err != nil {
    return nil, nil
  }

  videos, err := yt.GetVideosByIdList(ids)
  if err != nil {
    return nil, nil
  }

  videoList := VideoList(videos)
  sort.Sort(videoList)

  return videoList, nil
}
