package bruxism

// SimplePlugin is a simple wrapper around a Plugin that can implement handlers by function reference.
type SimplePlugin struct {
	name        string
	LoadFunc    LoadFunc    `json:"-"`
	SaveFunc    SaveFunc    `json:"-"`
	MessageFunc MessageFunc `json:"-"`
	HelpFunc    HelpFunc    `json:"-"`
	StatsFunc   StatsFunc   `json:"-"`
}

// Name returns the name of the plugin.
func (p *SimplePlugin) Name() string {
	return p.name
}

// Load will load plugin state from a byte array.
func (p *SimplePlugin) Load(bot *Bot, service Service, data []byte) error {
	if p.LoadFunc != nil {
		return p.LoadFunc(bot, service, data)
	}
	return nil
}

// Save will save plugin state to a byte array.
func (p *SimplePlugin) Save() ([]byte, error) {
	if p.SaveFunc != nil {
		return p.SaveFunc()
	}
	return nil, nil
}

// Help returns a list of help strings that are printed when the user requests them.
func (p *SimplePlugin) Help(bot *Bot, service Service, message Message, detailed bool) []string {
	if p.HelpFunc != nil {
		return p.HelpFunc(bot, service, message, detailed)
	}
	return nil
}

// Message handler.
func (p *SimplePlugin) Message(bot *Bot, service Service, message Message) {
	defer MessageRecover()
	if p.MessageFunc != nil {
		p.MessageFunc(bot, service, message)
	}
}

func (p *SimplePlugin) Stats(bot *Bot, service Service, message Message) []string {
	if p.StatsFunc != nil {
		return p.StatsFunc(bot, service, message)
	}
	return nil
}

// NewSimplePlugin creates a new simple plugin.
func NewSimplePlugin(name string) *SimplePlugin {
	return &SimplePlugin{name: name}
}
