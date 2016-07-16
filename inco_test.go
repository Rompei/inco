package inco

import (
	"os"
	"testing"
)

func TestPost(t *testing.T) {
	msg := &Message{
		Username:  "test_bot",
		Text:      "Test",
		IconEmoji: ":ghost:",
		Channel:   "#test",
	}
	if err := Incoming(os.Getenv("SLACK_URL"), msg); err != nil {
		t.Errorf(err.Error())
	}
}
