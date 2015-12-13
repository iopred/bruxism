package main

import (
  "bytes"
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

type YouTube struct {
  config      *oauth2.Config
  token       *oauth2.Token
  Client      *http.Client
  MessageChan chan Message
  InsertChan  chan interface{}
  DeleteChan  chan interface{}
  fanFunding  FanFunding
  me          string
}

func NewYouTube() *YouTube {
  return &YouTube{
    MessageChan: make(chan Message, 200),
    InsertChan:  make(chan interface{}, 200),
    DeleteChan:  make(chan interface{}, 200),
    fanFunding:  FanFunding{Messages: make(map[string]*LiveChatMessage)},
  }
}

func (yt *YouTube) GetMe() (string, error) {
  resp, err := yt.Client.Get("https://www.googleapis.com/youtube/v3/channels?part=id&mine=true")
  if err != nil {
    return "", err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  type ChannelListResponse struct {
    Items []struct {
      Id string `json:"id"`
    } `json:"items"`
  }

  channelList := &ChannelListResponse{}
  err = json.Unmarshal(body, channelList)

  if len(channelList.Items) != 1 {
    return "", errors.New("Invalid response while requesting Me")
  }

  return channelList.Items[0].Id, nil
}

func (yt *YouTube) pollBroadcasts(broadcasts []*LiveBroadcast, err error) {
  if err != nil {
    log.Println(err)
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

func (yt *YouTube) pollMessages(broadcast *LiveBroadcast) {
  pageToken := ""
  for {
    liveChatMessageListResponse, err := yt.ListLiveChatMessages(broadcast.Snippet.LiveChatId, pageToken)

    if err != nil {
      log.Println(err)
    } else if liveChatMessageListResponse.Error != nil {
      log.Println(liveChatMessageListResponse.Error.NewError("polling messages"))
    } else {
      // Ignore the first results, we only want new chats.
      if pageToken != "" {
        for i, message := range liveChatMessageListResponse.Items {
          // Ignore messages from ourselves.
          if message.AuthorDetails.ChannelId == yt.me {
            continue
          }
          yt.MessageChan <- message
          switch message.Snippet.Type {
          case LiveChatMessageSnippetTypeFanFunding:
            yt.addFanFundingMessage(liveChatMessageListResponse.Items[i])
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
  yt.fanFunding.Lock()
  defer yt.fanFunding.Unlock()

  if yt.fanFunding.Messages[message.Id] == nil {
    yt.fanFunding.Messages[message.Id] = message
    yt.writeMessagesToFile([]*LiveChatMessage{message}, "youtubelatest.txt")
  }

  largest := message
  for _, check := range yt.fanFunding.Messages {
    if check.Snippet.FanFundingEventDetails.AmountMicros > largest.Snippet.FanFundingEventDetails.AmountMicros {
      largest = check
    }
  }

  yt.writeMessagesToFile([]*LiveChatMessage{largest}, "youtubelargest.txt")
}

func (yt *YouTube) generateOauthUrl() string {
  // Redirect user to Google's consent page to ask for permission
  // for the scopes specified above.
  url := yt.config.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
  clipboard.WriteAll(url)
  return fmt.Sprintf("Visit the following URL to generate an auth code, then rerun with -auth=<code> (It has also been copied to your clipboard):\n%v", url)
}

func (yt *YouTube) createConfig() error {
  b, err := ioutil.ReadFile(config)
  if err != nil {
    return err
  }

  yt.config, err = google.ConfigFromJSON(b, "https://www.googleapis.com/auth/youtube")
  return err
}

func (yt *YouTube) createToken() error {
  if auth != "" {
    // Let's exchange our code
    tok, err := yt.config.Exchange(oauth2.NoContext, auth)
    if err != nil {
      return err
    }
    yt.token = tok

    b, err := json.Marshal(yt.token)
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
      yt.token = &oauth2.Token{}
      // A token was found, unmarshall it and use it.
      err := json.Unmarshal(b, yt.token)
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

func (yt *YouTube) Delete(url string) (resp *http.Response, err error) {
  req, err := http.NewRequest("DELETE", url, nil)
  if err != nil {
    return nil, err
  }
  return yt.Client.Do(req)
}

func (yt *YouTube) ListLiveBroadcasts(params string) ([]*LiveBroadcast, error) {
  resp, err := yt.Client.Get("https://www.googleapis.com/youtube/v3/liveBroadcasts?part=id,snippet,status,contentDetails&" + params)
  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  broadcasts := &LiveBroadcastListResponse{}
  err = json.Unmarshal(body, broadcasts)
  if err != nil {
    return nil, err
  }

  if broadcasts.Error != nil {
    return nil, broadcasts.Error.NewError("getting broadcasts")
  }

  return broadcasts.Items, nil
}

func (yt *YouTube) ListLiveChatMessages(liveChatId string, pageToken string) (*LiveChatMessageListResponse, error) {
  pageTokenString := ""
  if pageToken != "" {
    pageTokenString = "&pageToken=" + pageToken
  }

  resp, err := yt.Client.Get("https://www.googleapis.com/youtube/v3/liveChat/messages?maxResults=50&part=id,snippet,authorDetails&liveChatId=" + liveChatId + pageTokenString)
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

func (yt *YouTube) InsertLiveChatMessage(liveChatMessage *LiveChatMessage) error {
  jsonString, err := json.Marshal(liveChatMessage)
  if err != nil {
    return err
  }

  resp, err := yt.Client.Post("https://www.googleapis.com/youtube/v3/liveChat/messages?part=snippet", "application/json", bytes.NewBuffer(jsonString))
  if err != nil {
    return err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }

  liveChatMessage = &LiveChatMessage{}
  err = json.Unmarshal(body, liveChatMessage)
  if err != nil {
    return err
  }

  if liveChatMessage.Error != nil {
    return liveChatMessage.Error.NewError("inserting LiveChatMessage")
  }

  return nil
}

func (yt *YouTube) DeleteLiveChatMessage(liveChatMessage *LiveChatMessage) error {
  resp, err := yt.Delete("https://www.googleapis.com/youtube/v3/liveChat/messages?id=" + liveChatMessage.Id)
  if err != nil {
    return err
  }
  return resp.Body.Close()
}

func (yt *YouTube) InsertLiveChatBan(liveChatBan *LiveChatBan) error {
  jsonString, err := json.Marshal(liveChatBan)
  if err != nil {
    return err
  }

  resp, err := yt.Client.Post("https://www.googleapis.com/youtube/v3/liveChatBans?part=snippet", "application/json", bytes.NewBuffer(jsonString))
  if err != nil {
    return err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }

  liveChatBan = &LiveChatBan{}
  err = json.Unmarshal(body, liveChatBan)
  if err != nil {
    return err
  }

  if liveChatBan.Error != nil {
    return liveChatBan.Error.NewError("inserting LiveChatBan")
  }

  return nil
}

func (yt *YouTube) DeleteLiveChatBan(liveChatBan *LiveChatBan) error {
  resp, err := yt.Delete("https://www.googleapis.com/youtube/v3/liveChatBans?id=" + liveChatBan.Id)
  if err != nil {
    return err
  }
  return resp.Body.Close()
}

func (yt *YouTube) ListLiveChatModerators(liveChatId string, pageToken string) (*LiveChatModeratorListResponse, error) {
  pageTokenString := ""
  if pageToken != "" {
    pageTokenString = "&pageToken=" + pageToken
  }

  resp, err := yt.Client.Get("https://www.googleapis.com/youtube/v3/liveChatModerators?maxResults=50&part=id,snippet&liveChatId=" + liveChatId + pageTokenString)
  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  liveChatModeratorListResponse := &LiveChatModeratorListResponse{}
  err = json.Unmarshal(body, liveChatModeratorListResponse)
  if err != nil {
    return nil, err
  }

  return liveChatModeratorListResponse, nil
}

func (yt *YouTube) InsertLiveChatModerator(liveChatModerator *LiveChatModerator) error {
  jsonString, err := json.Marshal(liveChatModerator)
  if err != nil {
    return err
  }

  resp, err := yt.Client.Post("https://www.googleapis.com/youtube/v3/liveChatModerators?part=snippet", "application/json", bytes.NewBuffer(jsonString))
  if err != nil {
    return err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }

  liveChatModerator = &LiveChatModerator{}
  err = json.Unmarshal(body, liveChatModerator)
  if err != nil {
    return err
  }

  if liveChatModerator.Error != nil {
    return liveChatModerator.Error.NewError("inserting LiveChatModerator")
  }

  return nil
}

func (yt *YouTube) DeleteLiveChatModerator(liveChatModerator *LiveChatModerator) error {
  resp, err := yt.Delete("https://www.googleapis.com/youtube/v3/liveChatModerators?id=" + liveChatModerator.Id)
  if err != nil {
    return err
  }
  return resp.Body.Close()
}

func (yt *YouTube) Name() string {
  return "YouTube"
}

func (yt *YouTube) MessageChannel() <-chan Message {
  return yt.MessageChan
}

func (yt *YouTube) Open() error {
  if err := yt.createConfig(); err != nil {
    return err
  }

  // An oauth URL was requested, error early.
  if url {
    return errors.New(yt.generateOauthUrl())
  }

  if err := yt.createToken(); err != nil {
    return err
  }

  yt.Client = yt.config.Client(oauth2.NoContext, yt.token)

  me, err := yt.GetMe()
  if err != nil {
    return err
  }
  yt.me = me

  yt.pollBroadcasts(yt.ListLiveBroadcasts("default=true"))
  yt.pollBroadcasts(yt.ListLiveBroadcasts("mine=true"))

  // This is a map of channel id's to channels, it is used to send messages to a goroutine that is rate limiting each chatroom.
  channelInsertChans := make(map[string]chan *LiveChatMessage)

  // Chat messages need to be separated by one second, they must be handled by a separate goroutine.
  insertLiveChatMessageLimited := func(liveChatMessage *LiveChatMessage) {
    channelInsertChan := channelInsertChans[liveChatMessage.Channel()]
    if channelInsertChan == nil {
      channelInsertChan = make(chan *LiveChatMessage, 200)
      channelInsertChans[liveChatMessage.Channel()] = channelInsertChan

      // Start a goroutine to rate limit sends.
      go func() {
        for {
          if err := yt.InsertLiveChatMessage(<-channelInsertChan); err != nil {
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
        case *LiveChatMessage:
          insertLiveChatMessageLimited(request)
        case *LiveChatBan:
          yt.InsertLiveChatBan(request)
        case *LiveChatModerator:
          yt.InsertLiveChatModerator(request)
        }
      case request := <-yt.DeleteChan:
        switch request := request.(type) {
        case *LiveChatMessage:
          yt.DeleteLiveChatMessage(request)
        case *LiveChatBan:
          yt.DeleteLiveChatBan(request)
        case *LiveChatModerator:
          yt.DeleteLiveChatModerator(request)
        }
      }

      // Sleep for a millisecond, this will guarantee a maximum QPS of 1000.
      time.Sleep(1 * time.Millisecond)
    }
  }()

  return nil
}

func (yt *YouTube) SendMessage(channel, message string) error {
  yt.InsertChan <- NewLiveChatMessage(channel, message)
  return nil
}

func (yt *YouTube) DeleteMessage(messageId string) error {
  yt.DeleteChan <- &LiveChatMessage{Id: messageId}
  return nil
}

func (yt *YouTube) BanUser(channel, user string, duration int) error {
  yt.InsertChan <- NewLiveChatBan(channel, user, duration)
  return nil
}
