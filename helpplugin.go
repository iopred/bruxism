package bruxism

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

// HelpHelp returns help for the help plugin.
func HelpHelp(bot *Bot, service Service, message Message) (string, string) {
	ticks := ""
	if service.Name() == DiscordServiceName {
		ticks = "`"
	}

	commands := []string{}

	for _, plugin := range bot.Services[service.Name()].Plugins {
		t := plugin.Help(bot, service, message, true)

		if t != nil && len(t) > 0 {
			commands = append(commands, strings.ToLower(plugin.Name()))
		}
	}

	sort.Strings(commands)

	return "[<topic>]", fmt.Sprintf("Returns generic help or help for a specific topic. Available topics: %s%s%s", ticks, strings.Join(commands, ", "), ticks)
}

// HelpCommand is a command for returning help text for all registered plugins on a service.
func HelpCommand(bot *Bot, service Service, message Message, command string, parts []string) {
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
		help = append([]string{fmt.Sprintf("All commands can be used in private messages without the `%s` prefix.", service.CommandPrefix())}, help...)
	}

	if len(parts) != 0 && len(help) == 0 {
		help = []string{fmt.Sprintf("Unknown topic: %s", parts[0])}
	}

	if service.SupportsMultiline() {
		if err := service.SendMessage(message.Channel(), strings.Join(help, "\n")); err != nil {
			log.Println(err)
		}
	} else {
		for _, h := range help {
			if err := service.SendMessage(message.Channel(), h); err != nil {
				log.Println(err)
			}
		}
	}
}
