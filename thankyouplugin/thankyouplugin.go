package thankyouplugin

import (
	"strings"

	"github.com/iopred/bruxism"
)

// ThankyouCommand returns bot statistics.
func ThankyouCommand(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, command string, parts []string) {
	out := "Septapus currently has no patrons.\nBecome a patron at: <https://www.patreon.com/iopred>"

	if service.SupportsMultiline() {
		service.SendMessage(message.Channel(), out)
	} else {
		service.SendMessage(message.Channel(), strings.Join(strings.Split(out, "\n"), ", "))
	}
}

// StatsHelp is the help for the stats command.
var StatsHelp = bruxism.NewCommandHelp("", "Thank you to all my patrons.")
