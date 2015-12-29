package bot

import (
	"fmt"
	"strings"
)

type InvitePlugin struct {
	help string
}

func (p *InvitePlugin) Name() string {
	return "Invite"
}

func (p *InvitePlugin) Help(bot *Bot, service Service) []string {
	help := ""
	switch service.Name() {
	case DiscordServiceName:
		help = "!invite [discordinvite] - Joins the provided Discord server."
	case YouTubeServiceName:
		help = "!invite [livechatid] - Joins the provided YouTube chat by id (this may be hard to find)."
	default:
		help = "!invite [channel] - Joins the provided channel."
	}
	return []string{help}
}

func (p *InvitePlugin) Load(bot *Bot, service Service, data []byte) error {
	return nil
}

func (p *InvitePlugin) Save() ([]byte, error) {
	return nil, nil
}

func (p *InvitePlugin) Message(bot *Bot, service Service, message Message) {
	if matchesCommand("invite", message) {
		_, parts := parseCommand(message)
		if len(parts) == 1 {
			join := parts[0]
			if service.Name() == DiscordServiceName {
				join = strings.Replace(join, "https://discord.gg/", "", -1)
				join = strings.Replace(join, "http://discord.gg/", "", -1)
			}
			if err := service.Join(join); err != nil {
				fmt.Printf("Error joining %v %v", service.Name(), err)
			}
		}
	}

}

func NewInvitePlugin() *InvitePlugin {
	return &InvitePlugin{}
}
