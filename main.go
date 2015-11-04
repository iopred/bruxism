package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "html"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "os/signal"
  "strconv"
  "sync"
  "time"

  "github.com/atotto/clipboard"

  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
)

const TEXT_MESSAGE_EVENT string = "text"
const FAN_FUNDING_EVENT string = "fanFundingEvent"

var url bool
var auth string

var config string
var token string

type FanFunding struct {
  sync.Mutex
  Messages map[string]*LiveChatMessage
}

var fanFunding FanFunding

func init() {
  flag.BoolVar(&url, "url", false, "Generates a URL that provides an auth code.")
  flag.StringVar(&auth, "auth", "", "Exchanges the provided auth code for an oauth2 token.")
  flag.StringVar(&config, "config", "oauth2config.json", "The filename that contains the oauth2 config.")
  flag.StringVar(&token, "token", "oauth2token.json", "The filename to store the oauth2 token.")
  flag.Parse()

  fanFunding = FanFunding{Messages: make(map[string]*LiveChatMessage)}
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
  Kind          string `json:"kind"`
  Etag          string `json:"etag"`
  NextPageToken string `json:"nextPageToken"`
  PageInfo      struct {
    TotalResults   int `json:"totalResults"`
    ResultsPerPage int `json:"resultsPerPage"`
  } `json:"pageInfo"`
  Items []LiveBroadcast `json:"items"`
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
    FallbackText           string `json:"fallbackText"`
    MessageText            string `json:"messageText"`
    HasDisplayContent      bool   `json:"hasDisplayContent"`
    FanFundingEventDetails struct {
      AmountMicros  string `json:"amountMicros"`
      Currency      string `json:"currency"`
      DisplayString string `json:"displayString"`
    } `json:"FanFundingEventDetails"`
  } `json:"snippet"`
  AuthorDetails struct {
    ChannelId       string `json:"channelId"`
    DisplayName     string `json:"displayName"`
    ProfileImageUrl string `json:"profileImageUrl"`
    IsVerified      bool   `json:"isVerified"`
    IsChatOwner     bool   `json:"isChatOwner"`
    IsChatModerator bool   `json:"isChatModerator"`
  } `json:"authorDetails"`
}

type LiveChatMessageListResponse struct {
  Kind          string `json:"kind"`
  Etag          string `json:"etag"`
  NextPageToken string `json:"nextPageToken"`
  PageInfo      struct {
    TotalResults   int `json:"totalResults"`
    ResultsPerPage int `json:"resultsPerPage"`
  } `json:"pageInfo"`
  Items []LiveChatMessage `json:"items"`
}

func getBroadcasts(client *http.Client, params string) []LiveBroadcast {
  resp, err := client.Get("https://www.googleapis.com/youtube/v3/liveBroadcasts?part=snippet,status,contentDetails&" + params)
  if err != nil {
    log.Println(err)
    return nil
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Println(err)
    return nil
  }

  broadcasts := &LiveBroadcasts{}
  err = json.Unmarshal(body, broadcasts)
  if err != nil {
    log.Println(err)
    return nil
  }

  return broadcasts.Items
}

func pollMessages(client *http.Client, liveChatId string) {
  pageToken := ""
  for {
    liveChatMessageListResponse := getMessages(client, liveChatId, pageToken)
    for i, message := range liveChatMessageListResponse.Items {
      switch message.Snippet.Type {
      case TEXT_MESSAGE_EVENT:
        if message.Snippet.MessageText != "" {
          fmt.Printf("%v: %v\n", message.AuthorDetails.DisplayName, html.UnescapeString(message.Snippet.MessageText))
        } else {
          fmt.Printf("%v: %v\n", message.AuthorDetails.DisplayName, html.UnescapeString(message.Snippet.FallbackText))
        }
      case FAN_FUNDING_EVENT:
        fmt.Println(html.UnescapeString(message.Snippet.FallbackText))
        addFanFundingMessage(&liveChatMessageListResponse.Items[i])
      }
    }
    pageToken = liveChatMessageListResponse.NextPageToken
    time.Sleep(5 + time.Second)
  }
}

