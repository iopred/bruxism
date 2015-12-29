package bot

import "log"

func echoMessageFunc(bot *Bot, service Service, message Message) {
  if !service.IsMe(message) {
    if err := service.SendMessage(message.Channel(), message.Message()); err != nil {
      log.Println(err)
    }
  }
}

func NewEchoPlugin() *SimplePlugin {
  return NewSimplePlugin("Echo", echoMessageFunc, nil)
}

func banMessageFunc(bot *Bot, service Service, message Message) {
  if !service.IsMe(message) {
    if err := service.BanUser(message.Channel(), message.UserId(), 10); err != nil {
      log.Println(err)
    }
  }
}

func NewBanPlugin() *SimplePlugin {
  return NewSimplePlugin("Ban", banMessageFunc, nil)
}

func deleteMessageFunc(bot *Bot, service Service, message Message) {
  if !service.IsMe(message) {
    if err := service.DeleteMessage(message.MessageId()); err != nil {
      log.Println(err)
    }
  }
}
