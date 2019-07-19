package indodax

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (cl *Client) GetInfo() (usrInfo *UserInfo, err error) {
	respBody, err := cl.curlPrivate(apiViewGetInfo, nil)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	respGetInfo := &respGetInfo{}

	err = json.Unmarshal(respBody, respGetInfo)
	if err != nil {
		err = fmt.Errorf("GetInfo: " + err.Error())
		return nil, err
	}

	cl.Info = respGetInfo.Return

	printDebug(cl.Info)

	return cl.Info, nil
}

func (cl *Client) TransHistory() (transHistory *TransHistory, err error) {
	respBody, err := cl.curlPrivate(apiViewTransactionHistory, nil)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	respTransHistory := &respTransHistory{}

	err = json.Unmarshal(respBody, respTransHistory)
	if err != nil {
		err = fmt.Errorf("TransHistory: " + err.Error())
		return nil, err
	}

	printDebug(respTransHistory)

	return respTransHistory.Return, nil
}

func (cl *Client) OpenOrders(pairName string) (openOrders []OpenOrders, err error) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}

	params := url.Values{}
	params.Set("pair", pairName)

	respBody, err := cl.curlPrivate(apiViewOpenOrders, params)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	respOpenOrders := &responseOpenOrders{}

	err = json.Unmarshal(respBody, respOpenOrders)
	if err != nil {
		err = fmt.Errorf("OpenOrders: " + err.Error())
		return nil, err
	}

	printDebug(respOpenOrders)

	return respOpenOrders.Return.Orders, nil
}

func (cl *Client) AllOpenOrders() (allOpenOrders map[string][]OpenOrders, err error) {
	respBody, err := cl.curlPrivate(apiViewOpenOrders, nil)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	respOpenOrders := &responseAllOpenOrders{}

	err = json.Unmarshal(respBody, respOpenOrders)
	if err != nil {
		err = fmt.Errorf("OpenOrders: " + err.Error())
		return nil, err
	}

	printDebug(respOpenOrders)

	return respOpenOrders.Return.Orders, nil
}

func (cl *Client) TradeHitory(
	pairName string, 
	count, startTradeID, endTradeID int64,
	sortOrder string,
	sinceTime *time.Time,
	endTime *time.Time,
	) (openOrders []TradeHistory, err error) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}

	params := url.Values{}
	params.Set("pair", pairName)

	if count > 0 {
		params.Set("count", strconv.FormatInt(count, 10))
	}
	if startTradeID > 0 {
		params.Set("from_id", strconv.FormatInt(startTradeID, 10))
	}
	if endTradeID > 0 {
		params.Set("end_id", strconv.FormatInt(endTradeID, 10))
	}

	sortOrder = strings.ToLower(sortOrder)
	switch sortOrder {
	case "asc":
		params.Set("order", "asc")
	case "desc":
		params.Set("order", "desc")
	}

	if sinceTime != nil {
		params.Set("since", strconv.FormatInt(sinceTime.Unix(), 10))
	}
	if endTime != nil {
		params.Set("end", strconv.FormatInt(endTime.Unix(), 10))
	}

	respBody, err := cl.curlPrivate(apiViewTradeHistory, params)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	respTradeHistory := &respTradeHistory{}

	err = json.Unmarshal(respBody, respTradeHistory)
	if err != nil {
		err = fmt.Errorf("OpenOrders: " + err.Error())
		return nil, err
	}

	printDebug(respTradeHistory)

	return respTradeHistory.Return.Trades, nil
}
