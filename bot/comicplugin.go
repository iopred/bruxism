package bot

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/iopred/comicgen"
)

var comic *comicgen.ComicGen

func init() {
	comic, _ = comicgen.NewComicGen()
}

type comicPlugin struct {
	SimplePlugin
	log map[string][]Message
}

func (p *comicPlugin) helpFunc(bot *Bot, service Service) []string {
	return commandHelp("comic", "<lines>", "Creates a comic from the last lines of chat.")
}

func (p *comicPlugin) messageFunc(bot *Bot, service Service, message Message) {
	log, ok := p.log[message.Channel()]
	if !ok {
		log = []Message{}
	}

	if matchesCommand("comic", message) && !service.IsMe(message) {
		if len(log) == 0 {
			return
		}

		var err error
		lines := 0
		linesString, parts := parseCommand(message)
		if len(parts) > 0 {
			lines, _ = strconv.Atoi(linesString)
		}

		if lines == 0 {
			lines = rand.Intn(comic.MaxLines())
		}

		if lines > len(log) {
			lines = len(log)
		}

		speakers := make(map[string]int)
		claimed := make(map[int]bool)
		avatars := make(map[int]string)

		messages := log[len(log)-lines:]

		script := []*comicgen.Message{}

		for _, message := range messages {
			speaker := 0
			if _, ok := speakers[message.UserName()]; !ok {
				for {
					speaker = rand.Intn(comic.Avatars())
					if _, ok := claimed[speaker]; !ok {
						claimed[speaker] = true

						a, err := service.GetAvatar(message)
						if err == nil {
							avatars[speaker] = a
						}

						break
					}
				}
				speakers[message.UserName()] = speaker
			} else {
				speaker = speakers[message.UserName()]
			}

			script = append(script, &comicgen.Message{
				Speaker: speaker,
				Text:    message.Message(),
			})
		}

		image, err := comic.MakeComic(&comicgen.Script{
			Messages: script,
			Author:   fmt.Sprintf("%s and %s", service.UserName(), message.UserName()),
			Avatars:  avatars,
		})

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
