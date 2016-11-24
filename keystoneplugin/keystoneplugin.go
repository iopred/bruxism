package keystoneplugin

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/iopred/bruxism"
)

type keystoneDungeonID int

const (
	BRH keystoneDungeonID = iota
	COS
	DHT
	EOA
	HOV
	MOS
	NL
	ARC
	VOW
)

type keystoneDungeon struct {
	Name    string
	Aliases []string
}

var dungeons map[keystoneDungeonID]*keystoneDungeon = map[keystoneDungeonID]*keystoneDungeon{
	BRH: &keystoneDungeon{
		Name:    "Black Rook Hold",
		Aliases: []string{"black rook hold", "brh", "black rook", "hold"},
	},
	COS: &keystoneDungeon{
		Name:    "Court of Stars",
		Aliases: []string{"court of stars", "cos", "court"},
	},
	DHT: &keystoneDungeon{
		Name:    "Darkheart Thicket",
		Aliases: []string{"darkheart thicket", "dht", "thicket", "darkheart"},
	},
	EOA: &keystoneDungeon{
		Name:    "Eye of Azshara",
		Aliases: []string{"eye of azshara", "eoa", "eye", "azshara"},
	},
	HOV: &keystoneDungeon{
		Name:    "Halls of Valor",
		Aliases: []string{"halls of valor", "hall of valor", "hov", "halls"},
	},
	MOS: &keystoneDungeon{
		Name:    "Maw of Souls",
		Aliases: []string{"maw of souls", "mos", "maw"},
	},
	NL: &keystoneDungeon{
		Name:    "Neltharion's Lair",
		Aliases: []string{"neltharion's lair", "nl", "neltharions lair", "nel", "nelth", "lair"},
	},
	ARC: &keystoneDungeon{
		Name:    "The Arcway",
		Aliases: []string{"the arcway", "arc", "arcway"},
	},
	VOW: &keystoneDungeon{
		Name:    "Vault of the Wardens",
		Aliases: []string{"vault of the wardens", "vow", "vault", "warden", "wardens"},
	},
}

type keystone struct {
	User      string
	Dungeon   keystoneDungeonID
	Level     int
	Depleted  bool
	Modifiers []string
}

func (k *keystone) String() string {
	str := fmt.Sprintf("Level %d **%s**", k.Level, dungeons[k.Dungeon].Name)
	if len(k.Modifiers) != 0 {
		str += " *(" + strings.Join(k.Modifiers, ", ") + ")*"
	}
	if k.Depleted {
		str += " - Depleted"
	}
	return str
}

type keystoneChannel struct {
	Users        map[string]*keystone
	LastModified time.Time
}

var location *time.Location

func init() {
	var err error
	location, err = time.LoadLocation("America/Los_Angeles")
	if err != nil {
		location = time.Now().Location()
	}
}

func lastTuesday(t time.Time) time.Time {
	year, month, day := t.Date()
	t = time.Date(year, month, day, 0, 0, 0, 0, location)
	for t.Weekday() != time.Tuesday {
		t = t.Add(-24 * time.Hour)
	}

	return t
}

func (c *keystoneChannel) check() {
	if time.Now().After(lastTuesday(c.LastModified).Add(24 * 7 * time.Hour)) {
		c.Users = map[string]*keystone{}
	}
}

func (c *keystoneChannel) add(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, userName, query string) bool {
	query = strings.ToLower(query)
	for dungeonID, dungeon := range dungeons {
		for _, alias := range dungeon.Aliases {
			if strings.Index(query, alias+" ") == 0 {
				query = query[len(alias)+1:]
				if len(query) == 0 {
					return false
				}

				parts := strings.Split(query, " ")

				level, err := strconv.Atoi(parts[0])
				if err != nil {
					return false
				}

				parts = parts[1:]

				depleted := false
				modifiers := []string{}

				for _, part := range parts {
					if strings.Index(part, "deplete") == 0 {
						depleted = true
					} else {
						modifiers = append(modifiers, part)
					}
				}

				c.Users[message.UserID()] = &keystone{
					User:      userName,
					Dungeon:   dungeonID,
					Level:     level,
					Depleted:  depleted,
					Modifiers: modifiers,
				}
				c.LastModified = time.Now()

				return true
			}
		}
	}
	return false
}

type keystoneList []*keystone

// Len is part of sort.Interface.
func (s keystoneList) Len() int {
	return len(s)
}

// Swap is part of sort.Interface.
func (s keystoneList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s keystoneList) Less(i, j int) bool {
	if s[i].Level == s[j].Level {
		return dungeons[s[i].Dungeon].Name > dungeons[s[j].Dungeon].Name
	}
	return s[i].Level > s[j].Level
}

func (c *keystoneChannel) list(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	if len(c.Users) == 0 {
		service.SendMessage(message.Channel(), "No keystones have been set this week.")
		return
	}

	keystones := keystoneList{}
	for _, keystone := range c.Users {
		keystones = append(keystones, keystone)
	}

	sort.Sort(keystones)

	content := ""
	for _, keystone := range keystones {
		if len(content) != 0 {
			content += "\n"
		}
		content += keystone.String()
		content += " - " + keystone.User
	}
	service.SendMessage(message.Channel(), content)
}

