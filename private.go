package indodax

import (
	"encoding/json"
	"fmt"
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
