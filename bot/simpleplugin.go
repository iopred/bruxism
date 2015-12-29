package bot

// SimplePlugin is a simple wrapper around a Plugin that can implement handlers by function reference.
type SimplePlugin struct {
	name    string
	load    LoadFunc
	save    SaveFunc
	message MessageFunc
	help    HelpFunc
}

// Name returns the name of the plugin.
func (p *SimplePlugin) Name() string {
	return p.name
}

// Load will load plugin state from a byte array.
func (p *SimplePlugin) Load(bot *Bot, service Service, data []byte) error {
	if p.load != nil {
		return p.load(bot, service, data)
	}
	return nil
}

// Save will save plugin state to a byte array.
func (p *SimplePlugin) Save() ([]byte, error) {
	if p.save != nil {
		return p.save()
	}
	return nil, nil
}

// Help returns a list of help strings that are printed when the user requests them.
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
