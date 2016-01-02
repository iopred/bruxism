package bot

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

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
	return commandHelp(service, "comic", "[<1-5>]", "Creates a comic from recent messages.")
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

func (p *comicPlugin) messageFunc(bot *Bot, service Service, message Message) {
	log, ok := p.log[message.Channel()]
	if !ok {
		log = []Message{}
	}

	if matchesCommand(service, "comic", message) && !service.IsMe(message) {
		if len(log) == 0 {
			return
		}

		service.Typing(message.Channel())

		var err error
		lines := 0
		linesString, parts := parseCommand(service, message)
		if len(parts) > 0 {
			lines, err = strconv.Atoi(linesString)
		}

		if lines == 0 {
			lines = 1 + int(math.Floor((math.Pow(2*rand.Float64()-1, 3)/2+0.5)*float64(comic.MaxLines()-1)))
		}

		if lines > len(log) {
			lines = len(log)
		}

		image, err := comic.MakeComic(makeScriptFromMessages(service, message, log[len(log)-lines:]))

		if err != nil {
			fmt.Println("Error creating comic", err)
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
					fmt.Println("Error uploading comic", err)
				}
			}()
		}
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
