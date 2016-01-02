package bot

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

// HelpCommand is a command for returning help text for all registered plugins on a service.
func HelpCommand(bot *Bot, service Service, message Message) {
	help := []string{}

	for _, plugin := range bot.Services[service.Name()].Plugins {
		h := plugin.Help(bot, service)
		if h != nil && len(h) > 0 {
			help = append(help, h...)
		}
	}

	sort.Strings(help)

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
func InviteCommand(bot *Bot, service Service, message Message) {
	_, parts := parseCommand(service, message)
	if len(parts) == 1 {
		join := parts[0]
		if service.Name() == DiscordServiceName {
			join = discordInviteID(join)
		}
		if err := service.Join(join); err != nil {
			if service.Name() == DiscordServiceName && err == AlreadyJoinedError {
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

// StatsHelp will return the help text for the stats command.
func StatsHelp(bot *Bot, service Service) (string, string) {
	return "", "Lists bot statistics."
}

// StatsCommand returns bot statistics.
func StatsCommand(bot *Bot, service Service, message Message) {
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
	fmt.Fprintf(w, "Uptime: \t%s\n", getDurationString(time.Now().Sub(startTime)))
	fmt.Fprintf(w, "Memory used: \t%s\n", humanize.Bytes(stats.Alloc))
	fmt.Fprintf(w, "Concurrent tasks: \t%d", runtime.NumGoroutine())
	if service.Name() == DiscordServiceName {
		fmt.Fprintf(w, "\n```")
	}
	w.Flush()

	out := buf.String()

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
