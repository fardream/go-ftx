package ftx

import (
	"net/http"
	"net/url"
)

// Client is the client to ftx api.
type Client struct {
	Client     *http.Client
	Api        string
	Secret     []byte
	Subaccount string
}

// NewClient creates a new client
func NewClient(api, secret, subaccount string) (*Client, error) {
	return &Client{
		Client:     &http.Client{},
		Api:        api,
		Secret:     []byte(secret),
		Subaccount: url.PathEscape(subaccount),
	}, nil
}
