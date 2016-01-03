package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/iopred/bruxism/bot"
)

var youtubeURL bool
var youtubeAuth string
var youtubeConfigFilename string
var youtubeTokenFilename string
var youtubeLiveChatIDs string
var discordEmail string
var discordPassword string
var imgurID string

func init() {
	flag.BoolVar(&youtubeURL, "youtubeuRL", false, "Generates a URL that provides an auth code.")
	flag.StringVar(&youtubeAuth, "youtubeauth", "", "Exchanges the provided auth code for an oauth2 token.")
	flag.StringVar(&youtubeConfigFilename, "youtubeconfig", "youtubeoauth2config.json", "The filename that contains the oauth2 config.")
	flag.StringVar(&youtubeTokenFilename, "youtubetoken", "youtubeoauth2token.json", "The filename to store the oauth2 token.")
	flag.StringVar(&youtubeLiveChatIDs, "youtubelivechatids", "", "Additional chat id's to poll.")
	flag.StringVar(&discordEmail, "discordemail", "", "Discord account email.")
	flag.StringVar(&discordPassword, "discordpassword", "", "Discord account password.")
	flag.StringVar(&imgurID, "imgurid", "", "Imgur client id.")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
}

func main() {
	b := bot.NewBot()
	b.ImgurID = imgurID
	youtube := bot.NewYouTube(youtubeURL, youtubeAuth, youtubeConfigFilename, youtubeTokenFilename, youtubeLiveChatIDs)
	discord := bot.NewDiscord(discordEmail, discordPassword)
	synirc := bot.NewIRC("irc.synirc.net", "Septapus", []string{"#logcabin"})

	b.RegisterService(youtube)
	b.RegisterService(discord)
	b.RegisterService(synirc)
	b.Open()

	cp := bot.NewCommandPlugin()
	cp.AddCommand("help", bot.HelpCommand, nil)
	cp.AddCommand("command", bot.HelpCommand, nil)
	cp.AddCommand("invite", bot.InviteCommand, bot.InviteHelp)
	cp.AddCommand("join", bot.InviteCommand, nil)
	cp.AddCommand("stats", bot.StatsCommand, bot.NewCommandHelp("", "Lists bot statistics."))

	b.RegisterPlugin(youtube, cp)
	b.RegisterPlugin(youtube, bot.NewSlowModePlugin())
	b.RegisterPlugin(youtube, bot.NewTopStreamersPlugin(youtube))
	b.RegisterPlugin(youtube, bot.NewStreamerPlugin(youtube))
	b.RegisterPlugin(youtube, bot.NewComicPlugin())

	b.RegisterPlugin(discord, cp)
	b.RegisterPlugin(discord, bot.NewTopStreamersPlugin(youtube))
	b.RegisterPlugin(discord, bot.NewStreamerPlugin(youtube))
	b.RegisterPlugin(discord, bot.NewPlayingPlugin())
	b.RegisterPlugin(discord, bot.NewComicPlugin())
	b.RegisterPlugin(discord, bot.NewDirectMessageInvitePlugin())

	b.RegisterPlugin(synirc, cp)
	b.RegisterPlugin(synirc, bot.NewTopStreamersPlugin(youtube))
	b.RegisterPlugin(synirc, bot.NewStreamerPlugin(youtube))
	b.RegisterPlugin(synirc, bot.NewComicPlugin())

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
