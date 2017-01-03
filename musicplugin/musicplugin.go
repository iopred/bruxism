package musicplugin

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/iopred/bruxism"
	"github.com/iopred/discordgo"
)

type MusicPlugin struct {
	sync.Mutex

	discord *bruxism.Discord

	VoiceConnections map[string]*voiceConnection
}

type voiceConnection struct {
	sync.Mutex
	debug bool

	GuildID      string
	ChannelID    string
	MaxQueueSize int
	Queue        []song

	close   chan struct{}
	control chan controlMessage
	playing *song
	conn    *discordgo.VoiceConnection
}

type controlMessage int

const (
	Skip controlMessage = iota
	Pause
	Resume
)

type song struct {
	AddedBy     string
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	FullTitle   string `json:"full_title"`
	Thumbnail   string `json:"thumbnail"`
	URL         string `json:"webpage_url"`
	Duration    int    `json:"duration"`
	Remaining   int
}

// New will create a new music plugin.
func New(discord *bruxism.Discord) bruxism.Plugin {

	p := &MusicPlugin{
		discord:          discord,
		VoiceConnections: make(map[string]*voiceConnection),
	}

	return p
}

// Name returns the name of the plugin.
func (p *MusicPlugin) Name() string {
	return "Music"
}

// Load will load plugin state from a byte array.
func (p *MusicPlugin) Load(bot *bruxism.Bot, service bruxism.Service, data []byte) (err error) {
	if service.Name() != bruxism.DiscordServiceName {
		panic("Music Plugin only supports Discord.")
	}

	if data != nil {
		if err = json.Unmarshal(data, p); err != nil {
			log.Println("musicplugin: loading data err:", err)
		}
	}

	go p.init()

	return nil
}

func (p *MusicPlugin) init() {
	<-time.After(1 * time.Second)
	for _, s := range p.discord.Sessions {
		if !s.DataReady {
			go p.init()
			return
		}
	}
	p.ready()
}

func (p *MusicPlugin) ready() {
	// Join all registered voice channels and start the playback queue
	for _, v := range p.VoiceConnections {
		if v.ChannelID == "" {
			continue
		}
		vc, err := p.join(v.ChannelID)
		if err != nil {
			log.Println("musicplugin: join channel err:", err)
			continue
		}
		p.gostart(vc)
	}
}

// Save will save plugin state to a byte array.
func (p *MusicPlugin) Save() ([]byte, error) {
	return json.Marshal(p)
}

// Help returns a list of help strings that are printed when the user requests them.
func (p *MusicPlugin) Help(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, detailed bool) []string {
	if service.IsPrivate(message) {
		return nil
	}

	// Only show help messages for guilds where we have a voice connection
	c, err := p.discord.Channel(message.Channel())
	if err != nil {
		log.Println("musicplugin: fetching channel err:", err.Error())
		return nil
	}

	_, ok := p.VoiceConnections[c.GuildID]
	if !ok {
		return nil
	}

	help := []string{
		bruxism.CommandHelp(service, "music", "<command>", "Music Plugin, see `help music`")[0],
	}

	if detailed {
		help = append(help, []string{
			"Examples:",
			bruxism.CommandHelp(service, "music", "join [channelid]", "Join your voice channel or the provided voice channel.")[0],
			bruxism.CommandHelp(service, "music", "leave", "Leave current voice channel.")[0],
			bruxism.CommandHelp(service, "music", "play [url]", "Start playing music and optionally enqueue provided url.")[0],
			bruxism.CommandHelp(service, "music", "info", "Information about this plugin and the currently playing song.")[0],
			bruxism.CommandHelp(service, "music", "pause", "Pause playback of current song.")[0],
			bruxism.CommandHelp(service, "music", "resume", "Resume playback of current song.")[0],
			bruxism.CommandHelp(service, "music", "skip", "Skip current song.")[0],
			bruxism.CommandHelp(service, "music", "stop", "Stop playing music.")[0],
			bruxism.CommandHelp(service, "music", "list", "List contents of queue.")[0],
			bruxism.CommandHelp(service, "music", "clear", "Clear all items from queue.")[0],
		}...)
	}

	return help
}

