package quotes

import (
	"bytes"
	"html/template"
	"math/rand"

	"github.com/BurntSushi/toml"
)

type (
	Quote struct {
		Text   string
		Author string
	}
	Quotes struct {
		Quotes []Quote `toml:"quote"`
	}
)

var quotes Quotes

// Load
func LoadQuotes(filePath string) error {
	// check if the file exists

	_, err := toml.DecodeFile(filePath, &quotes)
	if err != nil {
		return err
	}

	return nil

}

// GetRandomQuote loads a random quote from the loaded quotes.
func LoadRandomQuote(postsDirectory string) Quote {
	LoadQuotes(postsDirectory)
	randomIndex := rand.Intn(len(quotes.Quotes))
	return quotes.Quotes[randomIndex]

}

func FormatQuoteAsToot(quote Quote) string {

	tmpl := template.Must(template.New("quote").Parse("\"{{.Text}}\" - {{.Author}}"))
	buf := new(bytes.Buffer)
	tmpl.Execute(buf, quote)
	return buf.String()

}
