package playedplugin

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/iopred/bruxism"
	"github.com/iopred/discordgo"
)

type playedEntry struct {
	Name     string
	Duration time.Duration
}

type byDuration []*playedEntry

func (a byDuration) Len() int           { return len(a) }
func (a byDuration) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byDuration) Less(i, j int) bool { return a[i].Duration < a[j].Duration }

type playedUser struct {
	Entries     map[string]*playedEntry
	Current     string
	LastChanged time.Time
	FirstSeen   time.Time
}

func (p *playedUser) Update(name string, now time.Time) {
	if p.Current != "" {
		pe := p.Entries[p.Current]
		if pe == nil {
			pe = &playedEntry{
				Name: p.Current,
			}
			p.Entries[p.Current] = pe
		}
		pe.Duration += now.Sub(p.LastChanged)
	}

	p.Current = name
	p.LastChanged = now
}

type playedPlugin struct {
	sync.RWMutex
	Users map[string]*playedUser
}

// Load will load plugin state from a byte array.
func (p *playedPlugin) Load(bot *bruxism.Bot, service bruxism.Service, data []byte) error {
	if service.Name() != bruxism.DiscordServiceName {
		panic("Carbonitex Plugin only supports Discord.")
	}

	if data != nil {
		if err := json.Unmarshal(data, p); err != nil {
			log.Println("Error loading data", err)
		}
	}

	go p.Run(bot, service)
	return nil
}

// Save will save plugin state to a byte array.
func (p *playedPlugin) Save() ([]byte, error) {
	return json.Marshal(p)
}

func (p *playedPlugin) Update(user string, entry string) {
	fmt.Println(user, entry)
	p.Lock()
	defer p.Unlock()

	t := time.Now()

	u := p.Users[user]
	if u == nil {
		u = &playedUser{
			Entries:     map[string]*playedEntry{},
			Current:     entry,
			LastChanged: t,
			FirstSeen:   t,
		}
		p.Users[user] = u
	}
	u.Update(entry, t)
}

// Run is the background go routine that executes for the life of the plugin.
func (p *playedPlugin) Run(bot *bruxism.Bot, service bruxism.Service) {
	discord := service.(*bruxism.Discord)

	discord.Session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		for _, g := range r.Guilds {
			for _, pu := range g.Presences {
				e := ""
				if pu.Game != nil {
					e = pu.Game.Name
				}
				p.Update(pu.User.ID, e)
			}
		}
	})

	discord.Session.AddHandler(func(s *discordgo.Session, g *discordgo.GuildCreate) {
		if g.Unavailable == nil || *g.Unavailable {
			return
		}

		for _, pu := range g.Presences {
			e := ""
			if pu.Game != nil {
				e = pu.Game.Name
			}
			p.Update(pu.User.ID, e)
		}
	})

	discord.Session.AddHandler(func(s *discordgo.Session, pr *discordgo.PresencesReplace) {
		for _, pu := range *pr {
			e := ""
			if pu.Game != nil {
				e = pu.Game.Name
			}
			p.Update(pu.User.ID, e)
		}
	})

	discord.Session.AddHandler(func(s *discordgo.Session, pu *discordgo.PresenceUpdate) {
		e := ""
		if pu.Game != nil {
			e = pu.Game.Name
		}
		p.Update(pu.User.ID, e)
	})
}

// Help returns a list of help strings that are printed when the user requests them.
func (p *playedPlugin) Help(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, detailed bool) []string {
	if detailed {
		return nil
	}

	return bruxism.CommandHelp(service, "played", "[@username]", "Returns your most played games, or a users most played games if provided.")
}

var userIDRegex = regexp.MustCompile("<@!?([0-9]*)>")

func (p *playedPlugin) Message(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	defer bruxism.MessageRecover()
	if service.Name() == bruxism.DiscordServiceName && !service.IsMe(message) {
		if bruxism.MatchesCommand(service, "played", message) {
			query := strings.Join(strings.Split(message.RawMessage(), " ")[1:], " ")

			id := message.UserID()
			match := userIDRegex.FindStringSubmatch(query)
			if match != nil {
				id = match[1]
			}

			p.Lock()
			defer p.Unlock()

			u := p.Users[id]
			if u == nil {
				service.SendMessage(message.Channel(), "I haven't seen that user.")
				return
			}

			u.Update(u.Current, time.Now())

			pes := make(byDuration, len(u.Entries))
			i := 0
			for _, pe := range u.Entries {
				pes[i] = pe
				i++
			}

			sort.Sort(pes)

			messageText := fmt.Sprintf("First seen %s.\n", humanize.Time(u.FirstSeen))
			for i = 0; i < len(pes) && i < 5; i++ {
				pe := pes[i]
				messageText += fmt.Sprintf("%s: %.0f hours, %.0f minutes, %.0f seconds\n", pe.Name, pe.Duration.Hours(), pe.Duration.Minutes(), pe.Duration.Seconds())
			}
			service.SendMessage(message.Channel(), messageText)
		}
	}
}

// Name returns the name of the plugin.
func (p *playedPlugin) Name() string {
	return "Played"
}

// New will create a played plugin.
func New() bruxism.Plugin {
	return &playedPlugin{
		Users: map[string]*playedUser{},
	}
}
