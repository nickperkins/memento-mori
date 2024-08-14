package bot

import (
	"testing"

	"github.com/nickperkins/momento-mori/internal/config"
)

func TestNewBot(t *testing.T) {
	t.Run("NewBot", func(t *testing.T) {
		// test the NewBot function
		// create a new bot
		b := NewBot(&config.Config{AppEnv: "test", MastodonInstanceURL: "https://mastodon.social", AccessToken: "access_token", QuotesFile: "quotes.txt", SleepDuration: 5})
		// check if the bot was created correctly
		if b.client != nil {
			if b.client.Config.Server != "https://mastodon.social" {
				t.Errorf("Expected Server to be https://mastodon.social, got %s", b.client.Config.Server)
			}
			if b.client.Config.AccessToken != "access_token" {
				t.Errorf("Expected AccessToken to be access_token, got %s", b.client.Config.AccessToken)
			}

		}
		if b.PostsDirectory != "quotes.txt" {
			t.Errorf("Expected PostsDirectory to be quotes.txt, got %s", b.PostsDirectory)
		}
		if b.SleepDuration != 5 {
			t.Errorf("Expected SleepDuration to be 5, got %d", b.SleepDuration)
		}
	})
}
