package comicplugin

import (
	"github.com/iopred/bruxism"
	"testing"
	// "time"
)


// func TestComicPlugin(t *testing.T) {
// 	c := New()

// 	b := bruxism.NewBot()
// 	s := bruxism.NewMockService().SetName("discord").SetUserID("32").SetUserName("Septapus")
	
// 	b.RegisterService(s)
// 	b.RegisterPlugin(s, c)
// 	b.Open()

// 	s.NewMessage(bruxism.NewMockMessage().SetMessage("@Septapus customcomicsimple jordy: hello | xeno: hi!").SetUserName("James").SetUserID("0").SetChannel("#lgv"))
	
// 	time.Sleep(1 * time.Second)	
// }

func TestMessagesFromChatlog(t *testing.T) {
	c := &comicPlugin{
		log: make(map[string][]bruxism.Message),
	}

	b := bruxism.NewBot()
	s := bruxism.NewMockService().SetName("discord").SetUserID("32").SetUserName("Septapus")
	
	b.RegisterService(s)
	b.RegisterPlugin(s, c)
	b.Open()

	messages, err := c.MessagesFromChatlog(s, bruxism.NewMockMessage().SetMessage("customcomicsimple   jordy:    hello    |   xeno:   hi! ").SetUserName("James").SetUserID("0").SetChannel("#lgv"))
	if err != nil {
		t.Fatal(err)
	}

	type messageTest struct {
        author string
        text   string
	}

	var messageTests = []messageTest {
	    {"jordy", "hello"}, {"xeno", "hi!"},
	}


	if len(messageTests) != len(messages) {
		t.Errorf("mismatched lengths: expected: %d, actual %d", len(messageTests), len(messages))
	}

	for i, mt := range messageTests {
        if mt.author != messages[i].Author {
        	t.Errorf("Author: expected: %s, actual %s", mt.author, messages[i].Author)
        }
        if mt.text != messages[i].Text {
        	t.Errorf("Text: expected: %s, actual %s", mt.text, messages[i].Text)
        }
    }
}