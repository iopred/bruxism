package emojiplugin

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/iopred/bruxism"
)

var emoji map[string]bool = map[string]bool{}

func init() {
	f, err := ioutil.ReadDir("emoji/twitter")
	if err != nil {
		return
	}
	for _, file := range f {
		emoji[file.Name()] = true
	}
}

func emojiFile(s string) string {
	found := ""
	filename := ""
	for _, r := range s {
		if filename != "" {
			filename = fmt.Sprintf("%s-%x", filename, r)
		} else {
			filename = fmt.Sprintf("%x", r)
		}
		if emoji[fmt.Sprintf("%s.png", filename)] {
			found = filename
		}
	}
	return found
}

func emojiLoadFunc(bot *bruxism.Bot, service bruxism.Service, data []byte) error {
	if service.Name() != bruxism.DiscordServiceName {
		panic("Emoji Plugin only supports Discord.")
	}
	return nil
}

var discordRegex = regexp.MustCompile("<(a?):.*?:(.*?)>")

func emojiMessageFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	if service.Name() != bruxism.DiscordServiceName {
		return
	}

	if service.IsMe(message) {
		return
	}

	if !bruxism.MatchesCommand(service, "emoji", message) && !bruxism.MatchesCommand(service, "hugemoji", message) && !bruxism.MatchesCommand(service, "hugeemoji", message) {
		return
	}

	base := "emoji/twitter"
	if bruxism.MatchesCommand(service, "hugemoji", message) || bruxism.MatchesCommand(service, "hugeemoji", message) {
		base = "emoji/twitterhuge"
	}
	_, parts := bruxism.ParseCommand(service, message)
	if len(parts) == 1 {
		submatches := discordRegex.FindStringSubmatch(parts[0])
		if len(submatches) != 0 {
			fileType := "png"
			url := discordgo.EndpointEmoji(submatches[2])
			if submatches[1] == "a" {
				fileType = "gif"
				url = discordgo.EndpointEmojiAnimated(submatches[2])
			}
			h, err := http.Get(url)
			if err != nil {
				return
			}

			service.SendFile(message.Channel(), "emoji."+fileType, h.Body)
			h.Body.Close()

			return

		}

		s := strings.TrimSpace(parts[0])
		filename := emojiFile(s)
		if filename != "" {
			if f, err := os.Open(fmt.Sprintf("%s/%s.png", base, filename)); err == nil {
				defer f.Close()
				service.SendFile(message.Channel(), "emoji.png", f)

				return
			}
		}
	}
}

func emojiHelpFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, detailed bool) []string {
	help := bruxism.CommandHelp(service, "emoji", "<emoji>", "Returns a big version of an emoji.")

	if detailed {
		help = append(help, bruxism.CommandHelp(service, "hugemoji", "<emoji>", "Returns a huge version of an emoji.")[0])
	}

	return help
}

// New creates a new emoji plugin.
func New() bruxism.Plugin {
	p := bruxism.NewSimplePlugin("Emoji")
	p.LoadFunc = emojiLoadFunc
	p.MessageFunc = emojiMessageFunc
	p.HelpFunc = emojiHelpFunc
	return p
}
