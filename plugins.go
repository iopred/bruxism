package main

import "log"

type EchoPlugin struct{}

func (e *EchoPlugin) Name() string {
	return "Echo"
}

func (e *EchoPlugin) Register(bot *Bot, service Service, data []byte) error {
	messageChannel := bot.MessageChannel(service)
	go func() {
		for {
			message := <-messageChannel
			if !service.IsMe(message) {
				if err := service.SendMessage(message.Channel(), message.Message()); err != nil {
					log.Println(err)
				}
			}
		}
	}()
	return nil
}

func (e *EchoPlugin) Save() []byte {
	return nil
}

func NewEchoPlugin() *EchoPlugin {
	return &EchoPlugin{}
}

type BanPlugin struct{}

func (e *BanPlugin) Name() string {
	return "Echo"
}

func (e *BanPlugin) Register(bot *Bot, service Service, data []byte) error {
	messageChannel := bot.MessageChannel(service)
	go func() {
		for {
			message := <-messageChannel
			if !service.IsMe(message) {
				if err := service.BanUser(message.Channel(), message.UserId(), 10); err != nil {
					log.Println(err)
				}
			}
		}
	}()
	return nil
}

func (e *BanPlugin) Save() []byte {
	return nil
}

func NewBanPlugin() *BanPlugin {
	return &BanPlugin{}
}

type DeletePlugin struct{}

func (e *DeletePlugin) Name() string {
	return "Echo"
}

func (e *DeletePlugin) Register(bot *Bot, service Service, data []byte) error {
	messageChannel := bot.MessageChannel(service)
	go func() {
		for {
			message := <-messageChannel
			if !service.IsMe(message) {
				if err := service.DeleteMessage(message.MessageId()); err != nil {
					log.Println(err)
				}
			}
		}
	}()
	return nil
}

func (e *DeletePlugin) Save() []byte {
	return nil
}

func NewDeletePlugin() *DeletePlugin {
	return &DeletePlugin{}
}
