package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"

  "github.com/atotto/clipboard"

  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
)

var url bool
var auth string

var config string
var token string

func init() {
  flag.BoolVar(&url, "url", false, "Generates a URL that provides an auth code.")
  flag.StringVar(&auth, "auth", "", "Exchanges the provided auth code for an oauth2 token.")
  flag.StringVar(&config, "config", "oauth2config.json", "The filename that contains the oauth2 config.")
  flag.StringVar(&token, "token", "token.json", "The filename to store the oauth2 token.")
  flag.Parse()
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
  Kind     string `json:"kind"`
  Etag     string `json:"etag"`
  PageInfo struct {
    TotalResults   int `json:"totalResults"`
    ResultsPerPage int `json:"resultsPerPage"`
  } `json:"pageInfo"`
  Items []LiveBroadcast `json:"items"`
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

  fmt.Printf(string(body))

  broadcasts := &LiveBroadcasts{}
  err = json.Unmarshal(body, broadcasts)
  if err != nil {
    log.Println(err)
    return nil
  }

  return broadcasts.Items
}

func getMessages(client *http.Client, liveChatId string) {
  resp, err := client.Get("https://www.googleapis.com/youtube/v3/liveChat/messages?part=snippet&liveChatId=" + liveChatId)
  if err != nil {
    log.Println(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Println(err)
  }

  fmt.Println(string(body))
}

func generate(conf *oauth2.Config) {
  // Redirect user to Google's consent page to ask for permission
  // for the scopes specified above.
  url := conf.AuthCodeURL("state")
  fmt.Printf("Visit the following URL to generate an auth code, then rerun with -auth=<code> (It has also been copied to your clipboard):\n%v\n", url)
  clipboard.WriteAll(url)
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

  for _, broadcast := range broadcasts {
    // If the broadcast has ended, it can't have a valid chat.
    if broadcast.Status.LifeCycleStatus == "complete" {
      continue
    }

    getMessages(client, broadcast.Snippet.LiveChatId)
  }

}
