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

func (p *InvitePlugin) Help() string {
	return p.help
}

func (p *InvitePlugin) Register(bot *Bot, service Service, data []byte) error {
	conversion := func(str string) string {
		return str
	}

	switch service.Name() {
	case DiscordServiceName:
		p.help = "!invite [discordinvite] - Joins the provided Discord server."
		conversion = func(str string) string {
			str = strings.Replace(str, "https://discord.gg/", "", -1)
			str = strings.Replace(str, "http://discord.gg/", "", -1)
			return str
		}
	case YouTubeServiceName:
		p.help = "!invite [livechatid] - Joins the provided YouTube chat by id (this may be hard to find)."
	default:
		p.help = "!invite [channel] - Joins the provided channel."
	}

	messageChannel := bot.NewMessageChannel(service)
	go func() {
		for {
			message := <-messageChannel

			messageMessage := message.Message()

			if strings.HasPrefix(messageMessage, "!invite") {
				splits := strings.Split(messageMessage, "!invite ")

				if len(splits) != 2 {
					continue
				}

				if err := service.Join(conversion(splits[1])); err != nil {
					fmt.Println("Error joining", service, err)
				}
			}
		}
	}()
	return nil
}

func (p *InvitePlugin) Save() []byte {
	return nil
}

func NewInvitePlugin() *InvitePlugin {
	return &InvitePlugin{}
}
