package main

import (
	"io/ioutil"
	"log"
	"os"
)

type BotService struct {
	service         Service
	plugins         map[string]Plugin
	messageChannels []chan Message
}

type Bot struct {
	services map[string]*BotService
}

func NewBot() *Bot {
	return &Bot{
		services: make(map[string]*BotService, 0),
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
	b.services[serviceName] = &BotService{
		service:         service,
		plugins:         make(map[string]Plugin, 0),
		messageChannels: make([]chan Message, 0),
	}
}

func (b *Bot) RegisterPlugin(service Service, plugin Plugin) {
	b.services[service.Name()].plugins[plugin.Name()] = plugin
	plugin.Register(b, service, b.getData(service, plugin))
}

func (b *Bot) MessageChannel(service Service) <-chan Message {
	messageChannel := make(chan Message, 200)
	serviceName := service.Name()
	b.services[serviceName].messageChannels = append(b.services[serviceName].messageChannels, messageChannel)
	return messageChannel
}

func (b *Bot) listen(service Service, serviceMessageChannel <-chan Message) {
	serviceName := service.Name()
	for {
		message := <-serviceMessageChannel
		log.Printf("<%s> %s: %s\n", message.Channel(), message.UserName(), message.Message())
		messageChannels := b.services[serviceName].messageChannels
		for _, messageChannel := range messageChannels {
			messageChannel <- message
		}
	}
}

func (b *Bot) Open() {
	for _, service := range b.services {
		if messageChan, err := service.service.Open(); err == nil {
			go b.listen(service.service, messageChan)
		} else {
			log.Println("Error creating service %v: %v", service.service.Name(), err)
		}

	}
}

func (b *Bot) Save() {
	for _, service := range b.services {
		serviceName := service.service.Name()
		if err := os.Mkdir(serviceName, os.ModePerm); err != nil {
			if !os.IsExist(err) {
				log.Println("Error creating service directory.")
			}
		}
		for _, plugin := range service.plugins {
			data := plugin.Save()
			if data != nil {
				if err := ioutil.WriteFile(serviceName+"/"+plugin.Name(), data, os.ModePerm); err != nil {
					log.Printf("Error saving plugin %v %v. %v", serviceName, plugin.Name(), err)
				}
			}
		}
	}
}
