package bruxism

import (
	"testing"
)

func TestCommandMatches(t *testing.T) {
	tests := []struct {
		Prefix string
		Matches bool
		Command string
		Message string
		Argument string
	}{
		{"", true, "hello", "hello", ""},
		{"", true, "hello", " hello", ""},
		{"", true, "hello", "hello ", ""},
		{"", true, "hello", " hello ", ""},
		{"", true, "hello", "hello world", "world"},
		{"", true, "hello", " hello world", "world"},
		{"", true, "hello", "hello world ", "world"},
		{"", true, "hello", " hello world ", "world"},
		{"", false, "hello", " helloworld ", ""},
		{"", true, "", " hello world ", ""},
		{"", true, "", " hello world ", ""},
		{"!", true, "hello", "!hello", ""},
		{"!", true, "hello", " !hello", ""},
		{"!", true, "hello", "!hello ", ""},
		{"!", true, "hello", " !hello ", ""},
		{"!", true, "hello", "!hello world", "world"},
		{"!", true, "hello", " !hello world", "world"},
		{"!", true, "hello", "!hello world ", "world"},
		{"!", true, "hello", " !hello world ", "world"},
		{"!", false, "hello", " !helloworld ", ""},
		{"!", false, "hello", " ! hello world ", ""},
		{"!", false, "hello", "!", ""},
		{"!", false, "", "hello there", ""},
		{"!", false, "", "sup mang", ""},
		{"@Septapus ", true, "hello", "@Septapus hello", ""},
		{"@Septapus ", true, "hello", " @Septapus hello", ""},
		{"@Septapus ", true, "hello", "@Septapus hello ", ""},
		{"@Septapus ", true, "hello", " @Septapus hello ", ""},
		{"@Septapus ", true, "hello", "@Septapus hello world", "world"},
		{"@Septapus ", true, "hello", " @Septapus hello world", "world"},
		{"@Septapus ", true, "hello", "@Septapus hello world ", "world"},
		{"@Septapus ", true, "hello", " @Septapus hello world ", "world"},
		{"@Septapus ", false, "hello", " @Septapushello world ", ""},
		{"@Septapus ", false, "hello", "@Septapus", ""},
		{"@Septapus ", false, "", "hello there ", ""},
		{"@Septapus ", false, "", "sup mang", ""},
	}

	for _, test := range tests {
		m := &DiscordMessage{
			Content: &test.Message,
		}
		matches, ok := m.MatchesCommand(test.Prefix, test.Command)
		if !ok {
			t.Errorf("Not ok with message `%v`. Prefix: %v. Command: %v", test.Message, test.Prefix, test.Command)
		}
		if matches != test.Matches {
			t.Errorf("Incorrect matches with message `%v`: Have %v, Want %v. Prefix: %v. Command: %v", test.Message, matches, test.Matches, test.Prefix, test.Command)
		}
		if test.Argument != "" {
			if str, _, ok := m.ParseCommand(test.Prefix); !ok || str != test.Argument {
				t.Errorf("No argument %v with message %v", test.Argument, test.Message)
			}
		}
	}
}

func TestCommandMatchesMultiple(t *testing.T) {
	text := "@Septapus hello world"
	m := &DiscordMessage{
		Content: &text,
	}
	if matches, ok := m.MatchesCommand("@Septapus ", "hello"); !matches || !ok {
		t.Errorf("Match failure")
	}

	if matches, ok := m.MatchesCommand("@Septapus ", "butt"); matches || !ok {
		t.Errorf("Match failure")
	}

	if matches, ok := m.MatchesCommand("@Septapus ", "hello"); !matches || !ok {
		t.Errorf("Match failure")
	}
}

func TestCommandParse(t *testing.T) {
	text := "@Septapus hello world  how are you   doing"
	m := &DiscordMessage{
		Content: &text,
	}
	if rest, parts, ok := m.ParseCommand("@Septapus "); ok {
		if rest != "world how are you doing" {
			t.Errorf("Command rest incorrect: %v", rest)
		}
		if len(parts) != 5  {
			t.Errorf("Parts length incorrect: %v", len(parts))
		}
	} else {
		t.Errorf("Match failure")
	}
}