# Memento Mori

Memento Mori is a Mastodon bot written in Go that posts a random quote as a reminder of our mortality.

## Features

- Posts random quotes about mortality to a Mastodon account.
- Configurable posting intervals.
- Easy to deploy and run.

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/memento-mori.git
    cd memento-mori
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Build the project:

    ```sh
    go build -o memento-mori cmd/main.go
    ```

## Configuration

Create a `.env` file in the root directory with the following variables:

```sh
MASTODON_SERVER=<https://mastodon.example.com>
ACCESS_TOKEN=your_access_token
QUOTES_FILE=<path to your quote file>
POSTING_INTERVAL=<time in minutes>
```
