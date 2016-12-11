package mysonplugin

import (
	"github.com/iopred/bruxism"
	"github.com/iopred/discordgo"
)

func messageFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	if service.IsMe(message) || !bruxism.MatchesCommand(service, "myson", message) || service.Name() != bruxism.DiscordServiceName {
		return
	}

	discord := service.(*bruxism.Discord)
	discord.Session.ChannelMessageSendEmbed(message.Channel(), &discordgo.MessageEmbed{
		Color:       discord.UserColor(service.UserID(), message.Channel()),
		Description: "Don't ever talk to me or my son ever again.",
		Author: &discordgo.MessageEmbedAuthor{
			Name:    discord.NicknameForID(service.UserID(), service.UserName(), message.Channel()),
			IconURL: discordgo.EndpointUserAvatar(service.UserID(), discord.Session.State.User.Avatar),
		},
	})
}

func helpFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, detailed bool) []string {
	return nil
}

func New() bruxism.Plugin {
	p := bruxism.NewSimplePlugin("MySon")
	p.MessageFunc = messageFunc
	p.HelpFunc = helpFunc
	return p
}
