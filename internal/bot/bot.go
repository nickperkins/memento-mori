package bot

import (
	"context"
	"fmt"
	"time"

	"github.com/mattn/go-mastodon"
	"github.com/nickperkins/momento-mori/internal/config"
	"github.com/nickperkins/momento-mori/internal/flags"
	"github.com/nickperkins/momento-mori/internal/quotes"
)

// Bot represents the Mastodon bot.
type Bot struct {
	client         *mastodon.Client
	PostsDirectory string
	SleepDuration  int // in minutes
	context        context.Context
}

// NewBot creates a new instance of the Mastodon bot.f
func NewBot(context context.Context, config *config.Config) *Bot {
	c := mastodon.NewClient(&mastodon.Config{
		Server:      config.MastodonInstanceURL,
		AccessToken: config.AccessToken,
	})
	return &Bot{
		client:         c,
		PostsDirectory: config.QuotesFile,
		SleepDuration:  config.SleepDuration,
		context:        context,
	}
}

// GetRandomQuote loads a random post from the directory of files.
func (b *Bot) GetRandomQuote() (string, error) {
	return quotes.FormatQuoteAsToot(quotes.LoadRandomQuote(b.PostsDirectory)), nil
}

// PostToot posts the given content on Mastodon.
func (b *Bot) PostToot(content string) error {
	if flags.Flags.DryRun {
		fmt.Printf("Dry-run mode: Would have posted:\n%s\n", content)
		return nil
	}
	result, err := b.client.PostStatus(b.context, &mastodon.Toot{
		Status:      content,
		SpoilerText: "Death",
		Visibility:  "unlisted",
	})
	if err != nil {
		return fmt.Errorf("failed to post status: %w", err)
	}
	fmt.Printf("Post successfully posted on Mastodon: %s\n", result.URL)
	return nil
}

// Sleeps for the configured duration.
func (b *Bot) Sleep() {
	waitTime := b.SleepDuration * 60
	fmt.Printf("Sleeping for %d minutes\n", b.SleepDuration)
	<-time.After(time.Duration(waitTime) * time.Second)

}
