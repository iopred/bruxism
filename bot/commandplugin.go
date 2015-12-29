package bot

import (
	"fmt"
	"strings"
)

const commandDelimeter = "!"

// The function signature for command help methods.
type CommandHelpFunc func(bot *Bot, service Service) (string, string)

// The function signature for bot message commands.
type CommandMessageFunc func(joined string, parts []string) string

// Given a CommandMessageFunc, returns a MessageFunc that can be used by SimplePlugin.
func NewCommandMessageFunc(commandMessageFunc CommandMessageFunc) MessageFunc {
	return func(bot *Bot, service Service, message Message) {
		if response := commandMessageFunc(parseCommand(message)); response != "" {
			service.SendMessage(message.Channel(), response)
		}
	}
}

func matchesCommand(commandString string, message Message) bool {
	return strings.HasPrefix(message.Message(), commandDelimeter+commandString)
}

func parseCommand(message Message) (string, []string) {
	parts := strings.Split(message.Message(), " ")
	if len(parts) > 1 {
		rest := parts[1:]
		return strings.Join(rest, " "), parts[1:]
	}
	return "", []string{}
}

func commandHelp(command, arguments, help string) []string {
	if arguments != "" {
		return []string{fmt.Sprintf("%s%s %s - %s", commandDelimeter, command, arguments, help)}
	}
	return []string{fmt.Sprintf("%s%s - %s", commandDelimeter, command, help)}
}

type command struct {
	message MessageFunc
	help    CommandHelpFunc
}

// Command plugin is a plugin that can have commands registered and will handle messages matching that command by calling functions.
type CommandPlugin struct {
	commands map[string]*command
}

// Name returns the name of the plugin.
func (p *CommandPlugin) Name() string {
	return "Command"
}

// Load will load plugin state from a byte array.
func (p *CommandPlugin) Load(bot *Bot, service Service, data []byte) error {
	// TODO: Add a generic data store backed by json.
	return nil
}

// Save will save plugin state to a byte array.
func (p *CommandPlugin) Save() ([]byte, error) {
	// TODO: Add a generic data store backed by json.
	return nil, nil
}

// Help returns a list of help strings that are printed when the user requests them.
func (p *CommandPlugin) Help(bot *Bot, service Service) []string {
	help := make([]string, 0)
	for commandString, command := range p.commands {
		if command.help != nil {
			arguments, h := command.help(bot, service)
			help = append(help, commandHelp(commandString, arguments, h)...)
		}
	}
	return help
}

// Message handler.
// Iterates over the registered commands and executes them if the message matches.
func (p *CommandPlugin) Message(bot *Bot, service Service, message Message) {
	if !service.IsMe(message) {
		for commandString, command := range p.commands {
			if matchesCommand(commandString, message) {
				command.message(bot, service, message)
			}
		}
	}
}

// AddCommand adds a command.
func (p *CommandPlugin) AddCommand(commandString string, message MessageFunc, help CommandHelpFunc) {
	p.commands[commandString] = &command{
		message: message,
		help:    help,
	}
}

// AddSimpleCommand adds a simple command.
func (p *CommandPlugin) AddSimpleCommand(commandString string, message CommandMessageFunc, arguments, help string) {
	p.AddCommand(commandString, NewCommandMessageFunc(message), func(bot *Bot, service Service) (string, string) {
		return arguments, help
	})
}

// Create will create a new command plugin.
func NewCommandPlugin() *CommandPlugin {
	return &CommandPlugin{make(map[string]*command)}
}
