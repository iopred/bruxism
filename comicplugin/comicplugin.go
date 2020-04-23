package comicplugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"log"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"github.com/iopred/bruxism"
	"github.com/iopred/comicgen"
)

var comicEmoji = "<:comic:701265673707454494>"
var urinalEmoji = "<:urinal:701279094402318387>"

type comicPlugin struct {
	sync.Mutex

	bruxism.SimplePlugin
	log      map[string][]bruxism.Message
	cooldown map[string]time.Time
	Count	 map[string]int
	Comics   int
	Public   map[string]bool
}

// Load will load plugin state from a byte array.
func (p *comicPlugin) Load(bot *bruxism.Bot, service bruxism.Service, data []byte) error {
	if data != nil {
		if err := json.Unmarshal(data, p); err != nil {
			log.Println("Error loading data", err)
		}
	}

	p.cooldown = map[string]time.Time{}
	if p.Public == nil {
		p.Public = map[string]bool{}
	}
	if p.Count == nil {
		p.Count = map[string]int{}
	}

	return nil
}

// Save will save plugin state to a byte array.
func (p *comicPlugin) Save() ([]byte, error) {
	return json.Marshal(p)
}

// Help returns a list of help strings that are printed when the user requests them.
func (p *comicPlugin) Help(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, detailed bool) []string {
	help := bruxism.CommandHelp(service, "comic", "[1-10]", "Creates a comic from recent messages, or a number of messages if provided. eg: comic 3")

	ticks := ""
	if service.Name() == bruxism.DiscordServiceName {
		ticks = "`"
	}
	if detailed {
		help = append(help, []string{
			bruxism.CommandHelp(service, "customcomic", "[id|name:] <text> | [id|name:] <text>", fmt.Sprintf("Creates a custom comic. Available names: %s%s%s", ticks, strings.Join(comicgen.CharacterNames, ", "), ticks))[0],
			bruxism.CommandHelp(service, "customcomicsimple", "[id:] <text> | [id:] <text>", "Creates a simple custom comic.")[0],
		}...)

		if service.Name() == bruxism.DiscordServiceName && service.IsChannelOwner(message) {
			help = append(help, []string{
				bruxism.CommandHelp(service, "comicpublic", "", "Makes comics in this server public.")[0],
				bruxism.CommandHelp(service, "comicprivate", "", "Makes comics in this server private (default).")[0],
			}...)
		}

		help = append(help, []string{
			"Examples:",
			bruxism.CommandHelp(service, "comic", "5", "Creates a comic from the last 5 messages")[0],
			bruxism.CommandHelp(service, "customcomic", "A | B | C", "Creates a comic with 3 lines.")[0],
			bruxism.CommandHelp(service, "customcomic", "0: Hi! | 1: Hello! | 0: Goodbye.", "Creates a comic with 3 lines, the second line spoken by a different character")[0],
			bruxism.CommandHelp(service, "customcomic", "tiki: Hi! | jordy: Hello! | tiki: Goodbye.", "Creates a comic with 3 lines, containing tiki and jordy.")[0],
			bruxism.CommandHelp(service, "customcomicsimple", "0: Foo | 1: Bar", "Creates a comic with 2 lines, both spoken by different characters.")[0],
		}...)
	}

	return help
}

var emojiRegexp = regexp.MustCompile("<a?:[a-zA-Z0-9]+:([0-9]+)>")
var imageRegexp = regexp.MustCompile("^(http(s?):)([/|.|\\w|\\s|-])*\\.(?:jpg|jpeg|gif|png)($|\\?\\S*$)")

func makeScriptFromMessages(service bruxism.Service, message bruxism.Message, messages []bruxism.Message, room string) *comicgen.Script {
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

		replacements := map[string]string{}
		messageMessage := message.Message()
		matches := emojiRegexp.FindAllStringSubmatch(messageMessage, -1)
		for _, match := range matches {
			if _, ok := replacements[match[0]]; !ok {
				replacements[match[0]] = discordgo.EndpointEmoji(match[1])
			}
		}

		imageUrl := ""
		if imageRegexp.MatchString(messageMessage) {
			imageUrl = messageMessage
			messageMessage = ""
		}

		if dm, ok := message.(*bruxism.DiscordMessage); ok {
			at := dm.AttachmentURL()
			if at != "" {
				imageUrl = at
			}
		}

		script = append(script, &comicgen.Message{
			Speaker:      speaker,
			Text:         messageMessage,
			Author:       message.UserName(),
			Replacements: replacements,
			ImageUrl:	  imageUrl,
		})
	}
	return &comicgen.Script{
		Messages: script,
		Author:   fmt.Sprintf(service.UserName()),
		Avatars:  avatars,
		Type:     comicgen.ComicTypeChat,
		Room:	  room,
	}
}

