package telegram

import "net/http"

type Client struct {
	host     string
	basePath string // tg-bot.com/bot<token>
	Client   http.Client
}

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		Client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}
