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