func (p *comicPlugin) makeComic(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, globalID string, script *comicgen.Script) {
	p.Comics++
	comic := comicgen.NewComicGen("arial", service.Name() != bruxism.DiscordServiceName)
	image, err := comic.MakeComic(script)
	if err != nil {
		service.SendMessage(message.Channel(), fmt.Sprintf("Sorry %s, there was an error creating the comic. %s", message.UserName(), err))
	} else {
		go func() {
			if service.Name() == bruxism.DiscordServiceName {
				discord := service.(*bruxism.Discord)
				pe, err := discord.UserChannelPermissions(message.UserID(), message.Channel())
				if err == nil && pe&discordgo.PermissionAttachFiles != 0 {
					b := &bytes.Buffer{}
					err = png.Encode(b, image)
					if err != nil {
						service.SendMessage(message.Channel(), fmt.Sprintf("Sorry %s, there was a problem creating your comic.", message.UserName()))
						return
					}

					m, err := discord.Session.ChannelFileSend(message.Channel(), "comic.png", b)
					if err == nil {
						if p.Public[globalID] && len(m.Attachments) > 0 {
							service.SendMessage("367871992503730176", m.Attachments[0].URL)
						}
						p.log[message.Channel()] = nil
						p.Count[message.Channel()] = 0
						return
					}
				}
			}

			b := &bytes.Buffer{}
			err = png.Encode(b, image)
			if err != nil {
				service.SendMessage(message.Channel(), fmt.Sprintf("Sorry %s, there was a problem creating your comic.", message.UserName()))
				return
			}

			url, err := bot.UploadToImgur(b, "comic.png")
			if err != nil {
				service.SendMessage(message.Channel(), fmt.Sprintf("Sorry %s, there was a problem uploading the comic to imgur.", message.UserName()))
				log.Println("Error uploading comic: ", err)
				return
			}

			if service.Name() == bruxism.DiscordServiceName {
				service.SendMessage(message.Channel(), fmt.Sprintf("Here's your comic <@%s>: %s", message.UserID(), url))
			} else {
				service.SendMessage(message.Channel(), fmt.Sprintf("Here's your comic %s: %s", message.UserName(), url))
			}

			if p.Public[globalID] && service.Name() == bruxism.DiscordServiceName {
				service.SendMessage("367871992503730176", url)
			}
			p.log[message.Channel()] = nil
			p.Count[message.Channel()] = 0
		}()
	}
}

func (p *comicPlugin) checkCooldown(service bruxism.Service, message bruxism.Message) bool {
	cooldown := p.cooldown[message.UserID()]
	if cooldown.After(time.Now()) {
		service.SendMessage(message.Channel(), fmt.Sprintf("Sorry %s, you need to wait %s before creating another comic.", message.UserName(), humanize.Time(cooldown)))
		return true
	}
	p.cooldown[message.UserID()] = time.Now().Add(10 * time.Second)
	return false
}

