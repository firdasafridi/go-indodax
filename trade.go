package indodax

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type respTrade struct {
	Success int
	Return  *Trade
	Message string
	Error   string
}

//
// Trade containt status of order placed by user like recive asset, spend asset, sold asset, remain asset, fee, order id placed, and last balance after trade.
//
type Trade struct {
	Receive          float64
	ReceiveAssetName string
	Spend            float64
	SpendAssetName   string
	Sold             float64
	SoldAssetName    string
	Remain           float64
	RemainAssetName  string
	Fee              float64
	OrderID          int64
	Balance          map[string]float64
}

func (trade *Trade) UnmarshalJSON(b []byte) (err error) {
	var kv map[string]interface{}

	err = json.Unmarshal(b, &kv)
	if err != nil {
		return err
	}

	for k, v := range kv {
		k = strings.ToLower(k)

		switch k {
		case fieldNameFee:
			trade.Fee, err = strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
		case fieldNameOrderID:
			orderIDFloat, err2 := strconv.ParseFloat(fmt.Sprintf("%f", v), 64)
			err = err2
			trade.OrderID = int64(orderIDFloat)
		case fieldNameBalance:
			trade.Balance, err = jsonToMapStringFloat64(v.(map[string]interface{}))
		default:
			keyName := strings.Split(k, "_")
			if len(keyName) < 2 {
				continue
			}
			switch keyName[0] {
			case fieldNameRemain:
				trade.Remain, err = strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
				trade.RemainAssetName = keyName[1]
			case fieldNameReceive:
				trade.Receive, err = strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
				trade.ReceiveAssetName = keyName[1]
			case fieldNameSpend:
				trade.Spend, err = strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
				trade.SpendAssetName = keyName[1]
			case fieldNameSold:
				trade.Sold, err = strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
				trade.SoldAssetName = keyName[1]
			}
		}
		if err != nil {
			return err
		}
	}

	return nil
}
