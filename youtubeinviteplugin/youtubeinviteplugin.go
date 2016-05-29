package youtubeinviteplugin

import "github.com/iopred/bruxism"

// messageFunc is a command for accepting an YouTubeInvite to a channel.
func messageFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	if service.IsMe(message) || !bruxism.MatchesCommand(service, "youtubeinvite", message) {
		return
	}

	_, parts := bruxism.ParseCommand(service, message)

	if len(parts) == 1 {
		join := parts[0]

		ytservice := bot.Services[bruxism.YouTubeServiceName]
		if ytservice == nil {
			return
		}

		if err := ytservice.Join(join); err != nil {
			service.SendMessage(message.Channel(), err.Error())
		}
	}
}

func New() bruxism.Plugin {
	p := bruxism.NewSimplePlugin("YouTubeInvite")
	p.MessageFunc = messageFunc
	return p
}
