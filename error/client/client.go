package client

import (
	"net/http"

	"error/model"
)

type Client struct {
	Client *http.Client
}

func New() *Client {
	client := &Client{
		Client: &http.Client{
			Timeout: model.TimeOut,
		},
	}

	return client
}
