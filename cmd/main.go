package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nickperkins/momento-mori/internal/bot"
	"github.com/nickperkins/momento-mori/internal/config"
	"github.com/nickperkins/momento-mori/internal/flags"
	"github.com/nickperkins/momento-mori/internal/quotes"
)

var Version = "dev"

func init() {
	// Set up logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// command line arguments
	flags.SetupFlags(Version)

}

// main is the entry point of the application.
// It parses command line flags, loads the configuration, creates a new Mastodon bot,
// loads the quotes, and runs the bot.
func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	// Parse command line flags
	flags.ParseFlags()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create a new Mastodon bot
	b := bot.NewBot(cfg)

	// Load the quotes
	err = quotes.LoadQuotes(b.PostsDirectory)
	if err != nil {
		log.Fatalf("Failed to load quotes: %v", err)
		syscall.Exit(1)
	}

	// Run the bot
	runBot(ctx, b)

}

// runBot runs the bot indefinitely, periodically posting quotes.
// It takes a context and a bot instance as parameters.
// If the context is cancelled, the function will shut down the bot and return.
// If an error occurs while posting a quote, it will be logged.
// If the "RunOnce" flag is set, the function will return after posting a quote.
// After each iteration, the bot will sleep for a certain duration.
func runBot(ctx context.Context, b *bot.Bot) {
	for {

		select {
		case <-ctx.Done():
			log.Println("Context cancelled, shutting down bot.")
			return
		default:
			if err := postQuote(ctx, b); err != nil {
				log.Printf("Failed to post quote: %v", err)
			}
			if flags.Flags.RunOnce {
				return
			}
			b.Sleep(ctx)
		}

	}
}

// postQuote is a function that posts a random quote to Mastodon using the given bot.
// It retrieves a random quote using the GetRandomQuote method of the bot.
// If an error occurs while retrieving the quote, it returns an error with a descriptive message.
// After retrieving the quote, it adds the standard hashtags "#MementoMori" and "#Death" to the quote.
// Finally, it posts the quote to Mastodon using the PostToot method of the bot.
// If an error occurs while posting the quote, it returns an error with a descriptive message.
// If the function executes successfully, it returns nil.
func postQuote(ctx context.Context, b *bot.Bot) error {
	quote, err := b.GetRandomQuote()
	if err != nil {
		return fmt.Errorf("failed to get quote: %v", err)
	}

	// Add the standard hashtags
	quote += "\n\n#MementoMori #Death"

	if err = b.PostToot(ctx, quote); err != nil {
		return fmt.Errorf("failed to toot on Mastodon: %v", err)

	}
	return nil
}
