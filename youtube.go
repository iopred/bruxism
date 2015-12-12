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
  Config      *oauth2.Config
  Token       *oauth2.Token
  Client      *http.Client
  MessageChan chan Message
  Broadcasts  []*LiveBroadcast
  FanFunding  FanFunding
  Me          string
}

func NewYouTube() *YouTube {
  return &YouTube{
    MessageChan: make(chan Message, 200),
    Broadcasts:  make([]*LiveBroadcast, 0, 10),
    FanFunding:  FanFunding{Messages: make(map[string]*LiveChatMessage)},
  }
}

func (yt *YouTube) getMe() error {
  resp, err := yt.Client.Get("https://www.googleapis.com/youtube/v3/channels?part=id&mine=true")
  if err != nil {
    return err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }

  type ChannelListResponse struct {
    Items []struct {
      Id string `json:"id"`
    } `json:"items"`
  }

  channelList := &ChannelListResponse{}
  err = json.Unmarshal(body, channelList)

  if len(channelList.Items) != 1 {
    return errors.New("Invalid response while requesting Me")
  }

  yt.Me = channelList.Items[0].Id

  return nil
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
    return nil, broadcasts.Error.NewError("getting broadcasts")
  }

  return broadcasts.Items, nil
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
      log.Println(liveChatMessageListResponse.Error.NewError("polling messages"))
    } else {
      // Ignore the first results, we only want new chats.
      if pageToken != "" {
        for i, message := range liveChatMessageListResponse.Items {
          // Ignore messages from ourselves.
          if message.AuthorDetails.ChannelId == yt.Me {
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

func (yt *YouTube) MessageChannel() chan Message {
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

  yt.Client = yt.Config.Client(oauth2.NoContext, yt.Token)

  if err := yt.getMe(); err != nil {
    return err
  }

  yt.pollBroadcasts(yt.getBroadcasts("default=true"))
  yt.pollBroadcasts(yt.getBroadcasts("mine=true"))

  return nil
}

func (yt *YouTube) Send(channel, message string) error {
  liveChatMessage := &LiveChatMessage{
    Kind: LiveChatMessageKind,
    Snippet: &LiveChatMessageSnippet{
      LiveChatId: channel,
      Type:       LiveChatMessageSnippetTypeText,
      TextMessageDetails: &LiveChatMessageSnippetTextMessageDetails{
        MessageText: message,
      },
    },
  }

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
    return liveChatMessage.Error.NewError("sending message")
  }

  return nil
}
