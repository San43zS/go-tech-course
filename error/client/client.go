package client

import (
	"error/model"
	"net/http"
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
