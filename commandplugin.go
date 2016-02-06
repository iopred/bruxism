package bruxism

import (
	"fmt"
	"strings"
)

const commandDelimeter = "!"

// CommandHelpFunc is the function signature for command help methods.
type CommandHelpFunc func(bot *Bot, service Service) (string, string)

// CommandMessageFunc is the function signature for bot message commands.
type CommandMessageFunc func(bot *Bot, service Service, message Message, args string, parts []string)

// NewCommandHelp creates a new Command Help function.
func NewCommandHelp(args, help string) CommandHelpFunc {
	return func(bot *Bot, service Service) (string, string) {
		return args, help
	}
}

func MatchesCommandString(service Service, commandString string, private bool, message string) bool {
	lowerMessage := strings.ToLower(strings.Trim(message, " "))
	lowerPrefix := strings.ToLower(service.CommandPrefix())

	if strings.HasPrefix(lowerMessage, lowerPrefix) {
		lowerMessage = lowerMessage[len(lowerPrefix):]
	} else if !private {
		return false
	}

	lowerMessage = strings.Trim(lowerMessage, " ")
	lowerCommand := strings.ToLower(commandString)

	return lowerMessage == lowerCommand || strings.HasPrefix(lowerMessage, lowerCommand+" ")
}

func MatchesCommand(service Service, commandString string, message Message) bool {
	// Only new messages can trigger commands.
	if message.Type() != MessageTypeCreate {
		return false
	}
	return MatchesCommandString(service, commandString, service.IsPrivate(message), message.Message())
}

func ParseCommandString(service Service, message string) (string, []string) {
	message = strings.Trim(message, " ")

	lowerMessage := strings.ToLower(message)
	lowerPrefix := strings.ToLower(service.CommandPrefix())

	if strings.HasPrefix(lowerMessage, lowerPrefix) {
		message = message[len(lowerPrefix):]
	}
	message = strings.Trim(message, " ")

	parts := strings.Split(message, " ")
	if len(parts) > 1 {
		rest := parts[1:]
		return strings.Join(rest, " "), parts[1:]
	}
	return "", []string{}
}

func ParseCommand(service Service, message Message) (string, []string) {
	return ParseCommandString(service, message.Message())
}

func CommandHelp(service Service, command, arguments, help string) []string {
	ticks := ""
	if service.Name() == DiscordServiceName {
		ticks = "`"
	}

	if arguments != "" {
		return []string{fmt.Sprintf("%s%s%s %s%s - %s", ticks, service.CommandPrefix(), command, arguments, ticks, help)}
	}
	return []string{fmt.Sprintf("%s%s%s%s - %s", ticks, service.CommandPrefix(), command, ticks, help)}
}

type command struct {
	message CommandMessageFunc
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
func (p *CommandPlugin) Help(bot *Bot, service Service, detailed bool) []string {
	if detailed {
		return nil
	}
	help := []string{}
	for commandString, command := range p.commands {
		if command.help != nil {
			arguments, h := command.help(bot, service)
			help = append(help, CommandHelp(service, commandString, arguments, h)...)
		}
	}
	return help
}

// Message handler.
// Iterates over the registered commands and executes them if the message matches.
func (p *CommandPlugin) Message(bot *Bot, service Service, message Message) {
	defer messageRecover()
	if !service.IsMe(message) {
		for commandString, command := range p.commands {
			if MatchesCommand(service, commandString, message) {
				args, parts := ParseCommand(service, message)
				command.message(bot, service, message, args, parts)
				return
			}
		}
	}
}

// AddCommand adds a command.
func (p *CommandPlugin) AddCommand(commandString string, message CommandMessageFunc, help CommandHelpFunc) {
	p.commands[commandString] = &command{
		message: message,
		help:    help,
	}
}

// NewCommandPlugin will create a new command plugin.
func NewCommandPlugin() *CommandPlugin {
	return &CommandPlugin{make(map[string]*command)}
}