// Message handler.
func (p *MusicPlugin) Message(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	defer bruxism.MessageRecover()

	if service.IsMe(message) {
		return
	}

	if !bruxism.MatchesCommand(service, "music", message) && !bruxism.MatchesCommand(service, "mu", message) {
		return
	}

	if service.IsPrivate(message) {
		service.SendMessage(message.Channel(), "Sorry, this command doesn't work in private chat.")
		return
	}

	_, parts := bruxism.ParseCommand(service, message)

	if len(parts) == 0 {
		service.SendMessage(message.Channel(), strings.Join(p.Help(bot, service, message, true), "\n"))
		return
	}

	// Get the Channel (and GuildID) for this channel because it's needed in
	// a few locations below
	channel, err := p.discord.Channel(message.Channel())
	if err != nil {
		log.Println("musicplugin: fetching channel err:", err.Error())
		return
	}

	// grab pointer to this channels voice connection, if exists.
	vc, vcok := p.VoiceConnections[channel.GuildID]

	switch parts[0] {

	case "help":
		// display extended help information
		service.SendMessage(message.Channel(), strings.Join(p.Help(bot, service, message, true), "\n"))

	case "stats":
		// TODO: maybe provide plugin stats, total channels, total song queues, etc

	case "join":
		if !service.IsBotOwner(message) {
			service.SendMessage(message.Channel(), "Sorry, only bot owner can join, please ask him to run this command.")
			return
		}
		// join the voice channel of the caller or the provided channel ID

		channelID := ""
		if len(parts) > 1 {
			channelID = parts[1]
		}

		if channelID == "" {
			messageUserID := message.UserID()
			for _, g := range p.discord.Guilds() {
				for _, v := range g.VoiceStates {
					if v.UserID == messageUserID {
						channelID = v.ChannelID
					}
				}
			}

			if channelID == "" {
				service.SendMessage(message.Channel(), "I couldn't find you in any voice channels, please join one.")
				return
			}
		}

		_, err := p.join(channelID)
		if err != nil {
			service.SendMessage(message.Channel(), err.Error())
			break
		}

		service.SendMessage(message.Channel(), "Now, let's play some music!")

	case "leave":
		// leave voice channel for this Guild
		if !service.IsBotOwner(message) {
			service.SendMessage(message.Channel(), "Sorry, only bot owner can leave, please ask him to run this command.")
			return
		}

		if !vcok {
			service.SendMessage(message.Channel(), "There is no voice connection for this Guild.")
		}

		vc.conn.Close()
		delete(p.VoiceConnections, channel.GuildID)
		service.SendMessage(message.Channel(), "Closed voice connection.")

	case "debug":
		// enable or disable debug

		if !vcok {
			service.SendMessage(message.Channel(), fmt.Sprintf("There is no voice connection for this Guild."))
		}

		vc.Lock()
		vc.debug = !vc.debug
		service.SendMessage(message.Channel(), fmt.Sprintf("debug mode set to %v", vc.debug))
		vc.Unlock()

	case "play":
		// Start queue player and optionally enqueue provided songs

		p.gostart(vc)

		for _, v := range parts[1:] {
			url, err := url.Parse(v) // doesn't check much..
			if err != nil {
				continue
			}
			err = p.enqueue(vc, url.String(), service, message)
			if err != nil {
				// TODO: Might need improving.
				service.SendMessage(message.Channel(), err.Error())
			}
		}

	case "stop":
		// stop the queue player

		if vc.close != nil {
			close(vc.close)
			vc.close = nil
		}

		if vc.control != nil {
			close(vc.control)
			vc.control = nil
		}

	case "skip":
		// skip current song

		if vc.control == nil {
			return
		}
		vc.control <- Skip

	case "pause":
		// pause the queue player
		if vc.control == nil {
			return
		}
		vc.control <- Pause

	case "resume":
		// resume the queue player
		if vc.control == nil {
			return
		}
		vc.control <- Resume

	case "info":
		// report player settings, queue info, and current song

		msg := fmt.Sprintf("`Bruxism MusicPlugin:`\n")
		msg += fmt.Sprintf("`Voice Channel:` %s\n", vc.ChannelID)
		msg += fmt.Sprintf("`Queue Size:` %d\n", len(vc.Queue))

		if vc.playing == nil {
			service.SendMessage(message.Channel(), msg)
			break
		}

		msg += fmt.Sprintf("`Now Playing:`\n")
		msg += fmt.Sprintf("`ID:` %s\n", vc.playing.ID)
		msg += fmt.Sprintf("`Title:` %s\n", vc.playing.Title)
		msg += fmt.Sprintf("`Duration:` %ds\n", vc.playing.Duration)
		msg += fmt.Sprintf("`Remaining:` %ds\n", vc.playing.Remaining)
		msg += fmt.Sprintf("`Source URL:` <%s>\n", vc.playing.URL)
		msg += fmt.Sprintf("`Thumbnail:` %s\n", vc.playing.Thumbnail)
		service.SendMessage(message.Channel(), msg)

	case "list":
		// list top items in the queue

		if len(vc.Queue) == 0 {
			service.SendMessage(message.Channel(), "The music queue is empty.")
			return
		}

		var msg string

		i := 1
		i2 := 0
		for k, v := range vc.Queue {
			np := ""
			if vc.playing != nil && *vc.playing == v {
				np = "**(Now Playing)**"
			}
			d := time.Duration(v.Duration) * time.Second
			msg += fmt.Sprintf("`%.3d:%.15s` **%s** [%s] - *%s* %s\n", k, v.ID, v.Title, d.String(), v.AddedBy, np)

			if i >= 15 {
				service.SendMessage(message.Channel(), msg)
				msg = ""
				i = 0
				i2++

				if i2 >= 8 {
					// limit response to 8 messages (120 songs)
					return
				}
			}
			i++
		}

		service.SendMessage(message.Channel(), msg)

	case "clear":
		// clear all items from the queue
		vc.Lock()
		vc.Queue = []song{}
		vc.Unlock()

	default:
		service.SendMessage(message.Channel(), "Unknown music command, try `help music`")
	}
}

