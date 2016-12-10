package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/iopred/bruxism"
	"github.com/iopred/bruxism/carbonitexplugin"
	"github.com/iopred/bruxism/chartplugin"
	"github.com/iopred/bruxism/comicplugin"
	"github.com/iopred/bruxism/discordavatarplugin"
	"github.com/iopred/bruxism/emojiplugin"
	"github.com/iopred/bruxism/inviteplugin"
	"github.com/iopred/bruxism/liveplugin"
	"github.com/iopred/bruxism/musicplugin"
	"github.com/iopred/bruxism/mysonplugin"
	"github.com/iopred/bruxism/numbertriviaplugin"
	"github.com/iopred/bruxism/playedplugin"
	"github.com/iopred/bruxism/playingplugin"
	"github.com/iopred/bruxism/reminderplugin"
	"github.com/iopred/bruxism/statsplugin"
	"github.com/iopred/bruxism/streamerplugin"
	"github.com/iopred/bruxism/topstreamersplugin"
	"github.com/iopred/bruxism/triviaplugin"
	"github.com/iopred/bruxism/wormholeplugin"
	"github.com/iopred/bruxism/youtubejoinplugin"
)

var youtubeURL bool
var youtubeAuth string
var youtubeConfigFilename string
var youtubeTokenFilename string
var youtubeChannelIDs string
var youtubeInvitePort int
var discordToken string
var discordEmail string
var discordPassword string
var discordApplicationClientID string
var discordOwnerUserID string
var discordShards int
var ircServer string
var ircUsername string
var ircPassword string
var ircChannels string
var slackToken string
var slackOwnerUserID string
var imgurID string
var imgurAlbum string
var mashableKey string
var carbonitexKey string

func init() {
	flag.BoolVar(&youtubeURL, "youtubeurl", false, "Generates a URL that provides an auth code.")
	flag.StringVar(&youtubeAuth, "youtubeauth", "", "Exchanges the provided auth code for an oauth2 token.")
	flag.StringVar(&youtubeConfigFilename, "youtubeconfig", "youtubeoauth2config.json", "The filename that contains the oauth2 config.")
	flag.StringVar(&youtubeTokenFilename, "youtubetoken", "youtubeoauth2token.json", "The filename to store the oauth2 token.")
	flag.StringVar(&youtubeChannelIDs, "youtubechannelids", "", "Comma separated list of channel ids to poll.")
	flag.IntVar(&youtubeInvitePort, "youtubeinviteport", 7777, "The port to listen for invites.")
	flag.StringVar(&discordToken, "discordtoken", "", "Discord token.")
	flag.StringVar(&discordEmail, "discordemail", "", "Discord account email.")
	flag.StringVar(&discordPassword, "discordpassword", "", "Discord account password.")
	flag.StringVar(&discordOwnerUserID, "discordowneruserid", "", "Discord owner user id.")
	flag.StringVar(&discordApplicationClientID, "discordapplicationclientid", "", "Discord application client id.")
	flag.IntVar(&discordShards, "discordshards", 1, "Number of discord shards.")
	flag.StringVar(&ircServer, "ircserver", "", "IRC server.")
	flag.StringVar(&ircUsername, "ircusername", "", "IRC user name.")
	flag.StringVar(&ircPassword, "ircpassword", "", "IRC password.")
	flag.StringVar(&ircChannels, "ircchannels", "", "Comma separated list of IRC channels.")
	flag.StringVar(&slackToken, "slacktoken", "", "Slack token.")
	flag.StringVar(&slackOwnerUserID, "slackowneruserid", "", "Slack owner user id.")
	flag.StringVar(&imgurID, "imgurid", "", "Imgur client id.")
	flag.StringVar(&imgurAlbum, "imguralbum", "", "Imgur album id.")
	flag.StringVar(&mashableKey, "mashablekey", "", "Mashable key.")
	flag.StringVar(&carbonitexKey, "carbonitexkey", "", "Carbonitex key for discord server count tracking.")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
}

