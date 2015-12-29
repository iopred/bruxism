package bot

type SimplePlugin struct {
	name    string
	load    LoadFunc
	save    SaveFunc
	message MessageFunc
	help    HelpFunc
}

// The name of the plugin.
func (p *SimplePlugin) Name() string {
	return p.name
}

// Loads plugin state from a byte array.
func (p *SimplePlugin) Load(bot *Bot, service Service, data []byte) error {
	if p.load != nil {
		return p.load(bot, service, data)
	}
	return nil
}

// Saves plugin state to a byte array.
func (p *SimplePlugin) Save() ([]byte, error) {
	if p.save != nil {
		return p.save()
	}
	return nil, nil
}

// Returns a list of help strings that are printed when the user requests them.
func (p *SimplePlugin) Help(bot *Bot, service Service) []string {
	if p.help != nil {
		return p.help(bot, service)
	}
	return nil
}

// Message handler.
func (p *SimplePlugin) Message(bot *Bot, service Service, message Message) {
	if p.message != nil {
		p.message(bot, service, message)
	}
}

func NewSimplePlugin(name string) *SimplePlugin {
	return &SimplePlugin{name: name}
}
