package indodax

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type WsRequestHandler func(conn *websocket.Conn) error

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint string
}

func newWsConfig(endpoint string) *WsConfig {
	return &WsConfig{
		Endpoint: endpoint,
	}
}

type basicRequest struct {
	Id int `json:"id"`
}

type AuthenticationResponse struct {
	BasicResponse
	Result AuthenticationResponseResult `json:"result"`
}

var wsServe = func(cfg *WsConfig, request WsRequestHandler, handler WsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	Dialer := websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: false,
	}

	c, _, err := Dialer.Dial(cfg.Endpoint, nil)
	if err != nil {
		return nil, nil, err
	}
	c.SetReadLimit(655350)
	doneC = make(chan struct{})
	stopC = make(chan struct{})
	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer close(doneC)

		// First send authentication
		if err := wsAuthenticate(c); err != nil {
			errHandler(err)
			return
		}

		if err := request(c); err != nil {
			errHandler(err)
			return
		}

		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout)
		}

		// Wait for the stopC channel to be closed.  We do that in a
		// separate goroutine because ReadMessage is a blocking
		// operation.
		silent := false
		go func() {
			select {
			case <-stopC:
				silent = true
			case <-doneC:
			}
			_ = c.Close()
		}()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if !silent {
					errHandler(err)
				}
				return
			}
			handler(message)
		}
	}()
	return
}

type authenticationRequestParam struct {
	Token string `json:"token"`
}

type authenticationRequest struct {
	basicRequest
	Params authenticationRequestParam `json:"params"`
}

type BasicResponse struct {
	Id int `json:"id"`
}

type AuthenticationResponseResult struct {
	Client  string `json:"client"`
	Version string `json:"version"`
	Expires bool   `json:"expires"`
	Ttl     int    `json:"ttl"`
}

func wsAuthenticate(c *websocket.Conn) error {
	req := authenticationRequest{
		basicRequest: basicRequest{
			Id: 1,
		},
		Params: authenticationRequestParam{
			// This is documented here: https://github.com/btcid/indodax-official-api-docs/blob/master/ws3-websocket.md#authentication
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5NDY2MTg0MTV9.UR1lBM6Eqh0yWz-PVirw1uPCxe60FdchR8eNVdsskeo",
		},
	}
	if err := c.WriteJSON(req); err != nil {
		return fmt.Errorf("authentication error sending: %w", err)
	}

	_, message, err := c.ReadMessage()
	if err != nil {
		return fmt.Errorf("authentication error reading response: %w", err)
	}

	res := &AuthenticationResponseResult{}
	if err = json.Unmarshal(message, res); err != nil {
		return fmt.Errorf("authentication error parsing response: %w", err)
	}

	return nil
}

func keepAlive(c *websocket.Conn, timeout time.Duration) {
	tckr := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer tckr.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-tckr.C
			if time.Since(lastResponse) > timeout {
				_ = c.Close()

				return
			}
		}
	}()
}
