package numbertriviaplugin

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/iopred/bruxism"
)

func numberTrivia(bot *bruxism.Bot, num int, notfound bool) (string, error) {
	notfoundString := ""
	if notfound {
		notfoundString = "?notfound=floor"
	}
	r, err := http.NewRequest("GET", fmt.Sprintf("https://numbersapi.p.mashape.com/%d/trivia%s", num, notfoundString), nil)
	if err != nil {
		return "", err
	}

	r.Header.Set("X-Mashape-Authorization", bot.MashableKey)
	r.Header.Set("Accept", "text/plain")

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New(string(body))
	}

	return string(body), nil
}

// NumberTriviaCommand is a command for getting number trivial.
func NumberTriviaCommand(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, command string, parts []string) {
	notfound := true
	num := rand.Intn(1000)
	if len(parts) == 1 {
		if i, err := strconv.Atoi(parts[0]); err == nil {
			num = i
			notfound = false
		}
	}

	service.Typing(message.Channel())

	str, err := numberTrivia(bot, num, notfound)
	if err != nil {
		service.SendMessage(message.Channel(), "There was an error requesting trivia, sorry!")
		return
	}

	service.SendMessage(message.Channel(), str)
}

var NumberTriviaHelp = bruxism.NewCommandHelp("[<number>]", "Returns trivia for a random number or a specified number if provided.")
