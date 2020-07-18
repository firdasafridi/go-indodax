package indodax

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//
// This method gives user balances, user wallet, user id, username, profile picture and server's timestamp.
//
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
	if respGetInfo.Success != 1 {
		return nil, fmt.Errorf("GetInfo: " + respGetInfo.Message)
	}

	cl.Info = respGetInfo.Return

	printDebug(cl.Info)

	return cl.Info, nil
}

//
// This method gives list of deposits and withdrawals of all currencies
//
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
	if respTransHistory.Success != 1 {
		return nil, fmt.Errorf("TransHistory: " + respTransHistory.Message)
	}

	printDebug(respTransHistory)

	return respTransHistory.Return, nil
}

//
// This method gives the list of current open orders (buy and sell) by pair.
//
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
	if respOpenOrders.Success != 1 {
		return nil, fmt.Errorf("OpenOrders: " + respOpenOrders.Message)
	}

	printDebug(respOpenOrders)

	return respOpenOrders.Return.Orders, nil
}

//
// This method gives the list of current open orders (buy and sell) all pair.
//
func (cl *Client) AllOpenOrders() (allOpenOrders map[string][]OpenOrders, err error) {
	respBody, err := cl.curlPrivate(apiViewOpenOrders, nil)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	respOpenOrders := &responseAllOpenOrders{}

	err = json.Unmarshal(respBody, respOpenOrders)
	if err != nil {
		err = fmt.Errorf("AllOpenOrders: " + err.Error())
		return nil, err
	}
	if respOpenOrders.Success != 1 {
		return nil, fmt.Errorf("AllOpenOrders: " + respOpenOrders.Message)
	}

	printDebug(respOpenOrders)

	return respOpenOrders.Return.Orders, nil
}

//
// This method gives information about transaction in buying and selling history
//
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
		err = fmt.Errorf("TradeHitory: " + err.Error())
		return nil, err
	}
	if respTradeHistory.Success != 1 {
		return nil, fmt.Errorf("TradeHitory: " + respTradeHistory.Message)
	}

	printDebug(respTradeHistory)

	return respTradeHistory.Return.Trades, nil
}

//
// This method gives the list of order history (buy and sell).
//
func (cl *Client) OrderHistory(
	pairName string,
	count, from int64,
) (openOrders []OrderHistory, err error) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}

	params := url.Values{}
	params.Set("pair", pairName)

	if count > 0 {
		params.Set("count", strconv.FormatInt(count, 10))
	}
	if from > 0 {
		params.Set("from", strconv.FormatInt(from, 10))
	}

	respBody, err := cl.curlPrivate(apiViewOrderHistory, params)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	respOrderHistory := &respOrderHistory{}

	err = json.Unmarshal(respBody, respOrderHistory)
	if err != nil {
		err = fmt.Errorf("OrderHistory: " + err.Error())
		return nil, err
	}
	if respOrderHistory.Success != 1 {
		return nil, fmt.Errorf("OrderHistory: " + respOrderHistory.Message)
	}

	printDebug(respOrderHistory)

	return respOrderHistory.Return.Orders, nil
}

//
// Use getOrder to get specific order details.
//
func (cl *Client) GetOrder(
	pairName string,
	orderId int64,
) (getOrder *GetOrder, err error) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}

	params := url.Values{}
	params.Set("pair", pairName)

	if orderId > 0 {
		params.Set("order_id", strconv.FormatInt(orderId, 10))
	}

	respBody, err := cl.curlPrivate(apiViewGetOrder, params)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	respGetOrders := &respGetOrders{}

	err = json.Unmarshal(respBody, respGetOrders)
	if err != nil {
		err = fmt.Errorf("GetOrder: " + err.Error())
		return nil, err
	}
	if respGetOrders.Success != 1 {
		return nil, fmt.Errorf("GetOrder: " + respGetOrders.Message)
	}

	printDebug(respGetOrders)

	return respGetOrders.Return.Order, nil
}

