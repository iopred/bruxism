package main

import (
  "encoding/json"
  "errors"
  "flag"
  "fmt"
  "html"
  "io/ioutil"
  "log"
  "net/http"
  "sync"
  "time"

  "github.com/atotto/clipboard"

  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
)

const TEXT_MESSAGE_EVENT string = "textMessageEvent"
const FAN_FUNDING_EVENT string = "fanFundingEvent"

var url bool
var auth string
var config string
var token string

func init() {
  flag.BoolVar(&url, "youtubeurl", false, "Generates a URL that provides an auth code.")
  flag.StringVar(&auth, "youtubeauth", "", "Exchanges the provided auth code for an oauth2 token.")
  flag.StringVar(&config, "youtubeconfig", "youtubeoauth2config.json", "The filename that contains the oauth2 config.")
  flag.StringVar(&token, "youtubetoken", "youtubeoauth2token.json", "The filename to store the oauth2 token.")
  flag.Parse()
}

type FanFunding struct {
  sync.Mutex
  Messages map[string]*LiveChatMessage
}

type Error struct {
  Errors []struct {
    Domain  string `json:"domain"`
    Reason  string `json:"reason"`
    Message string `json:"message"`
  }
  Code    int    `json:"code"`
  Message string `json:"message"`
}

type LiveBroadcast struct {
  Kind    string `json:"kind"`
  Etag    string `json:"etag"`
  Id      string `json:"id"`
  Snippet struct {
    PublishedAt string `json:"publishedAt"`
    ChannelId   string `json:"channelId"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Thumbnails  map[string]struct {
      Url    string `json:"url"`
      Width  int    `json:"width"`
      Height int    `json:"height"`
    } `json:"thumbnails"`
    ScheduledStartTime string `json:"scheduledStartTime"`
    ScheduledEndTime   string `json:"scheduledEndTime"`
    ActualStartTime    string `json:"actualStartTime"`
    ActualEndTime      string `json:"actualEndTime"`
    IsDefaultBroadcast bool   `json:"isDefaultBroadcast"`
    LiveChatId         string `json:"liveChatId"`
  } `json:"snippet"`
  Status struct {
    LifeCycleStatus string `json:"lifeCycleStatus"`
    PrivacyStatus   string `json:"privacyStatus"`
    RecordingStatus string `json:"recordingStatus"`
  } `json:"status"`
  ContentDetails struct {
    BoundStreamId string `json:"boundStreamId"`
    MonitorStream struct {
      EnableMonitorStream    bool   `json:"enableMonitorStream"`
      BroadcastStreamDelayMs uint   `json:"broadcastStreamDelayMs"`
      EmbedHtml              string `json:"embedHtml"`
    } `json:"monitorStream"`
    EnableEmbed             bool `json:"enableEmbed"`
    EnableDvr               bool `json:"enableDvr"`
    EnableContentEncryption bool `json:"enableContentEncryption"`
    StartWithSlate          bool `json:"startWithSlate"`
    RecordFromStart         bool `json:"recordFromStart"`
    EnableClosedCaptions    bool `json:"enableClosedCaptions"`
  } `json:"contentDetails"`
}

type LiveBroadcasts struct {
  Error         *Error `json:"error"`
  Kind          string `json:"kind"`
  Etag          string `json:"etag"`
  NextPageToken string `json:"nextPageToken"`
  PageInfo      struct {
    TotalResults   int `json:"totalResults"`
    ResultsPerPage int `json:"resultsPerPage"`
  } `json:"pageInfo"`
  Items []*LiveBroadcast `json:"items"`
}

type LiveChatMessage struct {
  Kind    string `json:"kind"`
  Etag    string `json:"etag"`
  Id      string `json:"id"`
  Snippet struct {
    Type                   string `json:"type"`
    LiveChatId             string `json:"liveChatId"`
    AuthorChannelId        string `json:"authorChannelId"`
    PublishedAt            string `json:"publishedAt"`
    HasDisplayContent      bool   `json:"hasDisplayContent"`
    DisplayMessage         string `json:"displayMessage"`
    FanFundingEventDetails struct {
      AmountMicros        int    `json:"amountMicros,string"`
      Currency            string `json:"currency"`
      AmountDisplayString string `json:"amountDisplayString"`
      UserComment         string `json:"userComment"`
    } `json:"fanFundingEventDetails"`
    TextMessageDetails struct {
      MessageText string `json:"messageText"`
    } `json:"textMessageDetails"`
  } `json:"snippet"`
  AuthorDetails struct {
    ChannelId       string `json:"channelId"`
    ChannelUrl      string `json:"channelUrl"`
    DisplayName     string `json:"displayName"`
    ProfileImageUrl string `json:"profileImageUrl"`
    IsVerified      bool   `json:"isVerified"`
    IsChatOwner     bool   `json:"isChatOwner"`
    IsChatSponsor   bool   `json:"isChatSponsor"`
    IsChatModerator bool   `json:"isChatModerator"`
  } `json:"authorDetails"`
}

func (m *LiveChatMessage) Channel() string {
  return m.Snippet.LiveChatId
}

func (m *LiveChatMessage) User() string {
  return m.AuthorDetails.DisplayName
}

func (m *LiveChatMessage) Message() string {
  switch m.Snippet.Type {
  case TEXT_MESSAGE_EVENT:
    return html.UnescapeString(m.Snippet.TextMessageDetails.MessageText)
  }
  return html.UnescapeString(m.Snippet.DisplayMessage)
}

type LiveChatMessageListResponse struct {
  Error                 *Error `json:"error"`
  Kind                  string `json:"kind"`
  Etag                  string `json:"etag"`
  NextPageToken         string `json:"nextPageToken"`
  PollingIntervalMillis int    `json:"pollingIntervalMillis"`
  PageInfo              struct {
    TotalResults   int `json:"totalResults"`
    ResultsPerPage int `json:"resultsPerPage"`
  } `json:"pageInfo"`
  Items []*LiveChatMessage `json:"items"`
}

type YouTube struct {
  Config      *oauth2.Config
  Token       *oauth2.Token
  Client      *http.Client
  MessageChan chan Message
  Broadcasts  []*LiveBroadcast
  FanFunding  FanFunding
}

func NewYouTube() *YouTube {
  return &YouTube{
    MessageChan: make(chan Message, 200),
    Broadcasts:  make([]*LiveBroadcast, 0, 10),
    FanFunding:  FanFunding{Messages: make(map[string]*LiveChatMessage)},
  }
}

func (yt *YouTube) getBroadcasts(params string) ([]*LiveBroadcast, error) {
  resp, err := yt.Client.Get("https://www.googleapis.com/youtube/v3/liveBroadcasts?part=snippet,status,contentDetails&" + params)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  broadcasts := &LiveBroadcasts{}
  err = json.Unmarshal(body, broadcasts)
  if err != nil {
    return nil, err
  }

  if broadcasts.Error != nil {
    return nil, errors.New(broadcasts.Error.Message)
  }

  return broadcasts.Items, nil
}

func (yt *YouTube) pollBroadcasts(broadcasts []*LiveBroadcast, err error) {
  if err != nil {
    fmt.Println(err)
    return
  }

  for _, broadcast := range broadcasts {
    // If the broadcast has ended, it can't have a valid chat.
    if broadcast.Status.LifeCycleStatus == "complete" {
      continue
    }

    go yt.pollMessages(broadcast)
  }
}

func (yt *YouTube) getMessages(liveChatId string, pageToken string) (*LiveChatMessageListResponse, error) {
  pageTokenString := ""
  if pageToken != "" {
    pageTokenString = "&pageToken=" + pageToken
  }

  resp, err := yt.Client.Get("https://www.googleapis.com/youtube/v3/liveChat/messages?maxResults=50&part=snippet,authorDetails&liveChatId=" + liveChatId + pageTokenString)
  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  liveChatMessageListResponse := &LiveChatMessageListResponse{}
  err = json.Unmarshal(body, liveChatMessageListResponse)
  if err != nil {
    return nil, err
  }

  return liveChatMessageListResponse, nil
}

func (yt *YouTube) pollMessages(broadcast *LiveBroadcast) {
  pageToken := ""
  for {
    liveChatMessageListResponse, err := yt.getMessages(broadcast.Snippet.LiveChatId, pageToken)

    if err != nil {
      log.Println(err)
    } else if liveChatMessageListResponse.Error != nil {
      log.Println("Error", liveChatMessageListResponse.Error.Message)
    } else {
      for i, message := range liveChatMessageListResponse.Items {
        yt.MessageChan <- message
        switch message.Snippet.Type {
        case FAN_FUNDING_EVENT:
          yt.addFanFundingMessage(liveChatMessageListResponse.Items[i])
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

func (yt *YouTube) writeMessagesToFile(messages []*LiveChatMessage, filename string) {
  output := ""
  for _, message := range messages {
    output += html.UnescapeString(message.Snippet.DisplayMessage) + "\n"
  }
  err := ioutil.WriteFile(filename, []byte(output), 0777)
  if err != nil {
    log.Println(err)
  }
}

func (yt *YouTube) addFanFundingMessage(message *LiveChatMessage) {
  yt.FanFunding.Lock()
  defer yt.FanFunding.Unlock()

  if yt.FanFunding.Messages[message.Id] == nil {
    yt.FanFunding.Messages[message.Id] = message
    yt.writeMessagesToFile([]*LiveChatMessage{message}, "youtubelatest.txt")
  }

  largest := message
  for _, check := range yt.FanFunding.Messages {
    if check.Snippet.FanFundingEventDetails.AmountMicros > largest.Snippet.FanFundingEventDetails.AmountMicros {
      largest = check
    }
  }

  yt.writeMessagesToFile([]*LiveChatMessage{largest}, "youtubelargest.txt")
}

func (yt *YouTube) generateOauthUrl() string {
  // Redirect user to Google's consent page to ask for permission
  // for the scopes specified above.
  url := yt.Config.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
  clipboard.WriteAll(url)
  return fmt.Sprintf("Visit the following URL to generate an auth code, then rerun with -auth=<code> (It has also been copied to your clipboard):\n%v", url)
}

func (yt *YouTube) createConfig() error {
  b, err := ioutil.ReadFile(config)
  if err != nil {
    return err
  }

  yt.Config, err = google.ConfigFromJSON(b, "https://www.googleapis.com/auth/youtube")
  return err
}

func (yt *YouTube) createToken() error {
  if auth != "" {
    // Let's exchange our code
    tok, err := yt.Config.Exchange(oauth2.NoContext, auth)
    if err != nil {
      return err
    }
    yt.Token = tok

    b, err := json.Marshal(yt.Token)
    if err != nil {
      return err
    } else {
      err := ioutil.WriteFile(token, b, 0777)
      if err != nil {
        return err
      }
    }
  } else {
    b, err := ioutil.ReadFile(token)
    if err == nil {
      yt.Token = &oauth2.Token{}
      // A token was found, unmarshall it and use it.
      err := json.Unmarshal(b, yt.Token)
      if err != nil {
        return err
      }
    } else {
      // There was an error with the token, maybe it doesn't exist.
      // If we haven't been given an auth code, we must generate one.
      return errors.New(yt.generateOauthUrl())
    }
  }
  return nil
}

func (yt *YouTube) Name() string {
  return "YouTube"
}

func (yt *YouTube) Open() (chan Message, error) {
  if err := yt.createConfig(); err != nil {
    return nil, err
  }

  // An oauth URL was requested, error early.
  if url {
    return nil, errors.New(yt.generateOauthUrl())
  }

  if err := yt.createToken(); err != nil {
    return nil, err
  }

  yt.Client = yt.Config.Client(oauth2.NoContext, yt.Token)

  yt.pollBroadcasts(yt.getBroadcasts("default=true"))
  yt.pollBroadcasts(yt.getBroadcasts("mine=true"))

  return yt.MessageChan, nil
}

func (yt *YouTube) Send(channel, message string) error {
  return errors.New("Sending not supported.")
}
