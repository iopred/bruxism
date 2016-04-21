package discordavatarplugin

import (
	"fmt"
	"strings"

	"github.com/iopred/bruxism"
	"github.com/iopred/discordgo"
)

func sendAvatar(bot *bruxism.Bot, service bruxism.Service, channel, user string) {
	discord := service.(*bruxism.Discord)

	u, err := discord.Session.User(user)
	if err != nil {
		return
	}

	service.SendMessage(channel, discordgo.USER_AVATAR(u.ID, u.Avatar))
}

func avatarMessageFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	if service.Name() == bruxism.DiscordServiceName && !service.IsMe(message) {
		if bruxism.MatchesCommand(service, "avatar", message) {
			_, parts := bruxism.ParseCommand(service, message)
			switch len(parts) {
			case 0:
				sendAvatar(bot, service, message.Channel(), message.UserID())
			case 1:
				t := strings.TrimSpace(strings.Split(strings.TrimSpace(message.RawMessage()), " ")[2])

				if strings.Index(t, "<") == 0 {
					fmt.Println(t[2 : len(t)-1])
					sendAvatar(bot, service, message.Channel(), t[2:len(t)-1])
				}
			}
		}
	}
}

func avatarHelpFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, detailed bool) []string {
	if detailed {
		return nil
	}
	return bruxism.CommandHelp(service, "avatar", "[@username]", "Returns a big version of your avatar, or a users avatar if provided.")
}

// New creates a new discordavatar plugin.
func New() bruxism.Plugin {
	p := bruxism.NewSimplePlugin("discordavatar")
	p.MessageFunc = avatarMessageFunc
	p.HelpFunc = avatarHelpFunc
	return p
}
