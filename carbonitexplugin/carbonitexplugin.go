package carbonitexplugin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

func (p *carbonitexPlugin) Run(bot *bruxism.Bot, service bruxism.Service) {
	for {
		<-time.After(5 * time.Minute)

		resp, err := http.PostForm("https://www.carbonitex.net/discord/data/botdata.php", url.Values{"key": {p.key}, "servercount": {fmt.Sprintf("%d", service.ChannelCount())}})

		if err == nil {
			htmlData, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				resp.Body.Close()
			}
		}

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
