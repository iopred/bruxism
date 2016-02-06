package mtgplugin

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/iopred/bruxism"
	"github.com/scaleway/scaleway-cli/vendor/github.com/renstrom/fuzzysearch/fuzzy"
)

type MTGSet struct {
	Cards []*MTGCard `json:"cards"`
}

type MTGCard struct {
	Name      string  `json:"name"`
	ManaCost  string  `json:"manaCost"`
	Type      string  `json:"type"`
	Text      string  `json:"text"`
	ID        *int    `json:"multiverseid"`
	Power     *string `json:"power"`
	Toughness *string `json:"toughness"`
	Loyalty   *int    `json:"loyalty"`
}

var MTGCardMap map[string]*MTGCard = map[string]*MTGCard{}
var MTGCardNames []string

var MTGTextReplacer *strings.Replacer = strings.NewReplacer("(", "*(", ")", ")*")
var MTGCostReplacer *strings.Replacer = strings.NewReplacer("{", "", "}", "")
var MTGRestReplacer *strings.Replacer = strings.NewReplacer("*", "\\*")

func init() {
	file, err := os.Open("mtg/AllSets-x.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	MTGSets := map[string]*MTGSet{}

	d := json.NewDecoder(bufio.NewReader(file))
	err = d.Decode(&MTGSets)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, s := range MTGSets {
		for _, c := range s.Cards {
			c.ManaCost = MTGCostReplacer.Replace(c.ManaCost)
			c.Text = MTGCostReplacer.Replace(c.Text)
			MTGCardMap[c.Name] = c
			MTGCardNames = append(MTGCardNames, c.Name)
		}
	}
}

// MTGCommand is a command for getting information about MTG cards..
func MTGCommand(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, command string, parts []string) {
	cardNames := fuzzy.RankFindFold(command, MTGCardNames)
	if len(cardNames) == 0 {
		service.SendMessage(message.Channel(), "Could not find a card with that name, sorry.")
		return
	}

	sort.Sort(cardNames)

	card := MTGCardMap[cardNames[0].Target]

	rest := ""
	if card.Text != "" {
		rest += "\n"
	}
	if card.Power != nil {
		rest += MTGRestReplacer.Replace(fmt.Sprintf("%s/%s", *card.Power, *card.Toughness))
	}
	if card.Loyalty != nil {
		rest += MTGRestReplacer.Replace(fmt.Sprintf("%d", *card.Loyalty))
	}
	if card.ID != nil {
		if rest != "" && rest != "\n" {
			rest += "\n"
		}
		rest += fmt.Sprintf("(http://gatherer.wizards.com/Handlers/Image.ashx?multiverseid=%d&type=card)", *card.ID)
	}

	if service.Name() == bruxism.DiscordServiceName {
		service.SendMessage(message.Channel(), fmt.Sprintf("**%s** %s\n*%s*\n%s%s", card.Name, card.ManaCost, card.Type, MTGTextReplacer.Replace(card.Text), rest))
	} else {
		service.SendMessage(message.Channel(), strings.Replace(fmt.Sprintf("%s. %s. %s. %s%s", card.Name, card.Type, card.ManaCost, card.Text, rest), "\n", " ", -1))
	}
}
