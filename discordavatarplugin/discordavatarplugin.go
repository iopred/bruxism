package discordavatarplugin

import (
	"regexp"
	"strings"

	"github.com/iopred/bruxism"
	"github.com/iopred/discordgo"
)

var userIDRegex = regexp.MustCompile("<@!?([0-9]*)>")

func avatarLoadFunc(bot *bruxism.Bot, service bruxism.Service, data []byte) error {
	if service.Name() != bruxism.DiscordServiceName {
		panic("DiscordAvatar plugin only supports Discord.")
	}
	return nil
}

func avatarMessageFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	if service.IsMe(message) {
		return
	}

	if !bruxism.MatchesCommand(service, "avatar", message) {
		return
	}

	query := strings.Join(strings.Split(message.RawMessage(), " ")[1:], " ")

	id := message.UserID()
	match := userIDRegex.FindStringSubmatch(query)
	if match != nil {
		id = match[1]
	}

	discord := service.(*bruxism.Discord)

	u, err := discord.Session.User(id)
	if err != nil {
		return
	}

	service.SendMessage(message.Channel(), discordgo.EndpointUserAvatar(u.ID, u.Avatar))
}

func avatarHelpFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, detailed bool) []string {
	if detailed {
		return nil
	}
	return bruxism.CommandHelp(service, "avatar", "[@username]", "Returns a big version of your avatar, or a users avatar if provided.")
}

// New creates a new discordavatar plugin.
func New() bruxism.Plugin {
	p := bruxism.NewSimplePlugin("DiscordAvatar")
	p.LoadFunc = avatarLoadFunc
	p.MessageFunc = avatarMessageFunc
	p.HelpFunc = avatarHelpFunc
	return p
}