type keystonePlugin struct {
	sync.RWMutex
	Channels map[string]*keystoneChannel
}

// Name returns the name of the plugin.
func (p *keystonePlugin) Name() string {
	return "Keystone"
}

// Load will load plugin state from a byte array.
func (p *keystonePlugin) Load(bot *bruxism.Bot, service bruxism.Service, data []byte) error {
	if data != nil {
		if err := json.Unmarshal(data, p); err != nil {
			log.Println("Error loading data", err)
		}
	}

	if p.Channels == nil {
		p.Channels = make(map[string]*keystoneChannel)
	}

	return nil
}

// Save will save plugin state to a byte array.
func (p *keystonePlugin) Save() ([]byte, error) {
	return json.Marshal(p)
}

// Help returns a list of help strings that are printed when the user requests them.
func (p *keystonePlugin) Help(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, detailed bool) []string {
	help := []string{}

	if service.IsBotOwner(message) || service.IsModerator(message) {
		if p.Channels[message.Channel()] == nil {
			help = append(help, bruxism.CommandHelp(service, "start", "", "Starts keystone tracking in this channel.")[0])
		} else {
			help = append(help, bruxism.CommandHelp(service, "stop", "", "Stops keystone tracking in this channel.")[0])
		}
	}

	ticks := ""
	if service.Name() == bruxism.DiscordServiceName {
		ticks = "`"
	}

	if p.Channels[message.Channel()] != nil {
		help = append(help, []string{
			bruxism.CommandHelp(service, "set", "<dungeon> <level> [modifiers]", "Set your keystone for this week.")[0],
			bruxism.CommandHelp(service, "list", "", "Lists all this weeks keystones.")[0],
			bruxism.CommandHelp(service, "deplete", "", "Depletes your keystone")[0],
			bruxism.CommandHelp(service, "undeplete", "", "Undepletes your keystone")[0],
		}...)
	}

	if detailed {
		help = append(help, []string{
			"Examples:",
			fmt.Sprintf("%s%sset hov 5 teeming%s - Adds a Level 5 Halls of Valor keystone with teeming.", ticks, service.CommandPrefix(), ticks),
			fmt.Sprintf("%s%sset eye of azshara 2 depleted%s - Adds a depleted Level 2 Eye of Azshara keystone.", ticks, service.CommandPrefix(), ticks),
		}...)
	}

	if len(help) == 0 {
		return nil
	}
	return help
}

// Message handler.
func (p *keystonePlugin) Message(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	defer bruxism.MessageRecover()
	if !service.IsMe(message) {
		messageChannel := message.Channel()

		var userName string
		discord, ok := service.(*bruxism.Discord)
		if ok {
			userName = discord.Nickname(message)
		} else {
			userName = message.UserName()
		}

		if (service.IsBotOwner(message) || service.IsModerator(message)) && (bruxism.MatchesCommand(service, "start", message) || bruxism.MatchesCommand(service, "stop", message)) {
			p.Lock()
			defer p.Unlock()

			if bruxism.MatchesCommand(service, "start", message) {
				p.Channels[messageChannel] = &keystoneChannel{
					Users: map[string]*keystone{},
				}
				service.SendMessage(messageChannel, "This channel is now tracking mythic keystones.")
			} else {
				delete(p.Channels, messageChannel)
				service.SendMessage(messageChannel, "This channel is no longer tracking mythic keystones.")
			}
		} else if channel, ok := p.Channels[messageChannel]; ok {
			channel.check()

			if bruxism.MatchesCommand(service, "set", message) {
				query, parts := bruxism.ParseCommand(service, message)
				if len(parts) > 1 && channel.add(bot, service, message, userName, query) {
					channel.list(bot, service, message)
				} else {
					service.SendMessage(messageChannel, "Invalid keystone. Eg: `keystone set hall of valor 3 sanguine`")
				}
			} else if bruxism.MatchesCommand(service, "list", message) {
				channel.list(bot, service, message)
			} else if bruxism.MatchesCommand(service, "deplete", message) {
				keystone := channel.Users[message.UserID()]
				if keystone == nil {
					service.SendMessage(messageChannel, "You haven't set a keystone this week.")
				} else {
					keystone.Depleted = true
					keystone.User = userName
					service.SendMessage(messageChannel, fmt.Sprintf("Keystone depleted: %s", keystone.String()))
				}
			} else if bruxism.MatchesCommand(service, "undeplete", message) {
				keystone := channel.Users[message.UserID()]
				if keystone == nil {
					service.SendMessage(messageChannel, "You haven't set a keystone this week.")
				} else {
					keystone.Depleted = false
					keystone.User = userName
					service.SendMessage(messageChannel, fmt.Sprintf("Keystone undepleted: %s", keystone.String()))
				}
			}
		}
	}
}

// Stats will return the stats for a plugin.
func (p *keystonePlugin) Stats(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) []string {
	return nil
}

// New will create a new wormhole plugin.
func New() bruxism.Plugin {
	return &keystonePlugin{}
}
