package bot

import (
  "encoding/json"
  "log"
  "strings"
)

type SlowModePlugin struct {
  Enabled map[string]bool
}

func (p *SlowModePlugin) Name() string {
  return "SlowMode"
}

func (p *SlowModePlugin) Help() string {
  return "!slowmode [on|off] - Turn slow mode on or off."
}

func (p *SlowModePlugin) Register(bot *Bot, service Service, data []byte) error {
  if data != nil {
    if err := json.Unmarshal(data, p); err != nil {
      log.Println("Error loading data", err)
    }
  }

  messageChannel := bot.NewMessageChannel(service)
  go func() {
    for {
      message := <-messageChannel

      if service.IsMe(message) {
        continue
      }

      messageChannel := message.Channel()
      messageMessage := strings.Replace(message.Message(), " ", "", -1)
      messageModerator := message.IsModerator()

      if strings.HasPrefix(messageMessage, "!slowmode") {
        enabled := p.Enabled[messageChannel]
        if messageMessage == "!slowmodeon" {
          enabled = true
        } else if messageMessage == "!slowmodeoff" {
          enabled = false
        }

        changed := enabled != p.Enabled[messageChannel]
        if changed {
          p.Enabled[messageChannel] = enabled
        }

        if enabled {
          if changed {
            service.SendMessage(messageChannel, "Slow mode is now on (You will be temporarily banned for 30 seconds when you chat).")
          } else {
            service.SendMessage(messageChannel, "Slow mode is on (You will be temporarily banned for 30 seconds when you chat).")
          }
        } else {
          if changed {
            service.SendMessage(messageChannel, "Slow mode is now off.")
          } else {
            service.SendMessage(messageChannel, "Slow mode is off.")
          }
        }
      } else if p.Enabled[messageChannel] && !messageModerator {
        if err := service.BanUser(messageChannel, message.UserId(), 30); err != nil {
          log.Println(err)
        }
      }
    }
  }()
  return nil
}

func (p *SlowModePlugin) Save() []byte {
  if data, err := json.Marshal(p); err == nil {
    return data
  }
  return nil
}

func NewSlowModePlugin() *SlowModePlugin {
  return &SlowModePlugin{
    Enabled: make(map[string]bool),
  }
}