// Message handler.
func (p *comicPlugin) Message(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	defer bruxism.MessageRecover()

	if service.IsMe(message) {
		return
	}

	p.Lock()
	defer p.Unlock()

	log, ok := p.log[message.Channel()]
	if !ok {
		log = []bruxism.Message{}
	}
	count := p.Count[message.Channel()]

	globalID := message.Channel()
	if discord, ok := service.(*bruxism.Discord); ok {
		channel, err := discord.Channel(globalID)
		if err == nil {
			globalID = channel.GuildID
		}
	}

	emoji, ok := message.MatchesCommand("", comicEmoji)
	if !ok {
		emoji = false
	}

	urinal, ok := message.MatchesCommand("", urinalEmoji)
	if !ok {
		urinal = false
	}

	if service.Name() == bruxism.DiscordServiceName && bruxism.MatchesCommand(service, "comicpublic", message) && service.IsChannelOwner(message) {
		p.Public[globalID] = true
		service.SendMessage(message.Channel(), "Comics are now public. Visit https://discord.gg/HWN9pwj to see them all.")
	} else if service.Name() == bruxism.DiscordServiceName && bruxism.MatchesCommand(service, "comicprivate", message) && service.IsChannelOwner(message) {
		delete(p.Public, globalID)
		service.SendMessage(message.Channel(), "Comics are no longer public.")
	} else if bruxism.MatchesCommand(service, "customcomic", message) || bruxism.MatchesCommand(service, "customcomicsimple", message) {

		ty := comicgen.ComicTypeChat
		if bruxism.MatchesCommand(service, "customcomicsimple", message) {
			ty = comicgen.ComicTypeSimple
		}

		service.Typing(message.Channel())

		str, _ := bruxism.ParseCommand(service, message)

		messages := []*comicgen.Message{}

		splits := strings.Split(str, "|")
		if len(splits) == 0 || (len(splits) == 1 && len(splits[0]) == 0) {
			service.SendMessage(message.Channel(), fmt.Sprintf("Sorry %s, you didn't add any text.", message.UserName()))
			return
		}
		
		if p.checkCooldown(service, message) {
			return
		}
		
		if len(splits) > 10 {
			splits = splits[:10]
		}
		for _, line := range splits {
			line := strings.Trim(line, " ")

			text := ""
			speaker := 0
			author := ""
			if strings.Index(line, ":") != -1 {
				lineSplit := strings.Split(line, ":")

				author = strings.Trim(lineSplit[0], " ")

				var err error
				speaker, err = strconv.Atoi(author)
				if err != nil {
					speaker = -1
				}

				text = strings.Trim(strings.Join(lineSplit[1:], ":"), " ")
			} else {
				text = line
			}

			messages = append(messages, &comicgen.Message{
				Speaker: speaker,
				Text:    text,
				Author:  author,
			})
		}

		p.makeComic(bot, service, message, globalID, &comicgen.Script{
			Messages: messages,
			Author:   fmt.Sprintf(service.UserName()),
			Type:     ty,
		})
	} else if bruxism.MatchesCommand(service, "comic", message) || emoji || urinal {
		if len(log) == 0 && count == 0{
			service.SendMessage(message.Channel(), fmt.Sprintf("Sorry %s, I don't have enough messages to make a comic yet.", message.UserName()))
			return
		}

		if p.checkCooldown(service, message) {
			return
		}
		service.Typing(message.Channel())

		lines := 0
		
		_, parts := bruxism.ParseCommand(service, message)

		if emoji || urinal {
			_ , parts, _ = message.ParseCommand("")
		}

		if len(parts) > 0 {
			lines, _ = strconv.Atoi(parts[0])
		}

		if lines <= 0 {
			lines = 1 + int(math.Floor((math.Pow(2*rand.Float64()-1, 3)/2+0.5)*float64(5)))
		}

		var messages []bruxism.Message
		if len(log) != 0 {
			if lines > len(log) {
				lines = len(log)
			}
			messages = log[len(log)-lines:]
		} else if count != 0 {
			if lines > count {
				lines = count
			}
			if lines > 10 {
				lines = 10
			}
			if discord, ok := service.(*bruxism.Discord); ok {
				ms, err := discord.Session.ChannelMessages(message.Channel(), lines, message.MessageID(), "", "")
				if err != nil {
					service.SendMessage(message.Channel(), fmt.Sprintf("Sorry %s, there was an error creating the comic.", message.UserName()))
					return
				}
				if len(ms) < lines {
					lines = len(ms)
				}
				messages = make([]bruxism.Message, lines)
				for i := 0; i < lines; i++ {
					messages[lines - i - 1] = &bruxism.DiscordMessage{
						Discord:          discord,
						DiscordgoMessage: ms[i],
						MessageType:      bruxism.MessageTypeDelete,
					}
				}
			} 
		}

		room := ""
		for _, u := range parts {
			u = strings.ToLower(u)
			if u == "urinal" || u == "toilet" {
				urinal = true
			}
		}

		if urinal {
			room = "rooms/urinal.png"
		}

		p.makeComic(bot, service, message, globalID, makeScriptFromMessages(service, message, messages, room))
	} else {
		messageHistory := false
		if service.Name() == bruxism.DiscordServiceName {
			pe, err := service.(*bruxism.Discord).UserChannelPermissions(message.UserID(), message.Channel())
			if err == nil && pe&discordgo.PermissionReadMessageHistory != 0 {
				// messageHistory = true
			}
		}
		switch message.Type() {
		case bruxism.MessageTypeCreate:
			if messageHistory {
				p.Count[message.Channel()]++
				return
			}
			if len(log) < 10 {
				log = append(log, message)
			} else {
				log = append(log[1:], message)
			}
		case bruxism.MessageTypeUpdate:
			if messageHistory{
				return
			}
			for i, m := range log {
				if m.MessageID() == message.MessageID() {
					log[i] = message
					break
				}
			}
		case bruxism.MessageTypeDelete:
			if messageHistory {
				p.Count[message.Channel()]--
				return
			}
			for i, m := range log {
				if m.MessageID() == message.MessageID() {
					log = append(log[:i], log[i+1:]...)
					break
				}
			}
		}
		p.log[message.Channel()] = log
	}
}

// Name returns the name of the plugin.
func (p *comicPlugin) Name() string {
	return "Comic"
}

// Stats will return the stats for a plugin.
func (p *comicPlugin) Stats(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) []string {
	return []string{fmt.Sprintf("Comics created: \t%d\n", p.Comics)}
}

// New will create a new comic plugin.
func New() bruxism.Plugin {
	return &comicPlugin{
		log: make(map[string][]bruxism.Message),
		Count: make(map[string]int),
	}
}
