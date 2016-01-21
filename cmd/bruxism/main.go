package main

import (
	"flag"
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
var youtubeLiveVideoIDs string
var youtubeLiveChatIDs string
var discordToken string
var discordEmail string
var discordPassword string
var imgurID string
var imgurAlbum string
var mashableKey string

func init() {
	flag.BoolVar(&youtubeURL, "youtubeurl", false, "Generates a URL that provides an auth code.")
	flag.StringVar(&youtubeAuth, "youtubeauth", "", "Exchanges the provided auth code for an oauth2 token.")
	flag.StringVar(&youtubeConfigFilename, "youtubeconfig", "youtubeoauth2config.json", "The filename that contains the oauth2 config.")
	flag.StringVar(&youtubeTokenFilename, "youtubetoken", "youtubeoauth2token.json", "The filename to store the oauth2 token.")
	flag.StringVar(&youtubeLiveVideoIDs, "youtubelivevideoids", "", "Additional video id's to poll.")
	flag.StringVar(&youtubeLiveChatIDs, "youtubelivechatids", "", "Additional chat id's to poll.")
	flag.StringVar(&discordToken, "discordtoken", "", "Discord token.")
	flag.StringVar(&discordEmail, "discordemail", "", "Discord account email.")
	flag.StringVar(&discordPassword, "discordpassword", "", "Discord account password.")
	flag.StringVar(&imgurID, "imgurid", "", "Imgur client id.")
	flag.StringVar(&imgurAlbum, "imguralbum", "", "Imgur album id.")
	flag.StringVar(&mashableKey, "mashablekey", "", "Mashable key.")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
}

func main() {
	bot := bruxism.NewBot()
	bot.ImgurID = imgurID
	bot.ImgurAlbum = imgurAlbum
	bot.MashableKey = mashableKey

	youtube := bruxism.NewYouTube(youtubeURL, youtubeAuth, youtubeConfigFilename, youtubeTokenFilename, youtubeLiveVideoIDs, youtubeLiveChatIDs)

	var discord *bruxism.Discord
	if discordToken != "" {
		discord = bruxism.NewDiscord(discordToken)
	} else {
		discord = bruxism.NewDiscord(discordEmail, discordPassword)
	}
	synirc := bruxism.NewIRC("irc.synirc.net", "Septapus", []string{"#septapus"})

	bot.RegisterService(youtube)
	bot.RegisterService(discord)
	bot.RegisterService(synirc)
	bot.Open()

	cp := bruxism.NewCommandPlugin()
	cp.AddCommand("help", bruxism.HelpCommand, bruxism.HelpHelp)
	cp.AddCommand("command", bruxism.HelpCommand, nil)
	cp.AddCommand("commands", bruxism.HelpCommand, nil)
	cp.AddCommand("invite", bruxism.InviteCommand, bruxism.InviteHelp)
	cp.AddCommand("join", bruxism.InviteCommand, nil)
	cp.AddCommand("stats", bruxism.StatsCommand, bruxism.NewCommandHelp("", "Lists bot statistics."))
	cp.AddCommand("info", bruxism.StatsCommand, nil)
	cp.AddCommand("stat", bruxism.StatsCommand, nil)
	cp.AddCommand("numbertrivia", bruxism.NumberTriviaCommand, bruxism.NewCommandHelp("[<number>]", "Returns trivia for a random number or a specified number if provided."))
	cp.AddCommand("mtg", bruxism.MTGCommand, bruxism.NewCommandHelp("<cardname>", "Returns information about a Magic: The Gathering card."))

	bot.RegisterPlugin(youtube, cp)
	bot.RegisterPlugin(youtube, bruxism.NewSlowModePlugin())
	bot.RegisterPlugin(youtube, bruxism.NewTopStreamersPlugin(youtube))
	bot.RegisterPlugin(youtube, bruxism.NewStreamerPlugin(youtube))
	bot.RegisterPlugin(youtube, bruxism.NewComicPlugin())
	bot.RegisterPlugin(youtube, bruxism.NewReminderPlugin())

	bot.RegisterPlugin(discord, cp)
	bot.RegisterPlugin(discord, bruxism.NewTopStreamersPlugin(youtube))
	bot.RegisterPlugin(discord, bruxism.NewStreamerPlugin(youtube))
	bot.RegisterPlugin(discord, bruxism.NewPlayingPlugin())
	bot.RegisterPlugin(discord, bruxism.NewComicPlugin())
	bot.RegisterPlugin(discord, bruxism.NewDirectMessageInvitePlugin())
	bot.RegisterPlugin(discord, bruxism.NewReminderPlugin())

	bot.RegisterPlugin(synirc, cp)
	bot.RegisterPlugin(synirc, bruxism.NewTopStreamersPlugin(youtube))
	bot.RegisterPlugin(synirc, bruxism.NewStreamerPlugin(youtube))
	bot.RegisterPlugin(synirc, bruxism.NewComicPlugin())
	bot.RegisterPlugin(synirc, bruxism.NewReminderPlugin())

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
