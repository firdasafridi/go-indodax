package indodax

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/gorilla/websocket"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const (
	wsMainURL = "wss://ws3.indodax.com/ws/"
)

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 60
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = false
)

type BaseQuoteVolumePrice struct {
	BaseVolume  float64 `json:"base_volume"`
	QuoteVolume float64 `json:"quote_volume"`
	Price       float64 `json:"price"`
}

type OrderBookEvent struct {
	Pair string                 `json:"pair"`
	Ask  []BaseQuoteVolumePrice `json:"ask"`
	Bid  []BaseQuoteVolumePrice `json:"bid"`
}

type WsOrderBookEventHandler func(event *OrderBookEvent)

type OrderBookSymbol struct {
	Base  string
	Quote string
}

type ChannelRequestParam struct {
	Channel string `json:"channel"`
}

type ChannelRequest struct {
	basicRequest
	Method int                 `json:"method"`
	Params ChannelRequestParam `json:"params"`
}

func WsOrderBookServe(symbol OrderBookSymbol, handler WsOrderBookEventHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	baseSymbol := strings.ToLower(symbol.Base)
	quoteSymbol := strings.ToLower(symbol.Quote)
	channel := fmt.Sprintf("market:order-book-%s%s", baseSymbol, quoteSymbol)

	cfg := newWsConfig(wsMainURL)
	wsReqHandler := func(c *websocket.Conn) error {
		return c.WriteJSON(ChannelRequest{
			basicRequest: basicRequest{
				Id: 2,
			},
			Method: 1,
			Params: ChannelRequestParam{
				Channel: channel,
			},
		})
	}
	lastOffset := &atomic.Int64{}
	wsHandler := func(message []byte) {
		prevOffset := lastOffset.Load()

		e, offset, sErr, isTargetedMessage := toOrderBookEvent(channel, symbol, prevOffset, message)
		// It's not targeted message, ignore it
		if !isTargetedMessage {
			return
		}

		// Error when serializing message, ignore it
		if sErr != nil {
			errHandler(sErr)
			return
		}

		// It's older message, ignore it
		if prevOffset >= offset {
			return
		}

		lastOffset.Store(offset)

		handler(e)
	}

	return wsServe(cfg, wsReqHandler, wsHandler, errHandler)
}

func toOrderBookEvent(
	channel string, symbol OrderBookSymbol, lastOffset int64, message []byte,
) (e *OrderBookEvent, offset int64, err error, isTargetedMessage bool) {
	rMessage, _, _, err := jsonparser.Get(message, "result")
	if err != nil {
		return nil, 0, nil, false
	}

	gotChannel, err := jsonparser.GetString(rMessage, "channel")
	if err != nil || gotChannel != channel {
		return nil, 0, nil, false
	}

	dMessage, _, _, err := jsonparser.Get(rMessage, "data")
	if err != nil {
		return nil, 0, nil, false
	}

	offset, err = jsonparser.GetInt(dMessage, "offset")
	if err != nil {
		return nil, 0, fmt.Errorf("fail to retrive offset: %w", err), true
	}

	if lastOffset >= offset {
		return nil, offset, nil, true
	}

	dMessage, _, _, err = jsonparser.Get(dMessage, "data")
	if err != nil {
		return nil, 0, nil, false
	}

	e = &OrderBookEvent{}

	paths := [][]string{
		{"pair"},
		{"ask"},
		{"bid"},
	}
	hasError := false
	baseVolume := symbol.Base + "_volume"
	quoteVolume := symbol.Quote + "_volume"
	price := "price"

	jsonparser.EachKey(dMessage, func(idx int, bytes []byte, vt jsonparser.ValueType, pErr error) {
		hasError = err != nil || pErr != nil
		if hasError {
			err = firstError(err, pErr)
			return
		}

		switch idx {
		case 0: // pair
			e.Pair = string(bytes)
		case 1: // ask
			e.Ask, err = toBaseQuoteVolumePrices(bytes, baseVolume, quoteVolume, price)
		case 2: // bid
			e.Bid, err = toBaseQuoteVolumePrices(bytes, baseVolume, quoteVolume, price)
		}
	}, paths...)

	return e, offset, err, true
}

func toBaseQuoteVolumePrices(
	ba []byte, baseVolumeName, quoteVolumeName, price string,
) (r []BaseQuoteVolumePrice, err error) {
	paths := [][]string{
		{baseVolumeName},
		{quoteVolumeName},
		{price},
	}

	r = make([]BaseQuoteVolumePrice, 0)
	hasError := false
	_, err = jsonparser.ArrayEach(ba, func(ba []byte, _ jsonparser.ValueType, _ int, pErr error) {
		hasError = hasError || pErr != nil
		if hasError {
			err = firstError(err, pErr)
			return
		}

		bqvp := BaseQuoteVolumePrice{}
		jsonparser.EachKey(ba, func(idx int, ba []byte, _ jsonparser.ValueType, pErr error) {
			hasError = hasError || pErr != nil
			if hasError {
				err = firstError(err, pErr)

				return
			}

			asString := string(ba)
			v, vErr := strconv.ParseFloat(asString, 64)
			if vErr != nil {
				err = fmt.Errorf("fail to parse price %s: %w", asString, vErr)
				hasError = true

				return
			}

			switch idx {
			case 0: // base
				bqvp.BaseVolume = v
			case 1: // quote
				bqvp.QuoteVolume = v
			case 2: // price
				bqvp.Price = v
			}
		}, paths...)

		r = append(r, bqvp)
	})

	return r, err
}

func firstError(fst, snd error) error {
	if fst != nil {
		return fst
	}

	return snd
}
