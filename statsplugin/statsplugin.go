package statsplugin

import (
	"bytes"
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/iopred/bruxism"
	"github.com/iopred/discordgo"
)

var statsStartTime = time.Now()

func getDurationString(duration time.Duration) string {
	return fmt.Sprintf(
		"%0.2d:%02d:%02d",
		int(duration.Hours()),
		int(duration.Minutes())%60,
		int(duration.Seconds())%60,
	)
}

// StatsCommand returns bot statistics.
func StatsCommand(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, command string, parts []string) {
	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)

	w := &tabwriter.Writer{}
	buf := &bytes.Buffer{}

	w.Init(buf, 0, 4, 0, ' ', 0)
	if service.Name() == bruxism.DiscordServiceName {
		fmt.Fprintf(w, "```\n")
	}
	fmt.Fprintf(w, "Bruxism: \t%s\n", bruxism.VersionString)
	if service.Name() == bruxism.DiscordServiceName {
		fmt.Fprintf(w, "Discordgo: \t%s\n", discordgo.VERSION)
	}
	fmt.Fprintf(w, "Go: \t%s\n", runtime.Version())
	fmt.Fprintf(w, "Uptime: \t%s\n", getDurationString(time.Now().Sub(statsStartTime)))
	fmt.Fprintf(w, "Memory used: \t%s / %s (%s garbage collected)\n", humanize.Bytes(stats.Alloc), humanize.Bytes(stats.Sys), humanize.Bytes(stats.TotalAlloc))
	fmt.Fprintf(w, "Concurrent tasks: \t%d\n", runtime.NumGoroutine())
	if service.Name() == bruxism.DiscordServiceName {
		discord := service.(*bruxism.Discord)
		fmt.Fprintf(w, "Connected servers: \t%d\n", service.ChannelCount())
		if len(discord.Sessions) > 1 {
			shards := 0
			for _, s := range discord.Sessions {
				if s.DataReady {
					shards++
				}
			}
			if shards == len(discord.Sessions) {
				fmt.Fprintf(w, "Shards: \t%d\n", shards)
			} else {
				fmt.Fprintf(w, "Shards: \t%d (%d connected)\n", len(discord.Sessions), shards)
			}
			guild, err := discord.Channel(message.Channel())
			if err == nil {
				id, err := strconv.Atoi(guild.ID)
				if err == nil {
					fmt.Fprintf(w, "Current shard: \t%d\n", ((id>>22)%len(discord.Sessions) + 1))
				}
			}
		}
	} else {
		fmt.Fprintf(w, "Connected channels: \t%d\n", service.ChannelCount())
	}

	plugins := bot.Services[service.Name()].Plugins
	names := []string{}
	for _, plugin := range plugins {
		names = append(names, plugin.Name())
		sort.Strings(names)
	}

	for _, name := range names {
		stats := plugins[name].Stats(bot, service, message)
		for _, stat := range stats {
			fmt.Fprint(w, stat)
		}
	}

	if service.Name() == bruxism.DiscordServiceName {
		fmt.Fprintf(w, "\n```")
	}

	w.Flush()
	out := buf.String()

	end := ""
	if IsSeptapus {
		end += "Septapus community: <https://discord.gg/HWN9pwj>\nPatreon: <https://www.patreon.com/iopred>\nBuilt with love by iopred."
	}

	if service.SupportsMultiline() {
		if end != "" {
			out += "\n" + end
		}
		service.SendMessage(message.Channel(), out)
	} else {
		service.SendMessage(message.Channel(), strings.Join(strings.Split(out, "\n"), " "))
		service.SendMessage(message.Channel(), strings.Join(strings.Split(end, "\n"), " "))
	}
}

var IsSeptapus bool = false

// StatsHelp is the help for the stats command.
var StatsHelp = bruxism.NewCommandHelp("", "Lists bot statistics.")
