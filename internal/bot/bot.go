package bot

import (
	"fmt"
	"time"

	"github.com/mattn/go-mastodon"
	"github.com/nickperkins/momento-mori/internal/config"
	"github.com/nickperkins/momento-mori/internal/flags"
	"github.com/nickperkins/momento-mori/internal/quotes"
	"golang.org/x/net/context"
)

// Bot represents the Mastodon bot.
type Bot struct {
	client         *mastodon.Client
	PostsDirectory string
	SleepDuration  int // in minutes
}

// NewBot creates a new instance of the Mastodon bot.f
func NewBot(config *config.Config) *Bot {
	c := mastodon.NewClient(&mastodon.Config{
		Server:      config.MastodonInstanceURL,
		AccessToken: config.AccessToken,
	})
	return &Bot{
		client:         c,
		PostsDirectory: config.QuotesFile,
		SleepDuration:  config.SleepDuration,
	}
}

// GetRandomQuote loads a random post from the directory of files.
func (b *Bot) GetRandomQuote() (string, error) {
	return quotes.FormatQuoteAsToot(quotes.LoadRandomQuote(b.PostsDirectory)), nil
}

// PostToot posts the given content on Mastodon.
func (b *Bot) PostToot(ctx context.Context, content string) error {
	if flags.Flags.DryRun {
		fmt.Printf("Dry-run mode: Would have posted:\n%s\n", content)
		return nil
	}
	result, err := b.client.PostStatus(ctx, &mastodon.Toot{
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
func (b *Bot) Sleep(ctx context.Context) {
	waitTime := b.SleepDuration * 60
	fmt.Printf("Sleeping for %d minutes\n", b.SleepDuration)
	// Sleep for the configured duration except if the context is cancelled
	// This allows the bot to exit immediately if the context is cancelled
	select {
	case <-ctx.Done():
		return
	case <-time.After(time.Duration(waitTime) * time.Second):
		return
	}

}