// join a specific voice channel
func (p *MusicPlugin) join(cID string) (vc *voiceConnection, err error) {

	c, err := p.discord.Channel(cID)
	if err != nil {
		return
	}

	if c.Type != "voice" {
		err = fmt.Errorf("That's not a voice channel.")
		return
	}

	// Get or Create the VoiceConnection object
	p.Lock()
	vc, ok := p.VoiceConnections[c.GuildID]
	if !ok {
		vc = &voiceConnection{}
		p.VoiceConnections[c.GuildID] = vc
	}
	p.Unlock()

	guild, err := p.discord.Guild(c.GuildID)
	if err != nil {
		return
	}

	guildId, err := strconv.Atoi(guild.ID)
	if err != nil {
		return
	}

	// NOTE: Setting mute to false, deaf to true.
	vc.conn, err = p.discord.Sessions[(guildId>>22)%len(p.discord.Sessions)].ChannelVoiceJoin(c.GuildID, cID, false, true)
	if err != nil {
		return
	}

	vc.GuildID = c.GuildID
	vc.ChannelID = cID

	return
}

// enqueue a song/playlest to a VoiceConnections Queue
func (p *MusicPlugin) enqueue(vc *voiceConnection, url string, service bruxism.Service, message bruxism.Message) (err error) {

	if vc == nil {
		return fmt.Errorf("Cannot enqueue to nil voice connection.")
	}

	if url == "" {
		return fmt.Errorf("Cannot enqueue an empty string.")
	}

	// TODO //////////////////////////////////////////////////////////////////
	// need to parse the url and have a way to know what we're doing
	// 1) option to queue local files
	// 2) option to queue saved playlists
	// 3) option to queue URL that can be passed directly to ffmpeg without youtube-dl
	// 4) option to queue youtube-dl playlist
	// 5) option to queue youtube-dl song
	// 6) option to queue youtube-dl search result

	// right now option 4 and 5 work, only.
	//////////////////////////////////////////////////////////////////////////

	cmd := exec.Command("./youtube-dl", "-i", "-j", "--youtube-skip-dash-manifest", url)
	if vc.debug {
		cmd.Stderr = os.Stderr
	}

	output, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
		service.SendMessage(message.Channel(), fmt.Sprintf("Error adding song to playlist."))
		return
	}

	err = cmd.Start()
	if err != nil {
		log.Println(err)
		service.SendMessage(message.Channel(), fmt.Sprintf("Error adding song to playlist."))
		return
	}
	defer func() {
		go cmd.Wait()
	}()

	scanner := bufio.NewScanner(output)

	for scanner.Scan() {
		s := song{}
		err = json.Unmarshal(scanner.Bytes(), &s)
		if err != nil {
			log.Println(err)
			continue
		}

		s.AddedBy = message.UserName()

		vc.Lock()
		vc.Queue = append(vc.Queue, s)
		vc.Unlock()
	}
	return
}

// little wrapper function for start() to fire it off in a
// go routine if it is not already running.
func (p *MusicPlugin) gostart(vc *voiceConnection) (err error) {

	vc.Lock()

	if vc == nil {
		vc.Unlock()
		return fmt.Errorf("gostart cannot start a nil voice connection queue")
	}

	if vc.close != nil || vc.control != nil {
		vc.Unlock()
		return fmt.Errorf("gostart will not start a voice connection with non-nil control channels")
	}

	vc.close = make(chan struct{})
	vc.control = make(chan controlMessage)

	// TODO can this be moved lower?
	vc.Unlock()

	go p.start(vc, vc.close, vc.control)

	return
}

