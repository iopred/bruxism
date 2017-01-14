package bruxism

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
)

type helpPlugin struct {
	Private map[string]bool
}

// Name returns the name of the service.
func (p *helpPlugin) Name() string {
	return "Help"
}

// Help returns a list of help strings that are printed when the user requests them.
func (p *helpPlugin) Help(bot *Bot, service Service, message Message, detailed bool) []string {
	privs := service.SupportsPrivateMessages() && !service.IsPrivate(message) && service.IsModerator(message)
	if detailed && !privs {
		return nil
	}

	ticks := ""
	if service.Name() == DiscordServiceName {
		ticks = "`"
	}

	commands := []string{}

	for _, plugin := range bot.Services[service.Name()].Plugins {
		hasDetailed := false

		if plugin == p {
			hasDetailed = privs
		} else {
			t := plugin.Help(bot, service, message, true)
			hasDetailed = t != nil && len(t) > 0
		}

		if hasDetailed {
			commands = append(commands, strings.ToLower(plugin.Name()))
		}
	}

	sort.Strings(commands)

	help := []string{}

	if len(commands) > 0 {
		help = append(help, CommandHelp(service, "help", "[topic]", fmt.Sprintf("Returns help for a specific topic. Available topics: %s%s%s", ticks, strings.Join(commands, ", "), ticks))[0])
	}

	if detailed {
		help = append(help, []string{
			CommandHelp(service, "setprivatehelp", "", "Sets help text to be sent through private messages in this channel.")[0],
			CommandHelp(service, "setpublichelp", "", "Sets the default help behavior for this channel.")[0],
		}...)
	}

	return help
}

func (p *helpPlugin) Message(bot *Bot, service Service, message Message) {
	if !service.IsMe(message) {
		if MatchesCommand(service, "help", message) || MatchesCommand(service, "command", message) {
			if service.Name() == YouTubeServiceName {
				service.SendMessage(message.Channel(), "Visit septapus dot com for help.")
				return
			}

			_, parts := ParseCommand(service, message)

			help := []string{}

			for _, plugin := range bot.Services[service.Name()].Plugins {
				var h []string
				if len(parts) == 0 {
					h = plugin.Help(bot, service, message, false)
				} else if len(parts) == 1 && strings.ToLower(parts[0]) == strings.ToLower(plugin.Name()) {
					h = plugin.Help(bot, service, message, true)
				}
				if h != nil && len(h) > 0 {
					help = append(help, h...)
				}
			}

			if len(parts) == 0 {
				sort.Strings(help)
				if service.SupportsPrivateMessages() {
					help = append([]string{fmt.Sprintf("All commands can be used in private messages without the `%s` prefix.", service.CommandPrefix())}, help...)
				}
			}

			if len(parts) != 0 && len(help) == 0 {
				help = []string{fmt.Sprintf("Unknown topic: %s", parts[0])}
			}

			if p.Private[message.Channel()] {
				service.SendMessage(message.Channel(), "Help has been sent via private message.")
				if service.SupportsMultiline() {
					service.PrivateMessage(message.UserID(), strings.Join(help, "\n"))
				} else {
					for _, h := range help {
						if err := service.PrivateMessage(message.UserID(), h); err != nil {
							break
						}
					}
				}
			} else {
				if service.SupportsMultiline() {
					service.SendMessage(message.Channel(), strings.Join(help, "\n"))
				} else {
					for _, h := range help {
						if err := service.SendMessage(message.Channel(), h); err != nil {
							break
						}
					}
				}
			}
		} else if MatchesCommand(service, "setprivatehelp", message) && service.SupportsPrivateMessages() && !service.IsPrivate(message) {
			if !service.IsModerator(message) {
				return
			}

			p.Private[message.Channel()] = true

			service.PrivateMessage(message.UserID(), fmt.Sprintf("Help text in <#%s> will be sent through private messages.", message.Channel()))
		} else if MatchesCommand(service, "setpublichelp", message) && service.SupportsPrivateMessages() && !service.IsPrivate(message) {
			if !service.IsModerator(message) {
				return
			}

			p.Private[message.Channel()] = false

			service.PrivateMessage(message.UserID(), fmt.Sprintf("Help text in <#%s> will be sent publically.", message.Channel()))
		}
	}
}

// Load will load plugin state from a byte array.
func (p *helpPlugin) Load(bot *Bot, service Service, data []byte) error {
	if data != nil {
		if err := json.Unmarshal(data, p); err != nil {
			log.Println("Error loading data", err)
		}
	}
	return nil
}

// Save will save plugin state to a byte array.
func (p *helpPlugin) Save() ([]byte, error) {
	return json.Marshal(p)
}

// Stats will return the stats for a plugin.
func (p *helpPlugin) Stats(bot *Bot, service Service, message Message) []string {
	return nil
}

// NeHelpPlugin will create a new help plugin.
func NewHelpPlugin() Plugin {
	p := &helpPlugin{
		Private: make(map[string]bool),
	}
	return p
}
