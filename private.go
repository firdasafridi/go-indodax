package indodax

import (
	"encoding/json"
	"fmt"
	"net/url"
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
