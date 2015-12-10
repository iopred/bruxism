package main

import (
  "log"
  "os"
  "os/signal"
)

var echo = false

func registerService(service Service) error {
  messageChan, err := service.Open()

  if err != nil {
    return err
  } else {
    go func() {
      for {
        message := <-messageChan
        log.Printf("<%s> %s: %s\n", message.Channel(), message.User(), message.Message())

        if echo {
          if err := service.Send(message.Channel(), message.Message()); err != nil {
            log.Println(err)
          }
        }
      }
    }()
  }
  return nil
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
