package bot

type SimplePlugin struct {
	name    string
	load    LoadFunc
	save    SaveFunc
	message MessageFunc
	help    HelpFunc
}

func (p *SimplePlugin) Name() string {
	return p.name
}

func (p *SimplePlugin) Load(bot *Bot, service Service, data []byte) error {
	if p.load != nil {
		return p.load(bot, service, data)
	}
	return nil
}

func (p *SimplePlugin) Save() ([]byte, error) {
	if p.save != nil {
		return p.save()
	}
	return nil, nil
}

func (p *SimplePlugin) Help(bot *Bot, service Service) []string {
	if p.help != nil {
		return p.help(bot, service)
	}
	return nil
}

func (p *SimplePlugin) Message(bot *Bot, service Service, message Message) {
	if p.message != nil {
		p.message(bot, service, message)
	}
}

func NewSimplePlugin(name string) *SimplePlugin {
	return &SimplePlugin{name: name}
}
