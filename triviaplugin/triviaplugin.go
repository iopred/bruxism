package triviaplugin

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/iopred/bruxism"
)

type question struct {
	Question string
	Answer   string
}

type questions []*question

func (q questions) Question() *question {
	n := rand.Intn(len(q))
	return q[n]
}

type trivia struct {
	All    questions
	Themes map[string]questions
}

func (t *trivia) Question(theme string) *question {
	q := t.Themes[theme]
	if q == nil {
		q = t.All
	}
	return q.Question()
}

var triviaQuestions *trivia = &trivia{
	All:    questions{},
	Themes: map[string]questions{},
}

type triviaScore struct {
	Name  string
	Score int
}

type triviaChannel struct {
	sync.RWMutex

	Channel    string
	Active     bool
	Theme      string
	Unanswered int

	Answer string
	Hint   string
	Asked  time.Time

	Scores map[string]*triviaScore

	hintChan chan bool
}

func (t *triviaChannel) Start(bot *bruxism.Bot, service bruxism.Service, theme string) {
	t.Lock()
	defer t.Unlock()

	if t.Active {
		return
	}

	service.SendMessage(t.Channel, "Trivia started.")

	t.Active = true
	t.Theme = theme
	t.Unanswered = 0

	go t.question(bot, service)

	return
}

func (t *triviaChannel) Stop(bot *bruxism.Bot, service bruxism.Service) {
	t.Lock()
	defer t.Unlock()

	if !t.Active {
		return
	}

	t.Active = false
	if t.hintChan != nil {
		close(t.hintChan)
		t.hintChan = nil
	}

	service.SendMessage(t.Channel, "Trivia stopped.")

	return
}

func (t *triviaChannel) Message(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	if strings.ToLower(message.Message()) == t.Answer {
		t.Lock()
		defer t.Unlock()

		if !t.Active {
			return
		}

		ts := t.Scores[message.UserID()]
		if ts == nil {
			ts = &triviaScore{}
			t.Scores[message.UserID()] = ts
		}
		ts.Name = message.UserName()
		ts.Score++

		service.SendMessage(message.Channel(), fmt.Sprintf("%s got it! The answer was %s.", message.UserName(), t.Answer))
		service.SendMessage(message.Channel(), fmt.Sprintf("%s answered in %d seconds and their score is now %d.", message.UserName(), int(time.Now().Sub(t.Asked).Seconds()), ts.Score))

		t.Unanswered = 0

		if t.hintChan != nil {
			close(t.hintChan)
			t.hintChan = nil
		}
	}
}

func (t *triviaChannel) question(bot *bruxism.Bot, service bruxism.Service) {
	t.Lock()
	question := triviaQuestions.Question(t.Theme)

	hintChan := make(chan bool)
	t.hintChan = hintChan
	t.Answer = strings.ToLower(question.Answer)
	t.Asked = time.Now()
	t.Unlock()

	service.SendMessage(t.Channel, question.Question)

	answer := strings.Split(question.Answer, "")
	hint := make([]string, len(answer))

	chars := 0
	for i, s := range answer {
		if s == " " {
			hint[i] = " "
		} else {
			chars++
			hint[i] = "-"
		}
	}

	hints := 3
	if hints > chars {
		hints = chars
	}

	hintTime := (time.Minute * 4) / time.Duration(hints+1)
	hintCount := chars / (hints + 1)

	func() {
		for {
			select {
			case <-hintChan:
				return
			case <-time.After(hintTime):
				if hints == 0 {
					service.SendMessage(t.Channel, fmt.Sprintf("Time's up! The answer was: %s.", question.Answer))
					t.Lock()
					t.Unanswered++
					if t.Unanswered > 4 {
						service.SendMessage(t.Channel, "Too many unanswered questions. Trivia stopped.")
						t.Active = false
					}
					t.Unlock()
					return
				} else {
					hints--

					service.SendMessage(t.Channel, "Hint: "+strings.Join(hint, ""))

					for i := 0; i < hintCount; i++ {
						for {
							r := rand.Intn(len(hint))
							if hint[r] == "-" {
								hint[r] = answer[r]
								break
							}
						}
					}
				}
			}
		}
	}()

	t.RLock()
	defer t.RUnlock()
	if t.Active {
		go t.question(bot, service)
	}
}

type triviaPlugin struct {
	sync.RWMutex
	// Map from ChannelID -> triviaChannel
	Channels map[string]*triviaChannel
}

// Name returns the name of the plugin.
func (p *triviaPlugin) Name() string {
	return "Trivia"
}

// Load will load plugin state from a byte array.
func (p *triviaPlugin) Load(bot *bruxism.Bot, service bruxism.Service, data []byte) error {
	if data != nil {
		if err := json.Unmarshal(data, p); err != nil {
			log.Println("Error loading data", err)
		}
	}

	for _, t := range p.Channels {
		if t.Active {
			go t.question(bot, service)
		}
	}

	return nil
}

// Save will save plugin state to a byte array.
func (p *triviaPlugin) Save() ([]byte, error) {
	return json.Marshal(p)
}

// Help returns a list of help strings that are printed when the user requests them.
func (p *triviaPlugin) Help(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, detailed bool) []string {
	if service.IsPrivate(message) || !(service.IsModerator(message) || service.IsBotOwner(message)) {
		return nil
	}

	return []string{
		bruxism.CommandHelp(service, "trivia", "<start|stop> [theme]", "Starts or stops trivia with an optional theme.")[0],
		bruxism.CommandHelp(service, "trivia", "<score>", "Returns your current trivia score.")[0],
	}
}

// Message handler.
func (p *triviaPlugin) Message(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	defer bruxism.MessageRecover()
	if !service.IsMe(message) && !service.IsPrivate(message) {
		messageChannel := message.Channel()

		isCommand := bruxism.MatchesCommand(service, "trivia", message)

		if isCommand && (service.IsModerator(message) || service.IsBotOwner(message)) {
			p.Lock()
			tc := p.Channels[messageChannel]
			if tc == nil {
				tc = &triviaChannel{
					Channel: messageChannel,
					Scores:  map[string]*triviaScore{},
				}
				p.Channels[messageChannel] = tc
			}
			p.Unlock()

			_, parts := bruxism.ParseCommand(service, message)

			if len(parts) == 0 {
				return
			}

			switch parts[0] {
			case "start":
				theme := ""
				if len(parts) >= 2 {
					theme = parts[1]
				}
				tc.Start(bot, service, theme)
			case "stop":
				tc.Stop(bot, service)
			}

		} else {
			if isCommand {
				_, parts := bruxism.ParseCommand(service, message)
				if len(parts) == 0 {
					return
				}
				if parts[0] == "score" {
					p.RLock()
					tc := p.Channels[messageChannel]

					if tc != nil {
						ts := tc.Scores[message.UserID()]
						if ts != nil {
							service.SendMessage(message.Channel(), fmt.Sprintf("%s's score is %d.", message.UserName(), ts.Score))
						} else {
							service.SendMessage(message.Channel(), fmt.Sprintf("%s's score is 0.", message.UserName()))
						}
					}

					p.RUnlock()
				}

				return
			}

			p.RLock()
			tc := p.Channels[messageChannel]
			p.RUnlock()
			if tc != nil {
				tc.Message(bot, service, message)
			}
		}
	}
}

// Stats will return the stats for a plugin.
func (p *triviaPlugin) Stats(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) []string {
	return nil
}

// New will create a new trivia plugin.
func New() bruxism.Plugin {
	return &triviaPlugin{
		Channels: map[string]*triviaChannel{},
	}
}
