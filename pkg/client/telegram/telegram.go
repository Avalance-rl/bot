package telegram

import "net/http"

type Client struct {
	host     string
	basePath string
	client   *http.Client
}

func NewClient(host string, token string, client *http.Client) *Client {
	return &Client{
		host:     host,
		basePath: basePath,
		client:   client,
	}
}

func newBasePath(token string) string {
	return "bot" + token
}
