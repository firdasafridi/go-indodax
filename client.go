package indodax

import (
	"net/http"
)

type Client struct {
	conn *http.Client
	env  *environment
}

func NewClient(key, secret string) (cl *Client, err error) {
	cl = &Client{
		conn: &http.Client{},
		env:  newEnvironment(),
	}

	if key != "" {
		cl.env.apiKey = key
		cl.env.apiSecret = secret
	}
	return cl, nil
}

