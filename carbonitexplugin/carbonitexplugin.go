package carbonitexplugin

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/iopred/bruxism"
)

type carbonitexPlugin struct {
	bruxism.SimplePlugin
	key string
}

func (p *carbonitexPlugin) carbonitexPluginLoadFunc(bot *bruxism.Bot, service bruxism.Service, data []byte) error {
	if service.Name() != bruxism.DiscordServiceName {
		panic("Carbonitex Plugin only supports Discord.")
	}

	go p.Run(bot, service)
	return nil
}

// Run will poll YouTube for channels going live and send messages.
func (p *carbonitexPlugin) Run(bot *bruxism.Bot, service bruxism.Service) {
	for {
		<-time.After(10 * time.Second)

		body, err := json.Marshal(struct {
			Key         string `json:"key"`
			ServerCount int    `json:"servercount"`
		}{p.key, service.ChannelCount()})
		if err != nil {
			return
		}

		http.Post("https://www.carbonitex.net/discord/data/botdata.php", "application/json", bytes.NewBuffer(body))

		<-time.After(55 * time.Minute)
	}

}

// New will create a new carbonitex plugin.
// This plugin reports the server count to the carbonitex service.
func New(key string) bruxism.Plugin {
	p := &carbonitexPlugin{
		SimplePlugin: *bruxism.NewSimplePlugin("Carbonitex"),
		key:          key,
	}
	p.LoadFunc = p.carbonitexPluginLoadFunc
	return p
}
