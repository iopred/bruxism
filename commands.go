package bruxism

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"sort"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/iopred/discordgo"
)

// HelpHelp returns help for the help plugin.
func HelpHelp(bot *Bot, service Service) (string, string) {
	ticks := ""
	if service.Name() == DiscordServiceName {
		ticks = "`"
	}

	commands := []string{}

	for _, plugin := range bot.Services[service.Name()].Plugins {
		t := plugin.Help(bot, service, true)

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
			h = plugin.Help(bot, service, false)
		} else if len(parts) == 1 && strings.ToLower(parts[0]) == strings.ToLower(plugin.Name()) {
			h = plugin.Help(bot, service, true)
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

// InviteHelp will return the help text for the invite command.
func InviteHelp(bot *Bot, service Service) (string, string) {
	switch service.Name() {
	case DiscordServiceName:
		return "<discordinvite>", "Joins the provided Discord server."
	case YouTubeServiceName:
		return "<livechatid>", "Joins the provided YouTube chat by id (this may be hard to find)."
	}
	return "<channel>", "Joins the provided channel."
}

// InviteCommand is a command for accepting an invite to a channel.
func InviteCommand(bot *Bot, service Service, message Message, command string, parts []string) {
	if len(parts) == 1 {
		join := parts[0]
		if service.Name() == DiscordServiceName {
			join = discordInviteID(join)
		}
		if err := service.Join(join); err != nil {
			if service.Name() == DiscordServiceName && err == ErrAlreadyJoined {
				service.PrivateMessage(message.UserID(), "I have already joined that server.")
				return
			}
			fmt.Printf("Error joining %v %v", service.Name(), err)
		} else if service.Name() == DiscordServiceName {
			service.PrivateMessage(message.UserID(), "I have joined that server.")
		}
	}
}

var startTime time.Time

func init() {
	startTime = time.Now()
}

func getDurationString(duration time.Duration) string {
	return fmt.Sprintf(
		"%0.2d:%02d:%02d",
		int(duration.Hours()),
		int(duration.Minutes())%60,
		int(duration.Seconds())%60,
	)
}

// StatsCommand returns bot statistics.
func StatsCommand(bot *Bot, service Service, message Message, command string, parts []string) {
	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)

	w := &tabwriter.Writer{}
	buf := &bytes.Buffer{}

	w.Init(buf, 0, 4, 0, ' ', 0)
	if service.Name() == DiscordServiceName {
		fmt.Fprintf(w, "```\n")
	}
	fmt.Fprintf(w, "Septapus: \t%s\n", VersionString)
	if service.Name() == DiscordServiceName {
		fmt.Fprintf(w, "Discordgo: \t%s\n", discordgo.VERSION)
	}
	fmt.Fprintf(w, "Go: \t%s\n", runtime.Version())
	fmt.Fprintf(w, "Uptime: \t%s\n", getDurationString(time.Now().Sub(startTime)))
	fmt.Fprintf(w, "Memory used: \t%s / %s\n", humanize.Bytes(stats.Alloc), humanize.Bytes(stats.TotalAlloc))
	fmt.Fprintf(w, "Concurrent tasks: \t%d\n", runtime.NumGoroutine())
	if service.Name() == DiscordServiceName {
		fmt.Fprintf(w, "Connected servers: \t%d\n", service.ChannelCount())
		fmt.Fprintf(w, "\n```")
	} else {
		fmt.Fprintf(w, "Connected channels: \t%d\n", service.ChannelCount())
	}
	w.Flush()

	out := buf.String() + "\nBuilt with love by iopred."

	if service.SupportsMultiline() {
		if err := service.SendMessage(message.Channel(), out); err != nil {
			log.Println(err)
		}
	} else {
		lines := strings.Split(out, "\n")
		for _, line := range lines {
			if err := service.SendMessage(message.Channel(), line); err != nil {
				log.Println(err)
			}
		}
	}
}
