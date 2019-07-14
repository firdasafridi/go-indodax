package indodax

import (
	"encoding/json"
	"fmt"
)

func (cl *Client) GetTicker(pairName string) (ticker *Ticker, err error) {
	if pairName == "" {
		return ticker, ErrInvalidPairName
	}

	urlPath := fmt.Sprintf(pathTicker, pairName)

	body, err := cl.curlPublic(urlPath)
	if err != nil {
		return ticker, fmt.Errorf("GetTicker: " + err.Error())
	}

	printDebug(string(body))

	tickerRes := tickerResponse{}
	err = json.Unmarshal(body, &tickerRes)
	if err != nil {
		return ticker, fmt.Errorf("GetTicker: " + err.Error())
	}

	ticker, err = tickerRes.toTicker(pairName)
	if err != nil {
		return ticker, fmt.Errorf("GetTicker: " + err.Error())
	}

	return ticker, nil
}
