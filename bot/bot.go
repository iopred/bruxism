package bot

import (
	"io/ioutil"
	"log"
	"os"
)

type BotService struct {
	Service
	Plugins         map[string]Plugin
	messageChannels []chan Message
}

type Bot struct {
	Services map[string]*BotService
}

func NewBot() *Bot {
	return &Bot{
		Services: make(map[string]*BotService, 0),
	}
}

func (b *Bot) getData(service Service, plugin Plugin) []byte {
	if b, err := ioutil.ReadFile(service.Name() + "/" + plugin.Name()); err == nil {
		return b
	}
	return nil
}

func (b *Bot) RegisterService(service Service) {
	serviceName := service.Name()
	b.Services[serviceName] = &BotService{
		Service: service,
		Plugins: make(map[string]Plugin, 0),
	}
}

func (b *Bot) RegisterPlugin(service Service, plugin Plugin) {
	b.Services[service.Name()].Plugins[plugin.Name()] = plugin
	plugin.Load(b, service, b.getData(service, plugin))
}

func (b *Bot) listen(service Service, messageChan <-chan Message) {
	serviceName := service.Name()
	for {
		message := <-messageChan
		log.Printf("<%s> %s: %s\n", message.Channel(), message.UserName(), message.Message())
		plugins := b.Services[serviceName].Plugins
		for _, plugin := range plugins {
			go plugin.Message(b, service, message)
		}
	}
}

func (b *Bot) Open() {
	for _, service := range b.Services {
		if messageChan, err := service.Open(); err == nil {
			go b.listen(service, messageChan)
		} else {
			log.Printf("Error creating service %v: %v\n", service.Name(), err)
		}

	}
}

func (b *Bot) Save() {
	for _, service := range b.Services {
		serviceName := service.Name()
		if err := os.Mkdir(serviceName, os.ModePerm); err != nil {
			if !os.IsExist(err) {
				log.Println("Error creating service directory.")
			}
		}
		for _, plugin := range service.Plugins {
			if data, err := plugin.Save(); err != nil {
				log.Printf("Error saving plugin %v %v. %v", serviceName, plugin.Name(), err)
			} else if data != nil {
				if err := ioutil.WriteFile(serviceName+"/"+plugin.Name(), data, os.ModePerm); err != nil {
					log.Printf("Error saving plugin %v %v. %v", serviceName, plugin.Name(), err)
				}
			}
		}
	}
}
