package indodax

import (
	"fmt"
	"net/http"
)

//
// Client represent common fields and environment Trading API
//
type Client struct {
	conn *http.Client
	env  *environment
	Info *UserInfo
}

//
// NewClient create and initialize new Indodax client.
//
// The token and secret parameters are used to authenticate the client when
// accessing private API.
//
// By default, the key and secret is read from environment variables
// "INDODAX_KEY" and "INDODAX_SECRET", the parameters will override the
// default value, if its set.
// If both environment variables and the parameters are empty, the client can
// only access the public API.
//
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

//
// TestAuthenticate is function to test connection the current client's using token and secret keys.
//
func (cl *Client) TestAuthentication() (err error) {

	// Test secet key by requesting User information
	_, err = cl.GetInfo()
	if err != nil {
		err = fmt.Errorf("Authenticate: " + err.Error())
		return err
	}

	return nil
}
