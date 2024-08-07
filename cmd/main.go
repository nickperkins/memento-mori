package main

import (
	"context"
	"log"
	"syscall"

	"github.com/nickperkins/momento-mori/internal/bot"
	"github.com/nickperkins/momento-mori/internal/config"
	"github.com/nickperkins/momento-mori/internal/flags"
	"github.com/nickperkins/momento-mori/internal/quotes"
)

func init() {
	// Set up logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// command line arguments
	flags.SetupFlags()

}

func main() {
	// Parse command line flags
	flags.ParseFlags()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create a new Mastodon bot
	b := bot.NewBot(context.Background(), cfg)

	// Load the quotes
	err = quotes.LoadQuotes(b.PostsDirectory)
	if err != nil {
		log.Fatalf("Failed to load quotes: %v", err)
		syscall.Exit(1)
	}

	// loop forever until the program is killed
	for {

		// Get a random quote
		quote, err := b.GetRandomQuote()
		if err != nil {
			log.Fatalf("Failed to load posts: %v", err)
		}

		// Add the standard hashtags
		quote += "\n\n#MementoMori #Death"

		// Post the random post on Mastodon
		err = b.PostToot(quote)
		if err != nil {
			log.Fatalf("Failed to post on Mastodon: %v", err)
		}
		if flags.Flags.RunOnce {
			break
		}
		b.Sleep()
	}
}
