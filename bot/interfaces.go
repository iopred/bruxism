package bot

type Message interface {
	Channel() string
	UserName() string
	UserId() string
	Message() string
	MessageId() string
	IsModerator() bool
}

type Service interface {
	Name() string
	UserName() string
	Open() (<-chan Message, error)
	IsMe(message Message) bool
	SendMessage(channel, message string) error
	DeleteMessage(messageId string) error
	BanUser(channel, user string, duration int) error
	SetPlaying(game string) error
	Join(join string) error
}

type LoadFunc func(*Bot, Service, []byte) error
type SaveFunc func() ([]byte, error)
type HelpFunc func(*Bot, Service) []string
type MessageFunc func(*Bot, Service, Message)

type Plugin interface {
	Name() string
	Load(*Bot, Service, []byte) error
	Save() ([]byte, error)
	Help(*Bot, Service) []string
	Message(*Bot, Service, Message)
}
