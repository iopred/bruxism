package bruxism

import (
	"fmt"
	"strings"
)

const commandDelimeter = "!"

// CommandHelpFunc is the function signature for command help methods.
type CommandHelpFunc func(bot *Bot, service Service, message Message) (string, string)

// CommandMessageFunc is the function signature for bot message commands.
type CommandMessageFunc func(bot *Bot, service Service, message Message, args string, parts []string)

// NewCommandHelp creates a new Command Help function.
func NewCommandHelp(args, help string) CommandHelpFunc {
	return func(bot *Bot, service Service, message Message) (string, string) {
		return args, help
	}
}

// MatchesCommandString returns true if a message matches a command.
// Commands will be matched ignoring case with a prefix if they are not private messages.
func MatchesCommandString(service Service, commandString string, private bool, message string) bool {
	lowerMessage := strings.ToLower(strings.TrimSpace(message))
	lowerPrefix := strings.ToLower(service.CommandPrefix())

	if strings.HasPrefix(lowerMessage, lowerPrefix) {
		lowerMessage = lowerMessage[len(lowerPrefix):]
	} else if !private {
		return false
	}

	lowerMessage = strings.TrimSpace(lowerMessage)
	lowerCommand := strings.ToLower(commandString)

	return lowerMessage == lowerCommand || strings.HasPrefix(lowerMessage, lowerCommand+" ")
}

// MatchesCommand returns true if a message matches a command.
func MatchesCommand(service Service, commandString string, message Message) bool {
	// Deleted messages can't trigger commands.
	if message.Type() == MessageTypeDelete {
		return false
	}
	return MatchesCommandString(service, commandString, service.IsPrivate(message), message.Message())
}

// ParseCommandString will strip all prefixes from a message string, and return that string, and a space separated tokenized version of that string.
func ParseCommandString(service Service, message string) (string, []string) {
	message = strings.TrimSpace(message)

	lowerMessage := strings.ToLower(message)
	lowerPrefix := strings.ToLower(service.CommandPrefix())

	if strings.HasPrefix(lowerMessage, lowerPrefix) {
		message = message[len(lowerPrefix):]
	}
	rest := strings.Fields(message)

	if len(rest) > 1 {
		rest = rest[1:]
		return strings.Join(rest, " "), rest
	}
	return "", []string{}
}

// ParseCommand parses a message.
func ParseCommand(service Service, message Message) (string, []string) {
	return ParseCommandString(service, message.Message())
}

// CommandHelp is a helper message that creates help text for a command.
// eg. CommandHelp(service, "foo", "<bar>", "Foo bar baz") will return:
//     !foo <bar> - Foo bar baz
// The string is automatatically styled in Discord.
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
func (p *CommandPlugin) Help(bot *Bot, service Service, message Message, detailed bool) []string {
	if detailed {
		return nil
	}
	help := []string{}
	for commandString, command := range p.commands {
		if command.help != nil {
			arguments, h := command.help(bot, service, message)
			help = append(help, CommandHelp(service, commandString, arguments, h)...)
		}
	}
	return help
}

// Message handler.
// Iterates over the registered commands and executes them if the message matches.
func (p *CommandPlugin) Message(bot *Bot, service Service, message Message) {
	defer MessageRecover()
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

// Stats will return the stats for a plugin.
func (p *CommandPlugin) Stats(bot *Bot, service Service, message Message) []string {
	return nil
}

// NewCommandPlugin will create a new command plugin.
func NewCommandPlugin() *CommandPlugin {
	return &CommandPlugin{make(map[string]*command)}
}
