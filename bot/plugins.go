package bot

import "log"

type EchoPlugin struct{}

func (p *EchoPlugin) Name() string {
  return "Echo"
}

func (p *EchoPlugin) Help() string {
  return ""
}

func (p *EchoPlugin) Register(bot *Bot, service Service, data []byte) error {
  messageChannel := bot.NewMessageChannel(service)
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

func (p *EchoPlugin) Save() []byte {
  return nil
}

func NewEchoPlugin() *EchoPlugin {
  return &EchoPlugin{}
}

type BanPlugin struct{}

func (p *BanPlugin) Name() string {
  return "Ban"
}

func (p *BanPlugin) Help() string {
  return ""
}

func (p *BanPlugin) Register(bot *Bot, service Service, data []byte) error {
  messageChannel := bot.NewMessageChannel(service)
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

func (p *BanPlugin) Save() []byte {
  return nil
}

func NewBanPlugin() *BanPlugin {
  return &BanPlugin{}
}

type DeletePlugin struct{}

func (p *DeletePlugin) Name() string {
  return "Delete"
}

func (p *DeletePlugin) Help() string {
  return ""
}

func (p *DeletePlugin) Register(bot *Bot, service Service, data []byte) error {
  messageChannel := bot.NewMessageChannel(service)
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

func (p *DeletePlugin) Save() []byte {
  return nil
}

func NewDeletePlugin() *DeletePlugin {
  return &DeletePlugin{}
}
