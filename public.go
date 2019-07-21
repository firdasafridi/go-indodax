package indodax

import (
	"encoding/json"
	"fmt"
)

//
// GetTicker containts the price summary like volume, last price, open buy, open sell of an individual pair.
//
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

//
// GetOrderBook containts the order book buy and sell of an individual pair.
//
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

//
// GetListTrades containts the historical trade of an individual pair.
//
func (cl *Client) GetListTrades(pairName string) (
	listTrade []*ListTrade, err error,
) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}

	urlPath := fmt.Sprintf(pathTrades, pairName)

	body, err := cl.curlPublic(urlPath)
	if err != nil {
		return nil, fmt.Errorf("GetListTrades: " + err.Error())
	}

	printDebug(string(body))

	err = json.Unmarshal(body, &listTrade)
	if err != nil {
		return nil, fmt.Errorf("GetListTrades: " + err.Error())
	}

	return listTrade, nil
}

//
// GetSummaries containts the price summary like volume, last price, open buy, open sell of all pair.
//
func (cl *Client) GetSummaries() (summaries *Summary, err error) {

	urlPath := pathSummaries
	body, err := cl.curlPublic(urlPath)
	if err != nil {
		return nil, fmt.Errorf("GetSummaries: " + err.Error())
	}

	printDebug(string(body))

	err = json.Unmarshal(body, &summaries)
	if err != nil {
		return nil, fmt.Errorf("GetSummaries: " + err.Error())
	}

	return summaries, nil
}
