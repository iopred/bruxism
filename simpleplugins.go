package bruxism

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func discordInviteID(id string) string {
	id = strings.Replace(id, "://discordapp.com/invite/", "://discord.gg/", -1)
	id = strings.Replace(id, "https://discord.gg/", "", -1)
	id = strings.Replace(id, "http://discord.gg/", "", -1)
	return id
}

func directMessageInviteMessageFunc(bot *Bot, service Service, message Message) {
	if service.Name() == DiscordServiceName && !service.IsMe(message) && service.IsPrivate(message) {
		messageMessage := message.Message()
		id := discordInviteID(messageMessage)
		if id != messageMessage && strings.HasPrefix(messageMessage, "http") {
			if err := service.Join(id); err != nil {
				if err == ErrAlreadyJoined {
					service.PrivateMessage(message.UserID(), "I have already joined that server.")
					return
				}
				fmt.Printf("Error joining %v %v", service.Name(), err)
				return
			}
			service.PrivateMessage(message.UserID(), "I have joined that server.")
		}
	}
}

// NewDirectMessageInvitePlugin creates a new direct message invite plugin.
func NewDirectMessageInvitePlugin() Plugin {
	p := NewSimplePlugin("DirectMessageInvite")
	p.MessageFunc = directMessageInviteMessageFunc
	return p
}

func emojiMessageFunc(bot *Bot, service Service, message Message) {
	if service.Name() == DiscordServiceName && !service.IsMe(message) {
		if MatchesCommand(service, "emoji", message) {
			_, parts := ParseCommand(service, message)
			if len(parts) == 1 {
				s := strings.TrimSpace(parts[0])
				for _, r := range s {
					if f, err := os.Open(fmt.Sprintf("emoji/twitter/%s.png", strconv.FormatInt(int64(r), 16))); err == nil {
						defer f.Close()
						service.SendFile(message.Channel(), "emjoi.png", f)
					}
					return
				}
			}

		}
	}
}

func emojiHelpFunc(bot *Bot, service Service, detailed bool) []string {
	if detailed {
		return nil
	}
	return CommandHelp(service, "emoji", "<emoji>", "Returns a big version of an emoji.")
}

// NewEmojiPlugin creates a new emoji plugin.
func NewEmojiPlugin() Plugin {
	p := NewSimplePlugin("Emoji")
	p.MessageFunc = emojiMessageFunc
	p.HelpFunc = emojiHelpFunc
	return p
}
