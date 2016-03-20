package mtgplugin

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/iopred/bruxism"
	"github.com/renstrom/fuzzysearch/fuzzy"
)

// MTGSet is a struct containing all the information on a set of MTG.
type MTGSet struct {
	Cards []*MTGCard `json:"cards"`
}

// MTGSet is a struct containing all the information on a card from MTG.
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

// MTGCardMap is a map from card name to card.
var MTGCardMap = map[string]*MTGCard{}

// MTGCardNames is an array of card names
var MTGCardNames []string

// MTGTextReplacer is a replacer for styling text content of cards.
var MTGTextReplacer = strings.NewReplacer("(", "*(", ")", ")*")

// MTGCostReplacer is a replacer for styling text content of card costs.
var MTGCostReplacer = strings.NewReplacer("{", "", "}", "")

// MTGRestReplacer is a replacer for styling card info.
var MTGRestReplacer = strings.NewReplacer("*", "\\*")

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

// MTGHelp is the help for the MTG plugin.
var MTGHelp = bruxism.NewCommandHelp("<cardname>", "Returns information about a Magic: The Gathering card.")
