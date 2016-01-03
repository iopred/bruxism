package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/iopred/bruxism"
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
	b := bruxism.New()
	b.ImgurID = imgurID
	youtube := bruxism.NewYouTube(youtubeURL, youtubeAuth, youtubeConfigFilename, youtubeTokenFilename, youtubeLiveChatIDs)
	discord := bruxism.NewDiscord(discordEmail, discordPassword)
	synirc := bruxism.NewIRC("irc.synirc.net", "Septapus", []string{"#logcabin"})

	b.RegisterService(youtube)
	b.RegisterService(discord)
	b.RegisterService(synirc)
	b.Open()

	cp := bruxism.NewCommandPlugin()
	cp.AddCommand("help", bruxism.HelpCommand, nil)
	cp.AddCommand("command", bruxism.HelpCommand, nil)
	cp.AddCommand("invite", bruxism.InviteCommand, bruxism.InviteHelp)
	cp.AddCommand("join", bruxism.InviteCommand, nil)
	cp.AddCommand("stats", bruxism.StatsCommand, bruxism.NewCommandHelp("", "Lists bot statistics."))

	b.RegisterPlugin(youtube, cp)
	b.RegisterPlugin(youtube, bruxism.NewSlowModePlugin())
	b.RegisterPlugin(youtube, bruxism.NewTopStreamersPlugin(youtube))
	b.RegisterPlugin(youtube, bruxism.NewStreamerPlugin(youtube))
	b.RegisterPlugin(youtube, bruxism.NewComicPlugin())

	b.RegisterPlugin(discord, cp)
	b.RegisterPlugin(discord, bruxism.NewTopStreamersPlugin(youtube))
	b.RegisterPlugin(discord, bruxism.NewStreamerPlugin(youtube))
	b.RegisterPlugin(discord, bruxism.NewPlayingPlugin())
	b.RegisterPlugin(discord, bruxism.NewComicPlugin())
	b.RegisterPlugin(discord, bruxism.NewDirectMessageInvitePlugin())

	b.RegisterPlugin(synirc, cp)
	b.RegisterPlugin(synirc, bruxism.NewTopStreamersPlugin(youtube))
	b.RegisterPlugin(synirc, bruxism.NewStreamerPlugin(youtube))
	b.RegisterPlugin(synirc, bruxism.NewComicPlugin())

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
