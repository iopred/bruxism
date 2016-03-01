package main

import (
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/iopred/bruxism"
	"github.com/iopred/bruxism/comicplugin"
	"github.com/iopred/bruxism/directmessageinviteplugin"
	"github.com/iopred/bruxism/emojiplugin"
	"github.com/iopred/bruxism/inviteplugin"
	"github.com/iopred/bruxism/liveplugin"
	"github.com/iopred/bruxism/mtgplugin"
	"github.com/iopred/bruxism/numbertriviaplugin"
	"github.com/iopred/bruxism/playingplugin"
	"github.com/iopred/bruxism/reminderplugin"
	"github.com/iopred/bruxism/slowmodeplugin"
	"github.com/iopred/bruxism/statsplugin"
	"github.com/iopred/bruxism/streamerplugin"
	"github.com/iopred/bruxism/topstreamersplugin"
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
	cp.AddCommand("invite", inviteplugin.InviteCommand, inviteplugin.InviteHelp)
	cp.AddCommand("join", inviteplugin.InviteCommand, nil)
	cp.AddCommand("stats", statsplugin.StatsCommand, statsplugin.StatsHelp)
	cp.AddCommand("info", statsplugin.StatsCommand, nil)
	cp.AddCommand("stat", statsplugin.StatsCommand, nil)
	cp.AddCommand("numbertrivia", numbertriviaplugin.NumberTriviaCommand, numbertriviaplugin.NumberTriviaHelp)
	cp.AddCommand("mtg", mtgplugin.MTGCommand, mtgplugin.MTGHelp)

	bot.RegisterPlugin(youtube, cp)
	bot.RegisterPlugin(youtube, slowmodeplugin.NewSlowModePlugin())
	bot.RegisterPlugin(youtube, topstreamersplugin.NewTopStreamersPlugin(youtube))
	bot.RegisterPlugin(youtube, streamerplugin.NewStreamerPlugin(youtube))
	bot.RegisterPlugin(youtube, comicplugin.New())
	bot.RegisterPlugin(youtube, reminderplugin.NewReminderPlugin())

	bot.RegisterPlugin(discord, cp)
	bot.RegisterPlugin(discord, topstreamersplugin.NewTopStreamersPlugin(youtube))
	bot.RegisterPlugin(discord, streamerplugin.NewStreamerPlugin(youtube))
	bot.RegisterPlugin(discord, playingplugin.NewPlayingPlugin())
	bot.RegisterPlugin(discord, comicplugin.New())
	bot.RegisterPlugin(discord, directmessageinviteplugin.NewDirectMessageInvitePlugin())
	bot.RegisterPlugin(discord, reminderplugin.NewReminderPlugin())
	bot.RegisterPlugin(discord, emojiplugin.NewEmojiPlugin())
	bot.RegisterPlugin(discord, liveplugin.NewLivePlugin(youtube))

	bot.RegisterPlugin(synirc, cp)
	bot.RegisterPlugin(synirc, topstreamersplugin.NewTopStreamersPlugin(youtube))
	bot.RegisterPlugin(synirc, streamerplugin.NewStreamerPlugin(youtube))
	bot.RegisterPlugin(synirc, comicplugin.New())
	bot.RegisterPlugin(synirc, reminderplugin.NewReminderPlugin())

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
