package main

import (
  "log"
  "os"
  "os/signal"
)

// Flip these to true to echo, delete or timeout every message.
var echo = false
var del = false
var ban = false

func registerService(service Service) error {
  go func() {
    messageChan := service.MessageChannel()
    for {
      message := <-messageChan
      log.Printf("<%s> %s: %s\n", message.Channel(), message.UserName(), message.Message())

      if echo {
        if err := service.SendMessage(message.Channel(), message.Message()); err != nil {
          log.Println(err)
        }
      }
      if del {
        if err := service.DeleteMessage(message.MessageId()); err != nil {
          log.Println(err)
        }
      }
      if ban {
        if err := service.BanUser(message.Channel(), message.UserId(), 10); err != nil {
          log.Println(err)
        }
      }
    }
  }()

  return service.Open()
}

func main() {
  if err := registerService(NewYouTube()); err != nil {
    log.Println(err)
  }

  if err := registerService(NewDiscord()); err != nil {
    log.Println(err)
  }

  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt, os.Kill)
  <-c
}
