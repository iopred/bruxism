package bot

import "strings"

const COMMAND_DELIMITER = "!"

type CommandMessageFunc func(joined string, parts []string) string

func NewCommandMessageFunc(commandMessageFunc CommandMessageFunc) MessageFunc {
	return func(bot *Bot, service Service, message Message) {
		if response := commandMessageFunc(parseCommand(message)); response != "" {
			service.SendMessage(message.Channel(), response)
		}
	}
}

func matchesCommand(commandString string, message Message) bool {
	return strings.HasPrefix(message.Message(), COMMAND_DELIMITER+commandString)
}

func parseCommand(message Message) (string, []string) {
	parts := strings.Split(message.Message(), " ")
	if len(parts) > 1 {
		rest := parts[1:]
		return strings.Join(rest, " "), parts[1:]
	}
	return "", []string{}
}

type Command struct {
	message MessageFunc
	help    HelpFunc
}

type CommandPlugin struct {
	commands map[string]*Command
}

func (p *CommandPlugin) Name() string {
	return "Command"
}

func (p *CommandPlugin) Load(bot *Bot, service Service, data []byte) error {
	// TODO: Add a generic data store backed by json.
	return nil
}

func (p *CommandPlugin) Save() ([]byte, error) {
	// TODO: Add a generic data store backed by json.
	return nil, nil
}

func (p *CommandPlugin) Help(bot *Bot, service Service) []string {
	help := make([]string, 0)
	for _, command := range p.commands {
		if command.help != nil {
			help = append(help, command.help(bot, service)...)
		}
	}
	return help
}

func (p *CommandPlugin) Message(bot *Bot, service Service, message Message) {
	if !service.IsMe(message) {
		for commandString, command := range p.commands {
			if matchesCommand(commandString, message) {
				command.message(bot, service, message)
			}
		}
	}
}

func (p *CommandPlugin) AddCommand(commandString string, message MessageFunc, help HelpFunc) {
	p.commands[commandString] = &Command{
		message: message,
		help:    help,
	}
}

func (p *CommandPlugin) AddSimpleCommand(commandString string, message CommandMessageFunc, help string) {
	var helpFunc HelpFunc
	if help != "" {
		helpFunc = func(bot *Bot, service Service) []string { return []string{help} }
	}
	p.AddCommand(commandString, NewCommandMessageFunc(message), helpFunc)
}

func NewCommandPlugin() *CommandPlugin {
	return &CommandPlugin{make(map[string]*Command)}
}
