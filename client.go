package indodax

import (
	"fmt"
	"net/http"
)

type Client struct {
	conn *http.Client
	env  *environment
	Info *UserInfo
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

func (cl *Client) TestAuthentication() (err error) {

	// Test secet key by requesting User information
	_, err = cl.GetInfo()
	if err != nil {
		err = fmt.Errorf("Authenticate: " + err.Error())
		return err
	}

	return nil
}