func getMessages(client *http.Client, liveChatId string, pageToken string) *LiveChatMessageListResponse {
  pageTokenString := ""
  if pageToken != "" {
    pageTokenString = "&pageToken=" + pageToken
  }
  resp, err := client.Get("https://www.googleapis.com/youtube/v3/liveChat/messages?maxResults=50&part=snippet,authorDetails&liveChatId=" + liveChatId + pageTokenString)
  if err != nil {
    log.Println(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Println(err)
  }

  liveChatMessageListResponse := &LiveChatMessageListResponse{}
  err = json.Unmarshal(body, liveChatMessageListResponse)
  if err != nil {
    log.Println(err)
    return nil
  }

  return liveChatMessageListResponse
}

func generate(conf *oauth2.Config) {
  // Redirect user to Google's consent page to ask for permission
  // for the scopes specified above.
  url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
  fmt.Printf("Visit the following URL to generate an auth code, then rerun with -auth=<code> (It has also been copied to your clipboard):\n%v\n", url)
  clipboard.WriteAll(url)
}

func addFanFundingMessage(message *LiveChatMessage) {
  fanFunding.Lock()
  defer fanFunding.Unlock()

  if fanFunding.Messages[message.Id] == nil {
    fanFunding.Messages[message.Id] = message
    writeMessagesToFile([]*LiveChatMessage{message}, "latest.txt")
  }

  largest := message
  largestValue, err := strconv.ParseFloat(message.Snippet.FanFundingEventDetails.AmountMicros, 64)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, check := range fanFunding.Messages {
    checkValue, err := strconv.ParseFloat(check.Snippet.FanFundingEventDetails.AmountMicros, 64)
    if err != nil {
      fmt.Println(err)
      continue
    }
    if checkValue > largestValue {
      largest = check
      largestValue = checkValue
    }
  }

  writeMessagesToFile([]*LiveChatMessage{largest}, "largest.txt")
}

func writeMessagesToFile(messages []*LiveChatMessage, filename string) {
  output := ""
  for _, message := range messages {
    output += html.UnescapeString(message.Snippet.FallbackText) + "\n"
  }
  err := ioutil.WriteFile(filename, []byte(output), 0777)
  if err != nil {
    log.Println(err)
  }
}

func main() {
  b, err := ioutil.ReadFile(config)
  if err != nil {
    log.Fatal(err)
  }

  conf, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/youtube")

  if url {
    generate(conf)
    return
  }

  var tok *oauth2.Token

  if auth != "" {
    // Let's exchange our code
    tok, err = conf.Exchange(oauth2.NoContext, auth)
    if err != nil {
      log.Fatal(err)
    }

    b, err := json.Marshal(tok)
    if err != nil {
      log.Println(err)
    } else {
      err := ioutil.WriteFile(token, b, 0777)
      if err != nil {
        log.Println(err)
      }
    }
  } else {
    b, err = ioutil.ReadFile(token)
    if err == nil {
      // A token was found, unmarshall it and use it.
      tok = &oauth2.Token{}
      err := json.Unmarshal(b, tok)
      if err != nil {
        log.Fatal(err)
      }
    } else {
      // There was an error with the token, maybe it doesn't exist.
      // If we haven't been given an auth code, we must generate one.
      generate(conf)
      return
    }
  }

  client := conf.Client(oauth2.NoContext, tok)

  broadcasts := make([]LiveBroadcast, 0, 10)

  // Merge the two sources of broadcasts.
  broadcasts = append(broadcasts, getBroadcasts(client, "default=true")...)
  broadcasts = append(broadcasts, getBroadcasts(client, "mine=true")...)

  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt, os.Kill)

  for _, broadcast := range broadcasts {
    // If the broadcast has ended, it can't have a valid chat.
    if broadcast.Status.LifeCycleStatus == "complete" {
      continue
    }

    go pollMessages(client, broadcast.Snippet.LiveChatId)
  }

  <-c
}
