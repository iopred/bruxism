package bruxism

import (
	"bytes"
	"encoding/json"
	"errors"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"runtime/debug"
)

// VersionString is the current version of the bot
const VersionString string = "0.3.1"

type serviceEntry struct {
	Service
	Plugins         map[string]Plugin
	messageChannels []chan Message
}

// Bot enables registering of Services and Plugins.
type Bot struct {
	Services map[string]*serviceEntry
	ImgurID  string
}

func messageRecover() {
	if r := recover(); r != nil {
		log.Println("Recovered:", string(debug.Stack()))
	}
}

// NewBot will create a new bot.
func NewBot() *Bot {
	return &Bot{
		Services: make(map[string]*serviceEntry, 0),
	}
}

func (b *Bot) getData(service Service, plugin Plugin) []byte {
	if b, err := ioutil.ReadFile(service.Name() + "/" + plugin.Name()); err == nil {
		return b
	}
	return nil
}

// RegisterService registers a service with the bot.
func (b *Bot) RegisterService(service Service) {
	if b.Services[service.Name()] != nil {
		log.Println("Service with that name already registered", service.Name())
	}
	serviceName := service.Name()
	b.Services[serviceName] = &serviceEntry{
		Service: service,
		Plugins: make(map[string]Plugin, 0),
	}
}

// RegisterPlugin registers a plugin on a service.
func (b *Bot) RegisterPlugin(service Service, plugin Plugin) {
	s := b.Services[service.Name()]
	if s.Plugins[plugin.Name()] != nil {
		log.Println("Plugin with that name already registered", plugin.Name())
	}
	s.Plugins[plugin.Name()] = plugin
	plugin.Load(b, service, b.getData(service, plugin))
}

func (b *Bot) listen(service Service, messageChan <-chan Message) {
	serviceName := service.Name()
	for {
		message := <-messageChan
		log.Printf("<%s> %s: %s\n", message.Channel(), message.UserName(), message.Message())
		plugins := b.Services[serviceName].Plugins
		for _, plugin := range plugins {
			go plugin.Message(b, service, message)
		}
	}
}

// Open will open all the current services and begins listening.
func (b *Bot) Open() {
	for _, service := range b.Services {
		if messageChan, err := service.Open(); err == nil {
			log.Println("Started service", service.Name())
			go b.listen(service, messageChan)
		} else {
			log.Printf("Error creating service %v: %v\n", service.Name(), err)
		}
	}
}

// Save will save the current plugin state for all plugins on all services.
func (b *Bot) Save() {
	for _, service := range b.Services {
		serviceName := service.Name()
		if err := os.Mkdir(serviceName, os.ModePerm); err != nil {
			if !os.IsExist(err) {
				log.Println("Error creating service directory.")
			}
		}
		for _, plugin := range service.Plugins {
			if data, err := plugin.Save(); err != nil {
				log.Printf("Error saving plugin %v %v. %v", serviceName, plugin.Name(), err)
			} else if data != nil {
				if err := ioutil.WriteFile(serviceName+"/"+plugin.Name(), data, os.ModePerm); err != nil {
					log.Printf("Error saving plugin %v %v. %v", serviceName, plugin.Name(), err)
				}
			}
		}
	}
}

// UploadToImgur uploads image data to Imgur and returns the url to it.
func (b *Bot) UploadToImgur(image image.Image, filename string) (string, error) {
	log.Println("Beginning upload.")

	if b.ImgurID == "" {
		return "", errors.New("No Imgur client ID provided.")
	}

	bodyBuf := &bytes.Buffer{}
	bodywriter := multipart.NewWriter(bodyBuf)

	writer, err := bodywriter.CreateFormFile("image", filename)
	if err != nil {
		return "", err
	}

	err = png.Encode(writer, image)
	if err != nil {
		return "", err
	}

	contentType := bodywriter.FormDataContentType()
	bodywriter.Close()

	r, err := http.NewRequest("POST", "https://api.imgur.com/3/image", bodyBuf)
	if err != nil {
		return "", err
	}

	r.Header.Set("Content-Type", contentType)
	r.Header.Set("Authorization", "Client-ID "+b.ImgurID)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New(string(body))
	}

	j := make(map[string]interface{})

	err = json.Unmarshal(body, &j)
	if err != nil {
		return "", err
	}

	return j["data"].(map[string]interface{})["link"].(string), nil
}