// "start" is a goroutine function that loops though the music queue and
// plays songs as they are added
func (p *MusicPlugin) start(vc *voiceConnection, close <-chan struct{}, control <-chan controlMessage) {

	if close == nil || control == nil || vc == nil {
		log.Println("musicplugin: start() exited due to nil channels")
		return
	}

	var i int
	var Song song

	// main loop keeps this going until close
	for {

		// exit if close channel is closed
		select {
		case <-close:
			log.Println("musicplugin: start() exited due to close channel.")
			return
		default:
		}

		// loop until voice connection is ready and songs are in the queue.
		if vc.conn == nil || vc.conn.Ready == false || len(vc.Queue) < 1 {
			time.Sleep(1 * time.Second)
			continue
		}

		// Get song to play and store it in local Song var
		vc.Lock()
		if len(vc.Queue)-1 >= i {
			Song = vc.Queue[i]
		} else {
			i = 0
			vc.Unlock()
			continue
		}
		vc.Unlock()

		vc.playing = &Song
		p.play(vc, close, control, Song)
		vc.playing = nil

		vc.Lock()
		if len(vc.Queue) > 0 {
			vc.Queue = append(vc.Queue[:i], vc.Queue[i+1:]...)
		}
		vc.Unlock()
	}
}

// play an individual song
func (p *MusicPlugin) play(vc *voiceConnection, close <-chan struct{}, control <-chan controlMessage, s song) {
	var err error

	if close == nil || control == nil || vc == nil || vc.conn == nil {
		log.Println("musicplugin: play exited because [close|control|vc|vc.conn] is nil.")
		return
	}

	ytdl := exec.Command("./youtube-dl", "-v", "-f", "bestaudio", "-o", "-", s.URL)
	if vc.debug {
		ytdl.Stderr = os.Stderr
	}
	ytdlout, err := ytdl.StdoutPipe()
	if err != nil {
		log.Println("musicplugin: ytdl StdoutPipe err:", err)
		return
	}
	ytdlbuf := bufio.NewReaderSize(ytdlout, 16384)

	ffmpeg := exec.Command("ffmpeg", "-i", "pipe:0", "-f", "s16le", "-ar", "48000", "-ac", "2", "pipe:1")
	ffmpeg.Stdin = ytdlbuf
	if vc.debug {
		ffmpeg.Stderr = os.Stderr
	}
	ffmpegout, err := ffmpeg.StdoutPipe()
	if err != nil {
		log.Println("musicplugin: ffmpeg StdoutPipe err:", err)
		return
	}
	ffmpegbuf := bufio.NewReaderSize(ffmpegout, 16384)

	dca := exec.Command("./dca", "-raw", "-i", "pipe:0")
	dca.Stdin = ffmpegbuf
	if vc.debug {
		dca.Stderr = os.Stderr
	}
	dcaout, err := dca.StdoutPipe()
	if err != nil {
		log.Println("musicplugin: dca StdoutPipe err:", err)
		return
	}
	dcabuf := bufio.NewReaderSize(dcaout, 16384)

	err = ytdl.Start()
	if err != nil {
		log.Println("musicplugin: ytdl Start err:", err)
		return
	}
	defer func() {
		go ytdl.Wait()
	}()

	err = ffmpeg.Start()
	if err != nil {
		log.Println("musicplugin: ffmpeg Start err:", err)
		return
	}
	defer func() {
		go ffmpeg.Wait()
	}()

	err = dca.Start()
	if err != nil {
		log.Println("musicplugin: dca Start err:", err)
		return
	}
	defer func() {
		go dca.Wait()
	}()

	// header "buffer"
	var opuslen int16

	// Send "speaking" packet over the voice websocket
	vc.conn.Speaking(true)

	// Send not "speaking" packet over the websocket when we finish
	defer vc.conn.Speaking(false)

	start := time.Now()
	for {

		select {
		case <-close:
			log.Println("musicplugin: play() exited due to close channel.")
			return
		default:
		}

		select {
		case ctl := <-control:
			switch ctl {
			case Skip:
				return
				break
			case Pause:
				done := false
				for {

					ctl, ok := <-control
					if !ok {
						return
					}
					switch ctl {
					case Skip:
						return
						break
					case Resume:
						done = true
						break
					}

					if done {
						break
					}

				}
			}
		default:
		}

		// read dca opus length header
		err = binary.Read(dcabuf, binary.LittleEndian, &opuslen)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return
		}
		if err != nil {
			log.Println("musicplugin: read opus length from dca err:", err)
			return
		}

		// read opus data from dca
		opus := make([]byte, opuslen)
		err = binary.Read(dcabuf, binary.LittleEndian, &opus)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return
		}
		if err != nil {
			log.Println("musicplugin: read opus from dca err:", err)
			return
		}

		// Send received PCM to the sendPCM channel
		vc.conn.OpusSend <- opus
		// TODO: Add a select and timeout to above
		// shouldn't ever block longer than maybe 18-25ms

		// this can cause a panic if vc becomes nil while waiting to send
		// on the opus channel. TODO fix..
		vc.playing.Remaining = (vc.playing.Duration - int(time.Since(start).Seconds()))

	}
}

// Stats will return the stats for a plugin.
func (p *MusicPlugin) Stats(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) []string {
	return nil
}
