package bruxism

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"

	"github.com/iopred/comicgen"
)

var comic *comicgen.ComicGen

func init() {
	comic, _ = comicgen.NewComicGen("arial")
}

type comicPlugin struct {
	SimplePlugin
	log map[string][]Message
}

func (p *comicPlugin) helpFunc(bot *Bot, service Service) []string {
	return []string{
		commandHelp(service, "comic", "[<1-5>]", "Creates a comic from recent messages.")[0],
		commandHelp(service, "customcomic", "[<id>:] <text> | [<id>:] <text>", "Creates a custom comic, eg: customcomic Hello | 1: World!")[0],
	}
}

func makeScriptFromMessages(service Service, message Message, messages []Message) *comicgen.Script {
	speakers := make(map[string]int)
	avatars := make(map[int]string)

	script := []*comicgen.Message{}

	for _, message := range messages {
		speaker, ok := speakers[message.UserName()]
		if !ok {
			speaker = len(speakers)
			speakers[message.UserName()] = speaker
			avatars[speaker] = message.UserAvatar()
		}

		script = append(script, &comicgen.Message{
			Speaker: speaker,
			Text:    message.Message(),
		})
	}
	return &comicgen.Script{
		Messages: script,
		Author:   fmt.Sprintf("%s and %s", service.UserName(), message.UserName()),
		Avatars:  avatars,
	}
}

func (p *comicPlugin) makeComic(bot *Bot, service Service, message Message, script *comicgen.Script) {
	image, err := comic.MakeComic(script)

	if err != nil {
		service.SendMessage(message.Channel(), fmt.Sprintf("Sorry %s, there was an error creating the comic: %s", err))
	} else {
		go func() {
			url, err := bot.UploadToImgur(image, "comic.png")
			if err == nil {
				if service.Name() == DiscordServiceName {
					service.SendMessage(message.Channel(), fmt.Sprintf("Here's your comic <@%s>: %s", message.UserID(), url))
				} else {
					service.SendMessage(message.Channel(), fmt.Sprintf("Here's your comic %s: %s", message.UserName(), url))
				}
			} else {
				service.SendMessage(message.Channel(), fmt.Sprintf("Sorry %s, there was a problem uploading the comic to imgur.", message.UserName()))
			}
		}()
	}
}

func (p *comicPlugin) messageFunc(bot *Bot, service Service, message Message) {
	if service.IsMe(message) {
		return
	}

	log, ok := p.log[message.Channel()]
	if !ok {
		log = []Message{}
	}

	if matchesCommand(service, "customcomic", message) {
		service.Typing(message.Channel())

		str, _ := parseCommand(service, message)

		avatars := map[int]string{0: message.UserAvatar()}
		messages := []*comicgen.Message{}

		splits := strings.Split(str, "|")
		for _, line := range splits {
			line := strings.Trim(line, " ")

			text := ""
			speaker := 0
			if strings.Index(line, ":") != -1 {
				lineSplit := strings.Split(line, ":")

				speaker, _ = strconv.Atoi(strings.Trim(lineSplit[0], " "))
				if speaker < 0 {
					speaker = 0
				}

				text = strings.Trim(lineSplit[1], " ")
			} else {
				text = line
			}

			messages = append(messages, &comicgen.Message{
				Speaker: speaker,
				Text:    text,
			})
		}

		if len(messages) == 0 {
			service.SendMessage(message.Channel(), fmt.Sprintf("Sorry %s, you didn't add any text.", message.UserName()))
			return
		}

		p.makeComic(bot, service, message, &comicgen.Script{
			Messages: messages,
			Author:   fmt.Sprintf("%s and %s", service.UserName(), message.UserName()),
			Avatars:  avatars,
		})
	} else if matchesCommand(service, "comic", message) {
		if len(log) == 0 {
			return
		}

		service.Typing(message.Channel())

		lines := 0
		linesString, parts := parseCommand(service, message)
		if len(parts) > 0 {
			lines, _ = strconv.Atoi(linesString)
		}

		if lines <= 0 {
			lines = 1 + int(math.Floor((math.Pow(2*rand.Float64()-1, 3)/2+0.5)*float64(comic.MaxLines()-1)))
		}

		if lines > len(log) {
			lines = len(log)
		}

		p.makeComic(bot, service, message, makeScriptFromMessages(service, message, log[len(log)-lines:]))
	} else {
		if len(log) < 10 {
			p.log[message.Channel()] = append(log, message)
		} else {
			p.log[message.Channel()] = append(log[1:], message)
		}
	}
}

// NewComicPlugin will create a new top streamers plugin.
func NewComicPlugin() Plugin {
	if comic == nil {
		return nil
	}

	p := &comicPlugin{
		SimplePlugin: *NewSimplePlugin("Comic"),
		log:          make(map[string][]Message),
	}
	p.message = p.messageFunc
	p.help = p.helpFunc
	return p
}
