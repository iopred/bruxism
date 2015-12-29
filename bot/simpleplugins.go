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
  p := NewSimplePlugin("Echo")
  p.message = echoMessageFunc
  return p
}

func banMessageFunc(bot *Bot, service Service, message Message) {
  if !service.IsMe(message) {
    if err := service.BanUser(message.Channel(), message.UserId(), 10); err != nil {
      log.Println(err)
    }
  }
}

func NewBanPlugin() *SimplePlugin {
  p := NewSimplePlugin("Ban")
  p.message = banMessageFunc
  return p
}

func deleteMessageFunc(bot *Bot, service Service, message Message) {
  if !service.IsMe(message) {
    if err := service.DeleteMessage(message.MessageId()); err != nil {
      log.Println(err)
    }
  }
}

func NewDeletePlugin() *SimplePlugin {
  p := NewSimplePlugin("Delete")
  p.message = deleteMessageFunc
  return p
}