//
// This method is for canceling an existing open buy order.
//
func (cl *Client) CancelOrderBuy(
	pairName string,
	orderId int64,
) (cancelOrder *CancelOrder, err error) {
	cancelOrder, err = cl.cancelOrder("buy", pairName, orderId)
	if err != nil {
		return nil, err
	}
	return cancelOrder, nil
}

//
// This method is for canceling an existing open sell order.
//
func (cl *Client) CancelOrderSell(
	pairName string,
	orderId int64,
) (cancelOrder *CancelOrder, err error) {
	cancelOrder, err = cl.cancelOrder("sell", pairName, orderId)
	if err != nil {
		return nil, err
	}
	return cancelOrder, nil
}

//
// This method is for canceling an existing open order.
//
func (cl *Client) cancelOrder(
	method, pairName string,
	orderId int64,
) (cancelOrder *CancelOrder, err error) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}
	if orderId == 0 {
		return nil, ErrInvalidOrderID
	}

	params := url.Values{}
	params.Set("pair", pairName)
	params.Set("type", method)
	params.Set("order_id", strconv.FormatInt(orderId, 10))

	respBody, err := cl.curlPrivate(apiTradeCancelOrder, params)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	respCancelOrder := &respCancelOrder{}

	err = json.Unmarshal(respBody, respCancelOrder)
	if err != nil {
		err = fmt.Errorf("CancelOrder Unmarshal: " + err.Error())
		return nil, err
	}
	if respCancelOrder.Success != 1 {
		return nil, fmt.Errorf("CancelOrder: " + respCancelOrder.Error)
	}

	printDebug(respCancelOrder)

	return respCancelOrder.Return, nil
}

//
// This method is for opening a new buy order
//
func (cl *Client) TradeBuy(
	pairName string,
	price, amount float64,
) (trade *Trade, err error) {

	keyName := strings.Split(pairName, "_")
	if len(keyName) < 2 {
		return nil, ErrInvalidPairName
	}
	assetName := keyName[1]
	trade, err = cl.trade(
		"buy", pairName, assetName,
		price, amount,
	)

	if err != nil {
		return nil, err
	}

	return trade, nil
}

//
// This method is for opening a new sell order
//
func (cl *Client) TradeSell(
	pairName string,
	price, amount float64,
) (trade *Trade, err error) {

	keyName := strings.Split(pairName, "_")
	if len(keyName) < 2 {
		return nil, ErrInvalidPairName
	}
	assetName := keyName[0]
	trade, err = cl.trade(
		"sell", pairName, assetName,
		price, amount,
	)

	if err != nil {
		return nil, err
	}

	return trade, nil
}

//
// This method is for opening a new order
//
func (cl *Client) trade(
	method, pairName, assetName string,
	price, amount float64,
) (trade *Trade, err error) {
	if pairName == "" {
		return nil, ErrInvalidPairName
	}
	if assetName == "" {
		return nil, ErrInvalidAssetName
	}
	if price == 0 {
		return nil, ErrInvalidPrice
	}
	if amount == 0 {
		return nil, ErrInvalidAmount
	}

	params := url.Values{}
	params.Set("pair", pairName)
	params.Set("type", method)
	params.Set("price", fmt.Sprintf("%.8f", price))
	params.Set(assetName, fmt.Sprintf("%.8f", amount))

	respBody, err := cl.curlPrivate(apiTrade, params)
	if err != nil {
		return nil, err
	}

	printDebug(string(respBody))

	respTrade := &respTrade{}

	err = json.Unmarshal(respBody, respTrade)
	if err != nil {
		err = fmt.Errorf("trade: " + err.Error())
		return nil, err
	}
	if respTrade.Success != 1 {
		return nil, fmt.Errorf("Trade: " + respTrade.Error)
	}

	printDebug(respTrade)

	return respTrade.Return, nil
}
