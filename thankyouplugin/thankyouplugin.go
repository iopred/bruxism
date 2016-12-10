package thankyouplugin

import (
	"encoding/json"
	"log"
	"strings"
	"sync"

	"github.com/iopred/bruxism"
)

type thankyouPlugin struct {
	sync.Mutex
	bruxism.SimplePlugin
	Thanks int
}

// Load will load plugin state from a byte array.
func (p *thankyouPlugin) Load(bot *bruxism.Bot, service bruxism.Service, data []byte) error {
	if data != nil {
		if err := json.Unmarshal(data, p); err != nil {
			log.Println("Error loading data", err)
		}
	}

	return nil
}

// Save will save plugin state to a byte array.
func (p *thankyouPlugin) Save() ([]byte, error) {
	return json.Marshal(p)
}

// Help returns a list of help strings that are printed when the user requests them.
func (p *thankyouPlugin) Help(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, detailed bool) []string {
	if detailed {
		return nil
	}

	return bruxism.CommandHelp(service, "thankyou", "", "Thank you to all my patrons.")
}

// Message handler.
func (p *thankyouPlugin) Message(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	if service.IsMe(message) || !bruxism.MatchesCommand(service, "thankyou", message) {
		return
	}

	p.Lock()
	p.Thanks++
	p.Unlock()

	out := "Septapus currently has no patrons.\nBecome a patron at: <https://www.patreon.com/iopred>"

	if service.SupportsMultiline() {
		service.SendMessage(message.Channel(), out)
	} else {
		service.SendMessage(message.Channel(), strings.Join(strings.Split(out, "\n"), " "))
	}
}

// Name returns the name of the plugin.
func (p *thankyouPlugin) Name() string {
	return "Thankyou"
}

// New will create a new thankyou plugin.
func New() bruxism.Plugin {
	return &thankyouPlugin{}
}