func main() {
	q := make(chan bool)

	// Set our variables.
	bot := bruxism.NewBot()
	bot.ImgurID = imgurID
	bot.ImgurAlbum = imgurAlbum
	bot.MashableKey = mashableKey

	// Generally CommandPlugins don't hold state, so we share one instance of the command plugin for all services.
	cp := bruxism.NewCommandPlugin()
	cp.AddCommand("invite", inviteplugin.InviteCommand, inviteplugin.InviteHelp)
	cp.AddCommand("join", inviteplugin.InviteCommand, nil)
	cp.AddCommand("stats", statsplugin.StatsCommand, statsplugin.StatsHelp)
	cp.AddCommand("info", statsplugin.StatsCommand, nil)
	cp.AddCommand("stat", statsplugin.StatsCommand, nil)
	if bot.MashableKey != "" {
		cp.AddCommand("numbertrivia", numbertriviaplugin.NumberTriviaCommand, numbertriviaplugin.NumberTriviaHelp)
	}
	cp.AddCommand("quit", func(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, args string, parts []string) {
		if service.IsBotOwner(message) {
			q <- true
		}
	}, nil)

	youtube := bruxism.NewYouTube(youtubeURL, youtubeAuth, youtubeConfigFilename, youtubeTokenFilename)
	err := youtube.Init()
	if err != nil {
		log.Fatal(err)
	}

	ytLiveChannel := bruxism.NewYTLiveChannel(youtube.Service)
	ytip := youtubejoinplugin.New(ytLiveChannel)

	bot.RegisterService(youtube)

	bot.RegisterPlugin(youtube, cp)

	bot.RegisterPlugin(youtube, chartplugin.New())
	bot.RegisterPlugin(youtube, comicplugin.New())
	bot.RegisterPlugin(youtube, liveplugin.New(ytLiveChannel))
	bot.RegisterPlugin(youtube, reminderplugin.New())
	bot.RegisterPlugin(youtube, streamerplugin.New(youtube))
	bot.RegisterPlugin(youtube, topstreamersplugin.New(youtube))
	bot.RegisterPlugin(youtube, triviaplugin.New())
	bot.RegisterPlugin(youtube, wormholeplugin.New())
	bot.RegisterPlugin(youtube, ytip)

	// Register the Discord service if we have an email or token.
	if (discordEmail != "" && discordPassword != "") || discordToken != "" {
		var discord *bruxism.Discord
		if discordToken != "" {
			discord = bruxism.NewDiscord(discordToken)
		} else {
			discord = bruxism.NewDiscord(discordEmail, discordPassword)
		}
		discord.ApplicationClientID = discordApplicationClientID
		discord.OwnerUserID = discordOwnerUserID
		discord.Shards = discordShards
		bot.RegisterService(discord)

		bot.RegisterPlugin(discord, cp)

		if carbonitexKey != "" {
			bot.RegisterPlugin(discord, carbonitexplugin.New(carbonitexKey))
		}

		bot.RegisterPlugin(discord, chartplugin.New())
		bot.RegisterPlugin(discord, comicplugin.New())
		bot.RegisterPlugin(discord, discordavatarplugin.New())
		bot.RegisterPlugin(discord, emojiplugin.New())
		bot.RegisterPlugin(discord, liveplugin.New(ytLiveChannel))
		bot.RegisterPlugin(discord, musicplugin.New(discord))
		bot.RegisterPlugin(discord, mysonplugin.New())
		bot.RegisterPlugin(discord, playedplugin.New())
		bot.RegisterPlugin(discord, playingplugin.New())
		bot.RegisterPlugin(discord, reminderplugin.New())
		bot.RegisterPlugin(discord, streamerplugin.New(youtube))
		bot.RegisterPlugin(discord, topstreamersplugin.New(youtube))
		bot.RegisterPlugin(discord, triviaplugin.New())
		bot.RegisterPlugin(discord, wormholeplugin.New())
	}

	// Register the IRC service if we have an IRC server and Username.
	if ircServer != "" && ircUsername != "" {
		irc := bruxism.NewIRC(ircServer, ircUsername, ircPassword, strings.Split(ircChannels, ","))
		bot.RegisterService(irc)

		bot.RegisterPlugin(irc, cp)

		bot.RegisterPlugin(irc, chartplugin.New())
		bot.RegisterPlugin(irc, comicplugin.New())
		bot.RegisterPlugin(irc, liveplugin.New(ytLiveChannel))
		bot.RegisterPlugin(irc, reminderplugin.New())
		bot.RegisterPlugin(irc, streamerplugin.New(youtube))
		bot.RegisterPlugin(irc, topstreamersplugin.New(youtube))
		bot.RegisterPlugin(irc, triviaplugin.New())
		bot.RegisterPlugin(irc, wormholeplugin.New())
	}

	if slackToken != "" {
		slack := bruxism.NewSlack(slackToken)
		slack.OwnerUserID = slackOwnerUserID
		bot.RegisterService(slack)

		bot.RegisterPlugin(slack, cp)

		bot.RegisterPlugin(slack, chartplugin.New())
		bot.RegisterPlugin(slack, comicplugin.New())
		bot.RegisterPlugin(slack, liveplugin.New(ytLiveChannel))
		bot.RegisterPlugin(slack, streamerplugin.New(youtube))
		bot.RegisterPlugin(slack, topstreamersplugin.New(youtube))
		bot.RegisterPlugin(slack, triviaplugin.New())
		bot.RegisterPlugin(slack, wormholeplugin.New())
	}

	// Start all our services.
	bot.Open()

	if youtubeChannelIDs != "" {
		channelIDs := strings.Split(youtubeChannelIDs, ",")

		for _, channelID := range channelIDs {
			ytip.Monitor(channelID)
		}
	}

	// Wait for a termination signal, while saving the bot state every minute. Save on close.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	t := time.Tick(1 * time.Minute)

out:
	for {
		select {
		case <-q:
			break out
		case <-c:
			break out
		case <-t:
			bot.Save()
		}
	}

	bot.Save()
}
