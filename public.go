package indodax

import (
	"encoding/json"
	"fmt"
)

func (cl *Client) GetTicker(pairName string) (ticker *Ticker, err error) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}

	urlPath := fmt.Sprintf(pathTicker, pairName)

	body, err := cl.curlPublic(urlPath)
	if err != nil {
		return nil, fmt.Errorf("GetTicker: " + err.Error())
	}

	printDebug(string(body))

	tickerRes := tickerResponse{}
	err = json.Unmarshal(body, &tickerRes)
	if err != nil {
		return nil, fmt.Errorf("GetTicker: " + err.Error())
	}

	ticker, err = tickerRes.toTicker(pairName)
	if err != nil {
		return nil, fmt.Errorf("GetTicker: " + err.Error())
	}

	return ticker, nil
}

func (cl *Client) GetOrderBook(pairName string) (orderBook *OrderBook, err error) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}

	urlPath := fmt.Sprintf(pathDepth, pairName)

	body, err := cl.curlPublic(urlPath)
	if err != nil {
		return nil, fmt.Errorf("GetOrderBook: " + err.Error())
	}

	printDebug(string(body))

	orderBook = &OrderBook{}
	err = json.Unmarshal(body, &orderBook)
	if err != nil {
		return nil, fmt.Errorf("GetOrderBook: " + err.Error())
	}

	return orderBook, nil
}
