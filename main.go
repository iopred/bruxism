package main

import (
  "flag"
  "fmt"
  "os"
  "os/signal"
  "time"

  "github.com/iopred/bruxism/bot"
)

var youtubeUrl bool
var youtubeAuth string
var youtubeConfigFilename string
var youtubeTokenFilename string
var youtubeLiveChatIds string
var discordEmail string
var discordPassword string

func init() {
  flag.BoolVar(&youtubeUrl, "youtubeurl", false, "Generates a URL that provides an auth code.")
  flag.StringVar(&youtubeAuth, "youtubeauth", "", "Exchanges the provided auth code for an oauth2 token.")
  flag.StringVar(&youtubeConfigFilename, "youtubeconfig", "youtubeoauth2config.json", "The filename that contains the oauth2 config.")
  flag.StringVar(&youtubeTokenFilename, "youtubetoken", "youtubeoauth2token.json", "The filename to store the oauth2 token.")
  flag.StringVar(&youtubeLiveChatIds, "youtubelivechatids", "", "Additional chat id's to poll.")
  flag.StringVar(&discordEmail, "discordemail", "", "Discord account email.")
  flag.StringVar(&discordPassword, "discordpassword", "", "Discord account password.")
  flag.Parse()
}

func main() {
  b := bot.NewBot()
  youtube := bot.NewYouTube(youtubeUrl, youtubeAuth, youtubeConfigFilename, youtubeTokenFilename, youtubeLiveChatIds)
  discord := bot.NewDiscord(discordEmail, discordPassword)

  b.RegisterService(youtube)
  b.RegisterService(discord)
  b.Open()

  cp := bot.NewCommandPlugin()
  cp.AddCommand("help", bot.HelpCommand, nil)
  cp.AddCommand("command", bot.HelpCommand, nil)
  cp.AddCommand("commands", bot.HelpCommand, nil)

  b.RegisterPlugin(youtube, cp)
  b.RegisterPlugin(youtube, bot.NewSlowModePlugin())
  b.RegisterPlugin(youtube, bot.NewTopStreamersPlugin(youtube))
  b.RegisterPlugin(youtube, bot.NewStreamerPlugin(youtube))
  b.RegisterPlugin(youtube, bot.NewInvitePlugin())

  b.RegisterPlugin(discord, cp)
  b.RegisterPlugin(discord, bot.NewTopStreamersPlugin(youtube))
  b.RegisterPlugin(discord, bot.NewStreamerPlugin(youtube))
  b.RegisterPlugin(discord, bot.NewPlayingPlugin())
  b.RegisterPlugin(discord, bot.NewInvitePlugin())

  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Recovered in f", r)
    }
    b.Save()
  }()

  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt, os.Kill)

  t := time.Tick(1 * time.Minute)

  for {
    select {
    case <-c:
      return
    case <-t:
      b.Save()
    }
  }
}
