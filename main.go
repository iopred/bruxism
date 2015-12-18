package main

import (
  "flag"
  "os"
  "os/signal"
  "time"
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
  bot := NewBot()
  youtube := NewYouTube(youtubeUrl, youtubeAuth, youtubeConfigFilename, youtubeTokenFilename, youtubeLiveChatIds)
  discord := NewDiscord(discordEmail, discordPassword)

  bot.RegisterService(youtube)
  bot.RegisterPlugin(youtube, NewHelpPlugin())
  bot.RegisterPlugin(youtube, NewSlowModePlugin())
  bot.RegisterPlugin(youtube, NewTopStreamersPlugin(youtube))

  bot.RegisterService(discord)
  bot.RegisterPlugin(discord, NewHelpPlugin())
  bot.RegisterPlugin(discord, NewTopStreamersPlugin(youtube))
  bot.Open()

  defer func() {
    recover()
    bot.Save()
  }()

  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt, os.Kill)

  t := time.Tick(1 * time.Minute)

  for {
    select {
    case <-c:
      return
    case <-t:
      bot.Save()
    }
  }
}
