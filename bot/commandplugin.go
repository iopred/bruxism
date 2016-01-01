package bot

import (
	"fmt"
	"strings"
)

const commandDelimeter = "!"

// CommandHelpFunc is the function signature for command help methods.
type CommandHelpFunc func(bot *Bot, service Service) (string, string)

// CommandMessageFunc is the function signature for bot message commands.
type CommandMessageFunc func(joined string, parts []string) string

// NewCommandMessageFunc will returns a MessageFunc that can be used by SimplePlugin, given a CommandMessageFunc.
func NewCommandMessageFunc(commandMessageFunc CommandMessageFunc) MessageFunc {
	return func(bot *Bot, service Service, message Message) {
		if response := commandMessageFunc(parseCommand(service, message)); response != "" {
			service.SendMessage(message.Channel(), response)
		}
	}
}

func matchesCommand(service Service, commandString string, message Message) bool {
	return strings.HasPrefix(strings.ToLower(message.Message()), strings.ToLower(service.CommandPrefix()+commandString))
}

func parseCommand(service Service, message Message) (string, []string) {
	m := message.Message()[len(service.CommandPrefix()):]

	parts := strings.Split(m, " ")
	if len(parts) > 1 {
		rest := parts[1:]
		return strings.Join(rest, " "), parts[1:]
	}
	return "", []string{}
}

func commandHelp(service Service, command, arguments, help string) []string {
	if arguments != "" {
		return []string{fmt.Sprintf("%s%s %s - %s", service.CommandPrefix(), command, arguments, help)}
	}
	return []string{fmt.Sprintf("%s%s - %s", service.CommandPrefix(), command, help)}
}

type command struct {
	message MessageFunc
	help    CommandHelpFunc
}

// CommandPlugin is a plugin that can have commands registered and will handle messages matching that command by calling functions.
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
	help := []string{}
	for commandString, command := range p.commands {
		if command.help != nil {
			arguments, h := command.help(bot, service)
			help = append(help, commandHelp(service, commandString, arguments, h)...)
		}
	}
	return help
}

// Message handler.
// Iterates over the registered commands and executes them if the message matches.
func (p *CommandPlugin) Message(bot *Bot, service Service, message Message) {
	if !service.IsMe(message) {
		for commandString, command := range p.commands {
			if matchesCommand(service, commandString, message) {
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

// NewCommandPlugin will create a new command plugin.
func NewCommandPlugin() *CommandPlugin {
	return &CommandPlugin{make(map[string]*command)}
}
