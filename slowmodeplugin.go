package main

import (
  "encoding/json"
  "fmt"
  "log"
  "strings"
)

type SlowModePlugin struct {
  Enabled map[string]bool
}

func (e *SlowModePlugin) Name() string {
  return "SlowMode"
}

func (e *SlowModePlugin) Register(bot *Bot, service Service, data []byte) error {
  if data != nil {
    if err := json.Unmarshal(data, e); err != nil {
      log.Println("Error loading data", err)
    }
  }

  messageChannel := bot.NewMessageChannel(service)
  go func() {
    for {
      message := <-messageChannel
      messageMessage := strings.Replace(message.Message(), " ", "", -1)
      messageChannel := message.Channel()
      messageModerator := message.IsModerator()

      if strings.HasPrefix(messageMessage, "!slowmode") {
        enabled := e.Enabled[messageChannel]
        if messageMessage == "!slowmodeon" {
          enabled = true
        } else if messageMessage == "!slowmodeoff" {
          enabled = false
        }

        changed := enabled != e.Enabled[messageChannel]
        if changed {
          e.Enabled[messageChannel] = enabled
        }

        if enabled {
          if changed {
            service.SendMessage(messageChannel, fmt.Sprintf("Slow mode is now on (You will be temporarily banned for 30 seconds when you chat).", label, hint))
          } else {
            service.SendMessage(messageChannel, fmt.Sprintf("Slow mode is on (You will be temporarily banned for 30 seconds when you chat).", label, hint))
          }
        } else {
          if changed {
            service.SendMessage(messageChannel, fmt.Sprintf("Slow mode is now off.", label, hint))
          } else {
            service.SendMessage(messageChannel, fmt.Sprintf("Slow mode is off.", label, hint))
          }
        }
      } else if e.Enabled[messageChannel] && !messageModerator {
        if err := service.BanUser(messageChannel, message.UserId(), 30); err != nil {
          log.Println(err)
        }
      }
    }
  }()
  return nil
}

func (e *SlowModePlugin) Save() []byte {
  if data, err := json.Marshal(e); err == nil {
    return data
  }
  return nil
}

func NewSlowModePlugin() *SlowModePlugin {
  return &SlowModePlugin{
    Enabled: make(map[string]bool),
  }
}
